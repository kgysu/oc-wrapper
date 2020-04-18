package main

import (
	"fmt"
	"github.com/kgysu/oc-wrapper/application"
	v3 "github.com/kgysu/oc-wrapper/client"
)

// Loads sample app from cluster/namespace
func main() {
	restConf, err := v3.GetRestConfig(false)
	if err != nil {
		panic(err)
	}
	namespace := v3.GetNamespace(false, "")
	appName := "sample"
	labelSelector := "app=sample"
	appFromNamespace, err := application.NewAppFromNamespace(appName, namespace, restConf, labelSelector)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Found project [%s] %d Items \n", appFromNamespace.Name, len(appFromNamespace.Items))
}
