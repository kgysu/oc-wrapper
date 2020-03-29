package files

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ReadFile(file string) ([]byte, error) {
	return ioutil.ReadFile(file)
}

func ReadAllFilesInFolder(folder string) (map[string][]byte, error) {
	files, err := filePathWalkDir(folder)
	if err != nil {
		return nil, err
	}
	filesData := make(map[string][]byte)
	for _, file := range files {
		fileData, err := ReadFile(file)
		if err != nil {
			log.Fatal(err)
		} else {
			filesData[file] = fileData
		}
	}
	return filesData, nil
}

func ReadConfigFile(folder string) ([]byte, error) {
	files, err := filePathWalkDir(folder)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if strings.HasSuffix(file, ConfigFileSuffix) {
			fileData, err := ReadFile(file)
			if err != nil {
				log.Fatal(err)
			} else {
				return fileData, nil
			}
		}
	}
	return nil, fmt.Errorf("No config found in folder [%s]\n", folder)
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
