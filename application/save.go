package application

import (
	"fmt"
	"github.com/kgysu/oc-wrapper/fileutils"
	"io"
)

func (app *Application) Save(w io.Writer, rootFolder string) {
	dir, err := fileutils.GetCurrentDir()
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	// create Folders if not exists
	err = checkAppPathStructure(dir+rootFolder, app.Name)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	// write Item files
	for _, item := range app.Items {
		filePath := dir + rootFolder + "/" + app.Name + "/" + item.GetFileName()
		err := item.WriteToFile(filePath)
		if err != nil {
			fmt.Fprintln(w, err.Error())
		}
	}
	fmt.Fprintf(w, "saved app [%s] in [%s] \n", app.Name, rootFolder)
}

func checkAppPathStructure(path string, appName string) error {
	err := fileutils.CreateIfNotExists(path)
	if err != nil {
		return err
	}
	err = fileutils.CreateIfNotExists(path + "/" + appName)
	if err != nil {
		return err
	}
	return nil
}
