package application

import (
	"github.com/kgysu/oc-wrapper/appitem"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

func NewAppFromNamespace(appName string, namespace string, restConf *rest.Config, options v12.ListOptions) (*Application, error) {
	var items []appitem.AppItem
	// Add All
	newItems, err := appitem.ListAll(namespace, restConf, options)
	if err != nil {
		return nil, err
	}
	items = append(items, newItems...)

	// Create App
	newOp := NewApp(appName)
	newOp.Items = items
	return newOp, nil
}
