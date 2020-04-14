package application

import (
	"github.com/kgysu/oc-wrapper/appitem"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

func NewAppFromNamespace(appName string, namespace string, restConf *rest.Config, labelSelector string) (*Application, error) {
	var items []appitem.AppItem
	// Add All
	newItems, err := appitem.ListAll(namespace, restConf, v12.ListOptions{
		LabelSelector: labelSelector,
		Limit:         0,
	})
	if err != nil {
		return nil, err
	}
	items = append(items, newItems...)

	// Create App
	newOp := NewAppWithLabel(appName, labelSelector)
	newOp.Items = items
	return newOp, nil
}
