package appitem

import (
	"github.com/kgysu/oc-wrapper/client"
	"github.com/kgysu/oc-wrapper/items"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

func ListAll(namespace string, restConf *rest.Config, options v12.ListOptions) ([]AppItem, error) {
	var resultItems []AppItem
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

	svas, err := ListServiceAccounts(namespace, restConf, options)
	if err != nil {
		return nil, err
	}
	resultItems = append(resultItems, svas...)

	sets, err := ListStatefulSets(namespace, restConf, options)
	if err != nil {
		return nil, err
	}
	resultItems = append(resultItems, sets...)

	roles, err := ListRoles(namespace, restConf, options)
	if err != nil {
		return nil, err
	}
	resultItems = append(resultItems, roles...)

	roleBindings, err := ListRoleBindings(namespace, restConf, options)
	if err != nil {
		return nil, err
	}
	resultItems = append(resultItems, roleBindings...)

	configMaps, err := ListConfigMaps(namespace, restConf, options)
	if err != nil {
		return nil, err
	}
	resultItems = append(resultItems, configMaps...)

	return resultItems, nil
}

// Todo add more Types
func ListDeploymentConfigs(namespace string, restConf *rest.Config, options v12.ListOptions) ([]AppItem, error) {
	api, err := client.GetDeploymentConfigsInterface(namespace, restConf)
	if err != nil {
		return nil, err
	}
	list, err := api.List(options)
	if err != nil {
		return nil, err
	}
	var resultItems []AppItem
	for _, it := range list.Items {
		resultItems = append(resultItems, items.NewOpDeploymentConfig(it))
	}
	return resultItems, nil
}

func ListServices(namespace string, restConf *rest.Config, options v12.ListOptions) ([]AppItem, error) {
	api, err := client.GetServicesInterface(namespace, restConf)
	if err != nil {
		return nil, err
	}
	list, err := api.List(options)
	if err != nil {
		return nil, err
	}
	var resultItems []AppItem
	for _, it := range list.Items {
		resultItems = append(resultItems, items.NewOpService(it))
	}
	return resultItems, nil
}

func ListRoutes(namespace string, restConf *rest.Config, options v12.ListOptions) ([]AppItem, error) {
	api, err := client.GetRoutesInterface(namespace, restConf)
	if err != nil {
		return nil, err
	}
	list, err := api.List(options)
	if err != nil {
		return nil, err
	}
	var resultItems []AppItem
	for _, it := range list.Items {
		resultItems = append(resultItems, items.NewOpRoute(it))
	}
	return resultItems, nil
}

func ListPods(namespace string, restConf *rest.Config, options v12.ListOptions) ([]AppItem, error) {
	api, err := client.GetPodsInterface(namespace, restConf)
	if err != nil {
		return nil, err
	}
	list, err := api.List(options)
	if err != nil {
		return nil, err
	}
	var resultItems []AppItem
	for _, it := range list.Items {
		resultItems = append(resultItems, items.NewOpPod(it))
	}
	return resultItems, nil
}

func ListStatefulSets(namespace string, restConf *rest.Config, options v12.ListOptions) ([]AppItem, error) {
	api, err := client.GetStatefulSetsInterface(namespace, restConf)
	if err != nil {
		return nil, err
	}
	list, err := api.List(options)
	if err != nil {
		return nil, err
	}
	var resultItems []AppItem
	for _, it := range list.Items {
		resultItems = append(resultItems, items.NewOpStatefulSet(it))
	}
	return resultItems, nil
}

func ListServiceAccounts(namespace string, restConf *rest.Config, options v12.ListOptions) ([]AppItem, error) {
	api, err := client.GetServiceAccountsInterface(namespace, restConf)
	if err != nil {
		return nil, err
	}
	list, err := api.List(options)
	if err != nil {
		return nil, err
	}
	var resultItems []AppItem
	for _, it := range list.Items {
		resultItems = append(resultItems, items.NewOpServiceAccount(it))
	}
	return resultItems, nil
}

func ListRoles(namespace string, restConf *rest.Config, options v12.ListOptions) ([]AppItem, error) {
	api, err := client.GetRolesInterface(namespace, restConf)
	if err != nil {
		return nil, err
	}
	list, err := api.List(options)
	if err != nil {
		return nil, err
	}
	var resultItems []AppItem
	for _, it := range list.Items {
		resultItems = append(resultItems, items.NewOpRole(it))
	}
	return resultItems, nil
}

func ListRoleBindings(namespace string, restConf *rest.Config, options v12.ListOptions) ([]AppItem, error) {
	api, err := client.GetRoleBindingsInterface(namespace, restConf)
	if err != nil {
		return nil, err
	}
	list, err := api.List(options)
	if err != nil {
		return nil, err
	}
	var resultItems []AppItem
	for _, it := range list.Items {
		resultItems = append(resultItems, items.NewOpRoleBinding(it))
	}
	return resultItems, nil
}

func ListConfigMaps(namespace string, restConf *rest.Config, options v12.ListOptions) ([]AppItem, error) {
	api, err := client.GetConfigMapsInterface(namespace, restConf)
	if err != nil {
		return nil, err
	}
	list, err := api.List(options)
	if err != nil {
		return nil, err
	}
	var resultItems []AppItem
	for _, it := range list.Items {
		resultItems = append(resultItems, items.NewOpConfigMap(it))
	}
	return resultItems, nil
}
