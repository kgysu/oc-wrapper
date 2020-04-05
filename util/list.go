package util

import (
	"github.com/kgysu/oc-wrapper/items"
	"github.com/kgysu/oc-wrapper/project"
	"github.com/kgysu/oc-wrapper/v3"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

func ListAll(namespace string, restConf *rest.Config, options v12.ListOptions) ([]project.OpenshiftItemInterface, error) {
	var resultItems []project.OpenshiftItemInterface
	dcs, err := ListDeploymentConfigs(namespace, restConf, options)
	if err != nil {
		return nil, err
	}
	resultItems = append(resultItems, dcs...)

	svcs, err := ListServices(namespace, restConf, options)
	if err != nil {
		return nil, err
	}
	resultItems = append(resultItems, svcs...)

	routes, err := ListRoutes(namespace, restConf, options)
	if err != nil {
		return nil, err
	}
	resultItems = append(resultItems, routes...)

	return resultItems, nil
}

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
		resultItems = append(resultItems, items.NewOpDeploymentConfig(it))
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
		resultItems = append(resultItems, items.NewOpService(it))
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
		resultItems = append(resultItems, items.NewOpRoute(it))
	}
	return resultItems, nil
}