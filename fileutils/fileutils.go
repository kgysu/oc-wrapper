package fileutils

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func ReadFile(file string) ([]byte, error) {
	return ioutil.ReadFile(file)
}

func CreateFile(filename string, data []byte) error {
	return ioutil.WriteFile(filename, data, 0644)
}

func ReplaceEnvs(content string, envs map[string]string) string {
	result := content
	for key, value := range envs {
		result = strings.ReplaceAll(result, "${"+key+"}", value)
	}
	return result
}

func CreateIfNotExists(folder string) error {
	if !ExistsFile(folder) {
		err := os.Mkdir(folder, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func ExistsFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}

func GetCurrentDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return dir, nil
}

func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func FilterFilesByType(files []string, fileType string) []string {
	var resultFiles []string
	for _, file := range files {
		if strings.HasSuffix(file, fileType) {
			resultFiles = append(resultFiles, file)
		}
	}
	return resultFiles
}

func GetEnvFileByName(files []string, name string) string {
	for _, file := range files {
		if strings.HasPrefix(file, name) {
			return file
		}
	}
	return ""
}

func EnvFilesToMap(files []string) (map[string]string, error) {
	currentEnvs := make(map[string]string)
	for _, envFile := range files {
		envData, err := os.Open(envFile)
		if err != nil {
			return nil, err
		}
		scanner := bufio.NewScanner(envData)
		for scanner.Scan() {
			text := scanner.Text()
			splittedEnv := strings.SplitN(text, "=", 2)
			if len(splittedEnv) == 2 {
				currentEnvs[splittedEnv[0]] = splittedEnv[1]
			}
		}
		if err := scanner.Err(); err != nil {
			return nil, err
		}
	}
	return currentEnvs, nil
}

func FilesInDir(root string) ([]string, error) {
	var resultFiles []string
	files, err := ioutil.ReadDir(root)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if !file.IsDir() {
			resultFiles = append(resultFiles, root+"/"+file.Name())
		}
	}
	return resultFiles, err
}

func FoldersInDir(root string) ([]string, error) {
	var folders []string
	files, err := ioutil.ReadDir(root)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.IsDir() {
			folders = append(folders, file.Name())
		}
	}
	return folders, err
}
