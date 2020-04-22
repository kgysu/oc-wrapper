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
	folder = filepath.FromSlash(folder)
	if !ExistsFile(folder) {
		err := os.Mkdir(folder, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func ExistsFile(file string) bool {
	file = filepath.FromSlash(file)
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
	return filepath.FromSlash(dir), nil
}

func FilePathWalkDir(root string) ([]string, error) {
	root = filepath.FromSlash(root)
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, filepath.FromSlash(path))
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
	root = filepath.FromSlash(root)
	var resultFiles []string
	files, err := ioutil.ReadDir(root)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if !file.IsDir() {
			resultFiles = append(resultFiles, filepath.FromSlash(root+"/"+file.Name()))
		}
	}
	return resultFiles, err
}

func FoldersInDir(root string) ([]string, error) {
	root = filepath.FromSlash(root)
	var folders []string
	files, err := ioutil.ReadDir(root)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.IsDir() {
			folders = append(folders, filepath.FromSlash(file.Name()))
		}
	}
	return folders, err
}
