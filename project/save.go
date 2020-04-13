package project

import (
	"fmt"
	"github.com/kgysu/oc-wrapper/config"
	"github.com/kgysu/oc-wrapper/fileutils"
	"io"
)

func (op *OpenshiftProject) Save(w io.Writer) {
	dir, err := fileutils.GetCurrentDir()
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	// create Folders if not exists
	err = checkProjectPathStructure(dir, op.Name)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	// write Item files
	for _, item := range op.Items {
		filePath := dir + config.GetRootFolderOrDefault() + "/" + op.Name + "/" + item.GetFileName()
		err := item.WriteToFile(filePath)
		if err != nil {
			fmt.Fprintln(w, err.Error())
		}
	}
	fmt.Fprintf(w, "saved project [%s] in [%s] \n", op.Name, config.GetRootFolderOrDefault())
}

func checkProjectPathStructure(currentDir string, projectName string) error {
	err := fileutils.CreateIfNotExists(currentDir + config.GetRootFolderOrDefault())
	if err != nil {
		return err
	}
	err = fileutils.CreateIfNotExists(currentDir + config.GetRootFolderOrDefault() + "/" + projectName)
	if err != nil {
		return err
	}
	return nil
}
