package util

import (
	"bufio"
	"fmt"
	"github.com/kgysu/oc-wrapper/config"
	"github.com/kgysu/oc-wrapper/project"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// File Helpers

func checkProjectPath(op *project.OpenshiftProject, currentDir string) error {
	err := createIfNotExists(currentDir + "/projects")
	if err != nil {
		return err
	}
	err = createIfNotExists(currentDir + "/projects" + "/" + op.Name)
	if err != nil {
		return err
	}
	return nil
}

func createIfNotExists(folder string) error {
	if !existsFile(folder) {
		err := os.Mkdir(folder, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func existsFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}

func getCurrentDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return dir, nil
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

func filterFilesByType(files []string, fileType string) []string {
	var resultFiles []string
	for _, file := range files {
		if strings.HasSuffix(file, fileType) {
			resultFiles = append(resultFiles, file)
		}
	}
	return resultFiles
}

func getEnvFileByName(files []string, name string) string {
	for _, file := range files {
		if strings.HasPrefix(file, name) {
			return file
		}
	}
	return ""
}

func envFilesToMap(files []string) (map[string]string, error) {
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
				if config.IsInDebugMode() {
					fmt.Printf("Found Env: %s=%s \n", splittedEnv[0], splittedEnv[1])
				}
				currentEnvs[splittedEnv[0]] = splittedEnv[1]
			}
		}
		if err := scanner.Err(); err != nil {
			return nil, err
		}
	}
	return currentEnvs, nil
}

func filesInDir(root string) ([]string, error) {
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

func foldersInDir(root string) ([]string, error) {
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
