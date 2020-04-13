package project

import (
	"github.com/kgysu/oc-wrapper/items"
	v3 "github.com/kgysu/oc-wrapper/v3"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

func ListAll(namespace string, restConf *rest.Config, options v12.ListOptions) ([]OpenshiftItemInterface, error) {
	var resultItems []OpenshiftItemInterface
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

	pods, err := ListPods(namespace, restConf, options)
	if err != nil {
		return nil, err
	}
	resultItems = append(resultItems, pods...)

	return resultItems, nil
}

// Todo add more Types
func ListDeploymentConfigs(namespace string, restConf *rest.Config, options v12.ListOptions) ([]OpenshiftItemInterface, error) {
	api, err := v3.GetDeploymentConfigsInterface(namespace, restConf)
	if err != nil {
		return nil, err
	}
	list, err := api.List(options)
	if err != nil {
		return nil, err
	}
	var resultItems []OpenshiftItemInterface
	for _, it := range list.Items {
		resultItems = append(resultItems, items.NewOpDeploymentConfig(it))
	}
	return resultItems, nil
}

func ListServices(namespace string, restConf *rest.Config, options v12.ListOptions) ([]OpenshiftItemInterface, error) {
	api, err := v3.GetServicesInterface(namespace, restConf)
	if err != nil {
		return nil, err
	}
	list, err := api.List(options)
	if err != nil {
		return nil, err
	}
	var resultItems []OpenshiftItemInterface
	for _, it := range list.Items {
		resultItems = append(resultItems, items.NewOpService(it))
	}
	return resultItems, nil
}

func ListRoutes(namespace string, restConf *rest.Config, options v12.ListOptions) ([]OpenshiftItemInterface, error) {
	api, err := v3.GetRoutesInterface(namespace, restConf)
	if err != nil {
		return nil, err
	}
	list, err := api.List(options)
	if err != nil {
		return nil, err
	}
	var resultItems []OpenshiftItemInterface
	for _, it := range list.Items {
		resultItems = append(resultItems, items.NewOpRoute(it))
	}
	return resultItems, nil
}

func ListPods(namespace string, restConf *rest.Config, options v12.ListOptions) ([]OpenshiftItemInterface, error) {
	api, err := v3.GetPodsInterface(namespace, restConf)
	if err != nil {
		return nil, err
	}
	list, err := api.List(options)
	if err != nil {
		return nil, err
	}
	var resultItems []OpenshiftItemInterface
	for _, it := range list.Items {
		resultItems = append(resultItems, items.NewOpPod(it))
	}
	return resultItems, nil
}
