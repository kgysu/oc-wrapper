package main

import (
	"fmt"
	"github.com/kgysu/oc-wrapper/application"
	"github.com/kgysu/oc-wrapper/fileutils"
)

// Loads the created sample app from disk
func main() {
	currentDir, err := fileutils.GetCurrentDir()
	if err != nil {
		panic(err)
	}
	appsFolder := currentDir + "/apps"
	namespace := "default"
	appsFromDisk, err := application.NewAppsFromDisk(appsFolder, namespace)
	if err != nil {
		panic(err)
	}
	for _, app := range appsFromDisk {
		fmt.Printf("found app [%s] \n", app.Name)
		for i, item := range app.Items {
			fmt.Printf("found item %d: %s \n", i, item.Info())
		}
	}
}
