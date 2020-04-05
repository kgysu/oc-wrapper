package util

import (
	"fmt"
	"github.com/kgysu/oc-wrapper/config"
	"github.com/kgysu/oc-wrapper/project"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

func NewProjectFromNamespace(projectName string, namespace string, restConf *rest.Config, options v12.ListOptions) (*project.OpenshiftProject, error) {
	var items []project.OpenshiftItemInterface

	// TODO add more Types
	// Add All
	newItems, err := ListAll(namespace, restConf, options)
	if err != nil {
		return nil, err
	}
	items = append(items, newItems...)

	// Create Project
	op := project.NewOpenshiftProject(projectName)
	op.Items = items
	return op, nil
}

func NewProjectsFromDisk(namespace string) ([]*project.OpenshiftProject, error) {
	dir, err := getCurrentDir()
	if err != nil {
		return nil, err
	}

	rootDir := dir + config.GetRootFolderOrDefault() + "/"
	fmt.Println(rootDir)
	folders, err := foldersInDir(rootDir)
	if err != nil {
		return nil, err
	}

	var projects []*project.OpenshiftProject
	for _, folder := range folders {
		projectFromDisk, err := NewProjectFromDisk(folder, namespace)
		if err != nil {
			// only log on error
			fmt.Println(err.Error())
		}
		projects = append(projects, projectFromDisk)
	}

	return projects, nil
}

func NewProjectFromDisk(projectName string, namespace string) (*project.OpenshiftProject, error) {
	var items []project.OpenshiftItemInterface
	dir, err := getCurrentDir()
	if err != nil {
		return nil, err
	}

	rootDir := dir + config.GetRootFolderOrDefault()
	projectDir := rootDir + "/" + projectName
	if !existsFile(rootDir) || !existsFile(projectDir) {
		return nil, fmt.Errorf("no projects found")
	}
	envsDir := rootDir + "/" + projectName + "/" + namespace
	err = createIfNotExists(envsDir)
	if err != nil {
		return nil, err
	}

	yamlFiles, err := filesInDir(projectDir)
	if err != nil {
		return nil, err
	}

	envfiles, err := filesInDir(envsDir)
	if err != nil {
		return nil, err
	}

	// TODO: environment specific Items
	//envYamlFiles := filterFilesByType(envfiles, ".yaml")
	envEnvFiles := filterFilesByType(envfiles, ".env")

	envsMap, err := envFilesToMap(envEnvFiles)
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
	op := project.NewOpenshiftProject(projectName)
	op.Items = items
	return op, nil
}

func SaveProjectToDisk(op *project.OpenshiftProject) error {
	dir, err := getCurrentDir()
	if err != nil {
		return err
	}

	// create Folders if not exists
	err = checkProjectPath(op, dir)
	if err != nil {
		return err
	}

	// write Item files
	for _, item := range op.Items {
		filePath := dir + config.GetRootFolderOrDefault() + "/" + op.Name + "/" + item.GetFileName()
		err := item.WriteToFile(filePath)
		if err != nil {
			return err
		}
	}
	return nil
}
