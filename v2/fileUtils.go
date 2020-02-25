package v2

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const JsonFileSuffix = ".json"
const EnvFileSuffix = ".env"
const ProjectSettingsFileName = "project" + JsonFileSuffix

type ThisProjectMeta struct {
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
}

func parseProjectFile(filePath string) (*ThisProjectMeta, error) {
	data, err := ioutil.ReadFile(filePath + "/" + ProjectSettingsFileName)
	if err != nil {
		return nil, err
	}
	thisProjectMeta := &ThisProjectMeta{}
	err = json.Unmarshal(data, &thisProjectMeta)
	if err != nil {
		return nil, err
	}
	return thisProjectMeta, nil
}

func parseFileAndReplaceEnvVars(filePath string, envs map[string]string, v interface{}, debug bool) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	dataAsString := string(data)
	replacedString := dataAsString
	for key, value := range envs {
		replacedString = strings.ReplaceAll(replacedString, "$"+key, value)
	}
	if debug {
		fmt.Println(replacedString)
	}
	err = json.Unmarshal([]byte(replacedString), &v)
	if err != nil {
		return err
	}
	return nil
}

func loadEnvsForNamespace(namespace string, envFiles []string, debug bool) (map[string]string, error) {
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
