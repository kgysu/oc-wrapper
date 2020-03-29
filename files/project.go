package files

import (
	"fmt"
	yamld "gopkg.in/yaml.v2"
)

const ConfigFileSuffix = "config.yaml"

type OpenshiftProject struct {
	Name       string
	ConfigFile string
	local      bool
	Items      []OpenshiftItem
}

type OpenshiftItem struct {
	Name    string
	Kind    string
	File    string
	Data    string
	RawData []byte
}

//func NewFromFolder(folder string) (*OpenshiftProject, error) {
//	fileDatas, err := ReadAllFilesInFolder(folder)
//	if err != nil {
//		return nil, err
//	}
//
//	var folderItems []OpenshiftItem
//	for file, data := range fileDatas {
//		if strings.HasSuffix(file, "config.yaml") {
//
//		}
//		item := NewOpenshiftItemFromFile("", file, data)
//		folderItems = append(folderItems, item)
//	}
//
//	project := &OpenshiftProject{
//		Name:       "",
//		ConfigFile: "",
//		local:      true,
//		Items:      folderItems,
//	}
//
//	return project, nil
//}

func NewProjectFromConfig(folder string) (*OpenshiftProject, error) {
	configData, err := ReadConfigFile(folder)
	if err != nil {
		return nil, err
	}
	var newProject OpenshiftProject
	err = yamld.Unmarshal(configData, &newProject)
	if err != nil {
		return nil, err
	}
	return &newProject, nil
}

func NewOpenshiftItemFromFile(name string, file string, kind string, rawdata []byte) OpenshiftItem {
	return OpenshiftItem{
		Name:    name,
		Kind:    kind,
		File:    file,
		RawData: rawdata,
	}
}

func NewOpenshiftItem(name string, file string, kind string, data string) OpenshiftItem {
	return OpenshiftItem{
		Name: name,
		Kind: kind,
		File: file,
		Data: data,
	}
}

func (op *OpenshiftProject) LoadProjectFilesData() {
	for i, item := range op.Items {
		fileData, err := ReadFile(item.File)
		if err != nil {
			fmt.Printf("could not read file [%s]\n", item.File)
		}
		op.Items[i].RawData = fileData
		op.Items[i].Data = string(fileData)
	}
}
