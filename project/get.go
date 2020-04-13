package project

import (
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

func NewProjectFromNamespace(projectName string, namespace string, restConf *rest.Config, options v12.ListOptions) (*OpenshiftProject, error) {
	var items []OpenshiftItemInterface
	// Add All
	newItems, err := ListAll(namespace, restConf, options)
	if err != nil {
		return nil, err
	}
	items = append(items, newItems...)

	// Create Project
	newOp := NewOpenshiftProject(projectName)
	newOp.Items = items
	return newOp, nil
}
