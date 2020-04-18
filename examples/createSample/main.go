package main

import (
	"github.com/kgysu/oc-wrapper/templates"
	"os"
)

// Creates a sample app in current directory's apps folder
func main() {
	appsFolder := "/apps"
	sampleProject := templates.GetSampleApp()
	sampleProject.Save(os.Stdout, appsFolder)
}
