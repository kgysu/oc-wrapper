package util

import (
	items2 "github.com/kgysu/oc-wrapper/items"
	"github.com/kgysu/oc-wrapper/project"
	"github.com/kgysu/oc-wrapper/v3"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

// Todo add more Types
func ListDeploymentConfigs(namespace string, restConf *rest.Config, options v12.ListOptions) ([]project.OpenshiftItemInterface, error) {
	api, err := v3.GetDeploymentConfigsInterface(namespace, restConf)
	if err != nil {
		return nil, err
	}
	list, err := api.List(options)
	if err != nil {
		return nil, err
	}
	var resultItems []project.OpenshiftItemInterface
	for _, it := range list.Items {
		resultItems = append(resultItems, items2.NewOpDeploymentConfig(&it))
	}
	return resultItems, nil
}

func ListServices(namespace string, restConf *rest.Config, options v12.ListOptions) ([]project.OpenshiftItemInterface, error) {
	api, err := v3.GetServicesInterface(namespace, restConf)
	if err != nil {
		return nil, err
	}
	list, err := api.List(options)
	if err != nil {
		return nil, err
	}
	var resultItems []project.OpenshiftItemInterface
	for _, it := range list.Items {
		resultItems = append(resultItems, items2.NewOpService(&it))
	}
	return resultItems, nil
}

func ListRoutes(namespace string, restConf *rest.Config, options v12.ListOptions) ([]project.OpenshiftItemInterface, error) {
	api, err := v3.GetRoutesInterface(namespace, restConf)
	if err != nil {
		return nil, err
	}
	list, err := api.List(options)
	if err != nil {
		return nil, err
	}
	var resultItems []project.OpenshiftItemInterface
	for _, it := range list.Items {
		resultItems = append(resultItems, items2.NewOpRoute(&it))
	}
	return resultItems, nil
}
