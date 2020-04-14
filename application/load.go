package application

import (
	"fmt"
	"github.com/kgysu/oc-wrapper/appitem"
	"github.com/kgysu/oc-wrapper/fileutils"
)

func NewAppsFromDisk(path, namespace string) ([]*Application, error) {
	folders, err := fileutils.FoldersInDir(path)
	if err != nil {
		return nil, err
	}

	var apps []*Application
	for _, folder := range folders {
		appFromDisk, err := NewAppFromDisk(path, folder, namespace)
		if err != nil {
			// only log on error
			fmt.Println(err.Error())
		}
		apps = append(apps, appFromDisk)
	}

	return apps, nil
}

func NewAppFromDisk(path, appName, namespace string) (*Application, error) {
	var items []appitem.AppItem

	appDir := path + "/" + appName
	envsDir := path + "/" + appName + "/" + namespace

	err := checkAppPathStructure(path, appName)
	if err != nil {
		return nil, err
	}

	err = fileutils.CreateIfNotExists(envsDir)
	if err != nil {
		return nil, err
	}

	yamlFiles, err := fileutils.FilesInDir(appDir)
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
		item, err := appitem.NewAppItemFromFile(file, envsMap)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			items = append(items, item)
		}
	}

	// Create App
	op := NewApp(appName)
	op.Items = items
	return op, nil
}
