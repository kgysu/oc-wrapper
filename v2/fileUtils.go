package v2

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	v1 "k8s.io/api/core/v1"
	"os"
	"path/filepath"
	"strings"
)

const JsonFileSuffix = ".json"
const EnvFileSuffix = ".env"
const ProjectSettingsFileName = "project" + JsonFileSuffix
const ConfigMapDataFolder = "config"
const EnvVarPrefix = "$"

type ThisProjectMeta struct {
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
}

func parseProjectFile(filePath string) (*ThisProjectMeta, error) {
	dataAsString, err := readFileData(filePath + "/" + ProjectSettingsFileName)
	if err != nil {
		return nil, err
	}
	thisProjectMeta := &ThisProjectMeta{}
	err = json.Unmarshal(dataAsString, &thisProjectMeta)
	if err != nil {
		return nil, err
	}
	return thisProjectMeta, nil
}

func parseFileAndReplaceEnvVars(filePath string, envs map[string]string, v interface{}, debug bool) error {
	dataAsString, err := readFileAndReplaceEnvs(filePath, envs)
	if err != nil {
		return err
	}
	if debug { // debug is an user requirement here
		fmt.Printf("ItemResult: [%v] \n", dataAsString)
	}
	err = json.Unmarshal([]byte(dataAsString), &v)
	if err != nil {
		return err
	}
	return nil
}

func parseConfigMap(filePath string, folderPath string, namespace string, envs map[string]string, cm *v1.ConfigMap, debug bool) error {
	dataAsString, err := readFileAsString(filePath)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(dataAsString), &cm)
	if err != nil {
		return err
	}
	// Only use files in /config if Data section is empty
	if len(cm.Data) == 0 {
		fmt.Println("Data ist empty")
		configMapData, err := getConfigMapDataByNamespace(folderPath, namespace, envs, debug)
		if err != nil {
			return err
		}
		cm.Data = configMapData
	}

	if debug { // debug is an user requirement here
		for key, val := range cm.Data {
			fmt.Printf("key=[%s], val=[%s]", key, val)
		}
	}
	return nil
}

func getConfigMapDataByNamespace(folder string, namespace string, envs map[string]string, debug bool) (map[string]string, error) {
	configMapData := make(map[string]string)
	folderPath := folder + "/" + ConfigMapDataFolder + "/" + namespace
	// Load all config files in <folder>/<config>/<namespace>/*
	files, err := filePathWalkDir(folderPath)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		fmt.Println(file)
		fileDataAsString, err := readFileAndReplaceEnvs(file, envs)
		if err != nil {
			onlyLogOnError(err)
			continue
		}
		fileName := strings.Replace(file, folderPath+"/", "", 1)
		configMapData[fileName] = fileDataAsString
	}
	if debug {
		for key, val := range configMapData {
			fmt.Printf("key=[%s], val=[%s] \n", key, val)
		}
	}
	return configMapData, nil
}

func readFileAndReplaceEnvs(file string, envs map[string]string) (string, error) {
	dataString, err := readFileAsString(file)
	if err != nil {
		return "", err
	}
	return replaceEnvs(dataString, envs), nil
}

func replaceEnvs(input string, envs map[string]string) string {
	result := input
	for key, value := range envs {
		result = strings.ReplaceAll(result, EnvVarPrefix+key, value)
	}
	return result
}

func readFileAsString(file string) (string, error) {
	data, err := readFileData(file)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func readFileData(file string) ([]byte, error) {
	return ioutil.ReadFile(file)
}

func loadEnvsByNamespace(namespace string, envFiles []string, debug bool) (map[string]string, error) {
	fileNameSuffixForCurrentNs := namespace + EnvFileSuffix
	currentEnvs := make(map[string]string)
	for _, envFile := range envFiles {
		if strings.HasSuffix(envFile, fileNameSuffixForCurrentNs) {
			envData, err := os.Open(envFile)
			if err != nil {
				return nil, err
			}
			scanner := bufio.NewScanner(envData)
			for scanner.Scan() {
				text := scanner.Text()
				if !(strings.HasPrefix(text, "#")) {
					splittedEnv := strings.SplitN(text, "=", 2)
					if len(splittedEnv) == 2 {
						if debug {
							fmt.Printf("Found: %s=%s \n", splittedEnv[0], splittedEnv[1])
						}
						currentEnvs[splittedEnv[0]] = splittedEnv[1]
					}
				}
			}
			if err := scanner.Err(); err != nil {
				return nil, err
			}
		}
	}
	return currentEnvs, nil
}

func filePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func splitByFileType(filePaths []string) ([]string, []string) {
	var envFiles []string
	var jsonFiles []string
	for _, filePath := range filePaths {
		if strings.HasSuffix(filePath, EnvFileSuffix) {
			envFiles = append(envFiles, filePath)
		}
		if strings.HasSuffix(filePath, JsonFileSuffix) {
			jsonFiles = append(jsonFiles, filePath)
		}
	}
	return envFiles, jsonFiles
}
