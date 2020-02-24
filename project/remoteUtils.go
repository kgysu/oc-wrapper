package project

import (
	"github.com/kgysu/oc-wrapper/wrapper"
	v14 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

func loadAllItemsFromServer(namespace string, options v14.ListOptions) ([]OpenshiftItem, []error) {
	return loadItemsByTypeFromServer(namespace, AllItemsKey, options)
}

func loadItemsByTypeFromServer(namespace string, itemTypes string, options v14.ListOptions) ([]OpenshiftItem, []error) {
	var items []OpenshiftItem
	var errs []error

	if isItemTypeEqualOrAll(itemTypes, ConfigMapKey) {
		configMaps, err := wrapper.ListConfigMaps(namespace, options)
		if err != nil {
			errs = append(errs, err)
		}
		for _, configMap := range configMaps {
			items = append(items, fromConfigMap(configMap))
		}
	}
	if isItemTypeEqualOrAll(itemTypes, DeploymentConfigKey) {
		deploymentConfigs, err := wrapper.ListDeploymentConfigs(namespace, options)
		if err != nil {
			errs = append(errs, err)
		}
		for _, dc := range deploymentConfigs {
			items = append(items, fromDeploymentConfig(dc))
		}
	}
	if isItemTypeEqualOrAll(itemTypes, EventKey) {
		events, err := wrapper.ListEvents(namespace, options)
		if err != nil {
			errs = append(errs, err)
		}
		for _, event := range events {
			items = append(items, fromEvent(event))
		}
	}
	if isItemTypeEqualOrAll(itemTypes, KubeStatefulSetKey) {
		kubeStatefulSets, err := wrapper.ListKubeStatefulSets(namespace, options)
		if err != nil {
			errs = append(errs, err)
		}
		for _, kubeStatefulSet := range kubeStatefulSets {
			items = append(items, fromKubeStatefulSet(kubeStatefulSet))
		}
	}
	if isItemTypeEqualOrAll(itemTypes, PodKey) {
		pods, err := wrapper.ListPods(namespace, options)
		if err != nil {
			errs = append(errs, err)
		}
		for _, pod := range pods {
			items = append(items, fromPod(pod))
		}
	}
	if isItemTypeEqualOrAll(itemTypes, PvKey) {
		pvs, err := wrapper.ListPersistentVolumes(namespace, options)
		if err != nil {
			errs = append(errs, err)
		}
		for _, pv := range pvs {
			items = append(items, fromPersistentVolume(pv))
		}
	}
	if isItemTypeEqualOrAll(itemTypes, PvClaimKey) {
		pvClaims, err := wrapper.ListPersistentVolumeClaims(namespace, options)
		if err != nil {
			errs = append(errs, err)
		}
		for _, pvClaim := range pvClaims {
			items = append(items, fromPersistentVolumeClaim(pvClaim))
		}
	}
	if isItemTypeEqualOrAll(itemTypes, ReplicationControllerKey) {
		rcs, err := wrapper.ListReplicationControllers(namespace, options)
		if err != nil {
			errs = append(errs, err)
		}
		for _, rc := range rcs {
			items = append(items, fromReplicationController(rc))
		}
	}
	if isItemTypeEqualOrAll(itemTypes, RoleKey) {
		roles, err := wrapper.ListRoles(namespace, options)
		if err != nil {
			errs = append(errs, err)
		}
		for _, role := range roles {
			items = append(items, fromRole(role))
		}
	}
	if isItemTypeEqualOrAll(itemTypes, RoleBindingKey) {
		roleBindings, err := wrapper.ListRoleBindings(namespace, options)
		if err != nil {
			errs = append(errs, err)
		}
		for _, roleBinding := range roleBindings {
			items = append(items, fromRoleBinding(roleBinding))
		}
	}
	if isItemTypeEqualOrAll(itemTypes, RouteKey) {
		routes, err := wrapper.ListRoutes(namespace, options)
		if err != nil {
			errs = append(errs, err)
		}
		for _, route := range routes {
			items = append(items, fromRoute(route))
		}
	}
	if isItemTypeEqualOrAll(itemTypes, SecretKey) {
		secrets, err := wrapper.ListSecrets(namespace, options)
		if err != nil {
			errs = append(errs, err)
		}
		for _, secret := range secrets {
			items = append(items, fromSecret(secret))
		}
	}
	if isItemTypeEqualOrAll(itemTypes, ServiceKey) {
		services, err := wrapper.ListServices(namespace, options)
		if err != nil {
			errs = append(errs, err)
		}
		for _, service := range services {
			items = append(items, fromService(service))
		}
	}
	if isItemTypeEqualOrAll(itemTypes, ServiceAccountKey) {
		serviceAccounts, err := wrapper.ListServiceAccounts(namespace, options)
		if err != nil {
			errs = append(errs, err)
		}
		for _, serviceAccount := range serviceAccounts {
			items = append(items, fromServiceAccount(serviceAccount))
		}
	}
	return items, errs
}

func isItemTypeEqualOrAll(itemTypes string, shouldEqual string) bool {
	return itemTypes == AllItemsKey || strings.Contains(itemTypes, shouldEqual)
}
