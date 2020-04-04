package util

import (
	"fmt"
	"github.com/kgysu/oc-wrapper/project"
	"io/ioutil"
	"os"
	"path/filepath"
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
		fmt.Printf("created dir [%s]\n", folder)
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
