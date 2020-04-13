package project

import (
	"fmt"
	"github.com/kgysu/oc-wrapper/config"
	"github.com/kgysu/oc-wrapper/fileutils"
)

func NewProjectsFromDisk(currentDir string, namespace string) ([]*OpenshiftProject, error) {
	rootDir := currentDir + config.GetRootFolderOrDefault() + "/"
	folders, err := fileutils.FoldersInDir(rootDir)
	if err != nil {
		return nil, err
	}

	var projects []*OpenshiftProject
	for _, folder := range folders {
		projectFromDisk, err := NewProjectFromDisk(currentDir, folder, namespace)
		if err != nil {
			// only log on error
			fmt.Println(err.Error())
		}
		projects = append(projects, projectFromDisk)
	}

	return projects, nil
}

func NewProjectFromDisk(currentDir string, projectName string, namespace string) (*OpenshiftProject, error) {
	var items []OpenshiftItemInterface

	rootDir := currentDir + config.GetRootFolderOrDefault()
	projectDir := rootDir + "/" + projectName
	envsDir := rootDir + "/" + projectName + "/" + namespace

	err := checkProjectPathStructure(currentDir, projectName)
	if err != nil {
		return nil, err
	}

	yamlFiles, err := fileutils.FilesInDir(projectDir)
	if err != nil {
		return nil, err
	}

	envfiles, err := fileutils.FilesInDir(envsDir)
	if err != nil {
		return nil, err
	}

	// TODO: environment specific Items
	//envYamlFiles := filterFilesByType(envfiles, ".yaml")
	envEnvFiles := fileutils.FilterFilesByType(envfiles, ".env")

	envsMap, err := fileutils.EnvFilesToMap(envEnvFiles)
	if err != nil {
		return nil, err
	}

	for _, file := range yamlFiles {
		item, err := NewOpenshiftItemFromFile(file, envsMap)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			items = append(items, item)
		}
	}

	// Create Project
	op := NewOpenshiftProject(projectName)
	op.Items = items
	return op, nil
}
