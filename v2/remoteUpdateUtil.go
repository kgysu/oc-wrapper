package v2

import (
	"fmt"
	v1 "github.com/openshift/api/apps/v1"
	v17 "github.com/openshift/api/route/v1"
	v12 "github.com/openshift/client-go/apps/clientset/versioned/typed/apps/v1"
	v18 "github.com/openshift/client-go/route/clientset/versioned/typed/route/v1"
	v13 "k8s.io/api/apps/v1"
	v15 "k8s.io/api/core/v1"
	v19 "k8s.io/api/rbac/v1"
	v14 "k8s.io/client-go/kubernetes/typed/apps/v1"
	v16 "k8s.io/client-go/kubernetes/typed/core/v1"
	v110 "k8s.io/client-go/kubernetes/typed/rbac/v1"
)

// Update
func UpdateAllItemsRemote(namespace string, items []OpenshiftItem) ([]OpenshiftItem, error) {
	kubeClient, err := GetKubeAppsV1Client()
	if err != nil {
		return nil, err
	}
	appsClient, err := GetAppsV1Client()
	if err != nil {
		return nil, err
	}
	coreClient, err := GetCoreV1Client()
	if err != nil {
		return nil, err
	}
	routeClient, err := GetRouteV1Client()
	if err != nil {
		return nil, err
	}
	rbacClient, err := GetRbacV1Client()
	if err != nil {
		return nil, err
	}
	var resultItems []OpenshiftItem

	for _, item := range items {
		fmt.Printf("Create %s: %s", item.kind, item.name)
		if item.kind == DeploymentConfigKey {
			updateDeploymentConfigItem(item, namespace, appsClient, &resultItems)
		}
		if item.kind == StatefulSetKey {
			updateStatefulSetItem(item, namespace, kubeClient, &resultItems)
		}
		if item.kind == ServiceKey {
			updateServiceItem(item, namespace, coreClient, &resultItems)
		}
		if item.kind == ServiceAccountKey {
			updateServiceAccountItem(item, namespace, coreClient, &resultItems)
		}
		if item.kind == SecretKey {
			updateSecretItem(item, namespace, coreClient, &resultItems)
		}
		if item.kind == ConfigMapKey {
			updateConfigMapItem(item, namespace, coreClient, &resultItems)
		}
		if item.kind == RouteKey {
			updateRouteItem(item, namespace, routeClient, &resultItems)
		}
		if item.kind == RoleKey {
			updateRoleItem(item, namespace, rbacClient, &resultItems)
		}
		if item.kind == RoleBindingKey {
			updateRoleBindingItem(item, namespace, rbacClient, &resultItems)
		}
	}

	return resultItems, nil
}

// DeploymentConfig
func updateDeploymentConfigItem(item OpenshiftItem, namespace string, appsClient *v12.AppsV1Client, resultItems *[]OpenshiftItem) {
	var realItem v1.DeploymentConfig
	err := item.ParseTo(&realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	updateDeploymentConfig(&realItem, namespace, appsClient, resultItems)
}
func updateDeploymentConfigFromRemote(fromRemote *v1.DeploymentConfig, item OpenshiftItem, namespace string, appsClient *v12.AppsV1Client, resultItems *[]OpenshiftItem) {
	var realItem v1.DeploymentConfig
	err := item.ParseTo(&realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	if fromRemote != nil {
		fromRemote.Name = realItem.Name
		fromRemote.Labels = realItem.Labels
		fromRemote.Annotations = realItem.Annotations
		updateDeploymentConfig(fromRemote, namespace, appsClient, resultItems)
	} else {
		updateDeploymentConfig(&realItem, namespace, appsClient, resultItems)
	}
}
func updateDeploymentConfig(realItem *v1.DeploymentConfig, namespace string, appsClient *v12.AppsV1Client, resultItems *[]OpenshiftItem) {
	updated, err := appsClient.DeploymentConfigs(namespace).Update(realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	*resultItems = append(*resultItems, NewOpenshiftItem(updated.Name, DeploymentConfigKey, updated.String()))
}

// StatefulSet
func updateStatefulSetItem(item OpenshiftItem, namespace string, appsClient v14.AppsV1Interface, resultItems *[]OpenshiftItem) {
	var realItem v13.StatefulSet
	err := item.ParseTo(&realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	updateStatefulSet(&realItem, namespace, appsClient, resultItems)
}
func updateStatefulSetFromRemote(fromRemote *v13.StatefulSet, item OpenshiftItem, namespace string, appsClient v14.AppsV1Interface, resultItems *[]OpenshiftItem) {
	var realItem v13.StatefulSet
	err := item.ParseTo(&realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	if fromRemote != nil {
		fromRemote.Name = realItem.Name
		fromRemote.Labels = realItem.Labels
		fromRemote.Annotations = realItem.Annotations
		updateStatefulSet(fromRemote, namespace, appsClient, resultItems)
	} else {
		updateStatefulSet(&realItem, namespace, appsClient, resultItems)
	}
}
func updateStatefulSet(realItem *v13.StatefulSet, namespace string, appsClient v14.AppsV1Interface, resultItems *[]OpenshiftItem) {
	updated, err := appsClient.StatefulSets(namespace).Update(realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	*resultItems = append(*resultItems, NewOpenshiftItem(updated.Name, StatefulSetKey, updated.String()))
}

// Service
func updateServiceItem(item OpenshiftItem, namespace string, client *v16.CoreV1Client, resultItems *[]OpenshiftItem) {
	var realItem v15.Service
	err := item.ParseTo(&realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	updateService(&realItem, namespace, client, resultItems)
}
func updateServiceFromRemote(fromRemote *v15.Service, item OpenshiftItem, namespace string, client *v16.CoreV1Client, resultItems *[]OpenshiftItem) {
	var realItem v15.Service
	err := item.ParseTo(&realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	if fromRemote != nil {
		fromRemote.Name = realItem.Name
		fromRemote.Labels = realItem.Labels
		fromRemote.Annotations = realItem.Annotations
		updateService(fromRemote, namespace, client, resultItems)
	} else {
		updateService(&realItem, namespace, client, resultItems)
	}
}
func updateService(realItem *v15.Service, namespace string, client *v16.CoreV1Client, resultItems *[]OpenshiftItem) {
	updated, err := client.Services(namespace).Update(realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	*resultItems = append(*resultItems, NewOpenshiftItem(updated.Name, ServiceKey, updated.String()))
}

// ServiceAccount
func updateServiceAccountItem(item OpenshiftItem, namespace string, client *v16.CoreV1Client, resultItems *[]OpenshiftItem) {
	var realItem v15.ServiceAccount
	err := item.ParseTo(&realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	updateServiceAccount(&realItem, namespace, client, resultItems)
}
func updateServiceAccountFromRemote(fromRemote *v15.ServiceAccount, item OpenshiftItem, namespace string, client *v16.CoreV1Client, resultItems *[]OpenshiftItem) {
	var realItem v15.ServiceAccount
	err := item.ParseTo(&realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	if fromRemote != nil {
		fromRemote.Name = realItem.Name
		fromRemote.Labels = realItem.Labels
		fromRemote.Annotations = realItem.Annotations
		updateServiceAccount(fromRemote, namespace, client, resultItems)
	} else {
		updateServiceAccount(&realItem, namespace, client, resultItems)
	}
}
func updateServiceAccount(realItem *v15.ServiceAccount, namespace string, client *v16.CoreV1Client, resultItems *[]OpenshiftItem) {
	updated, err := client.ServiceAccounts(namespace).Update(realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	*resultItems = append(*resultItems, NewOpenshiftItem(updated.Name, ServiceAccountKey, updated.String()))
}

// Secret
func updateSecretItem(item OpenshiftItem, namespace string, client *v16.CoreV1Client, resultItems *[]OpenshiftItem) {
	var realItem v15.Secret
	err := item.ParseTo(&realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	updateSecret(&realItem, namespace, client, resultItems)
}
func updateSecretFromRemote(fromRemote *v15.Secret, item OpenshiftItem, namespace string, client *v16.CoreV1Client, resultItems *[]OpenshiftItem) {
	var realItem v15.Secret
	err := item.ParseTo(&realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	if fromRemote != nil {
		fromRemote.Name = realItem.Name
		fromRemote.Labels = realItem.Labels
		fromRemote.Annotations = realItem.Annotations
		updateSecret(fromRemote, namespace, client, resultItems)
	} else {
		updateSecret(&realItem, namespace, client, resultItems)
	}
}
func updateSecret(realItem *v15.Secret, namespace string, client *v16.CoreV1Client, resultItems *[]OpenshiftItem) {
	updated, err := client.Secrets(namespace).Update(realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	*resultItems = append(*resultItems, NewOpenshiftItem(updated.Name, SecretKey, updated.String()))
}

// ConfigMap
func updateConfigMapItem(item OpenshiftItem, namespace string, client *v16.CoreV1Client, resultItems *[]OpenshiftItem) {
	var realItem v15.ConfigMap
	err := item.ParseTo(&realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	updateConfigMap(&realItem, namespace, client, resultItems)
}
func updateConfigMapFromRemote(fromRemote *v15.ConfigMap, item OpenshiftItem, namespace string, client *v16.CoreV1Client, resultItems *[]OpenshiftItem) {
	var realItem v15.ConfigMap
	err := item.ParseTo(&realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	if fromRemote != nil {
		fromRemote.Name = realItem.Name
		fromRemote.Labels = realItem.Labels
		fromRemote.Annotations = realItem.Annotations
		updateConfigMap(fromRemote, namespace, client, resultItems)
	} else {
		updateConfigMap(&realItem, namespace, client, resultItems)
	}
}
func updateConfigMap(realItem *v15.ConfigMap, namespace string, client *v16.CoreV1Client, resultItems *[]OpenshiftItem) {
	updated, err := client.ConfigMaps(namespace).Update(realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	*resultItems = append(*resultItems, NewOpenshiftItem(updated.Name, ConfigMapKey, updated.String()))
}

// Route
func updateRouteItem(item OpenshiftItem, namespace string, client *v18.RouteV1Client, resultItems *[]OpenshiftItem) {
	var realItem v17.Route
	err := item.ParseTo(&realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	updateRoute(&realItem, namespace, client, resultItems)
}
func updateRouteFromRemote(fromRemote *v17.Route, item OpenshiftItem, namespace string, client *v18.RouteV1Client, resultItems *[]OpenshiftItem) {
	var realItem v17.Route
	err := item.ParseTo(&realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	if fromRemote != nil {
		fromRemote.Name = realItem.Name
		fromRemote.Labels = realItem.Labels
		fromRemote.Annotations = realItem.Annotations
		updateRoute(fromRemote, namespace, client, resultItems)
	} else {
		updateRoute(&realItem, namespace, client, resultItems)
	}
}
func updateRoute(realItem *v17.Route, namespace string, client *v18.RouteV1Client, resultItems *[]OpenshiftItem) {
	updated, err := client.Routes(namespace).Update(realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	*resultItems = append(*resultItems, NewOpenshiftItem(updated.Name, RouteKey, updated.String()))
}

// Role
func updateRoleItem(item OpenshiftItem, namespace string, client *v110.RbacV1Client, resultItems *[]OpenshiftItem) {
	var realItem v19.Role
	err := item.ParseTo(&realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	updateRole(&realItem, namespace, client, resultItems)
}
func updateRoleFromRemote(fromRemote *v19.Role, item OpenshiftItem, namespace string, client *v110.RbacV1Client, resultItems *[]OpenshiftItem) {
	var realItem v19.Role
	err := item.ParseTo(&realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	if fromRemote != nil {
		fromRemote.Name = realItem.Name
		fromRemote.Labels = realItem.Labels
		fromRemote.Annotations = realItem.Annotations
		updateRole(fromRemote, namespace, client, resultItems)
	} else {
		updateRole(&realItem, namespace, client, resultItems)
	}
}
func updateRole(realItem *v19.Role, namespace string, client *v110.RbacV1Client, resultItems *[]OpenshiftItem) {
	updated, err := client.Roles(namespace).Update(realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	*resultItems = append(*resultItems, NewOpenshiftItem(updated.Name, RoleKey, updated.String()))
}

// RoleBinding
func updateRoleBindingItem(item OpenshiftItem, namespace string, client *v110.RbacV1Client, resultItems *[]OpenshiftItem) {
	var realItem v19.RoleBinding
	err := item.ParseTo(&realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	updateRoleBinding(&realItem, namespace, client, resultItems)
}
func updateRoleBindingFromRemote(fromRemote *v19.RoleBinding, item OpenshiftItem, namespace string, client *v110.RbacV1Client, resultItems *[]OpenshiftItem) {
	var realItem v19.RoleBinding
	err := item.ParseTo(&realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	if fromRemote != nil {
		fromRemote.Name = realItem.Name
		fromRemote.Labels = realItem.Labels
		fromRemote.Annotations = realItem.Annotations
		updateRoleBinding(fromRemote, namespace, client, resultItems)
	} else {
		updateRoleBinding(&realItem, namespace, client, resultItems)
	}
}
func updateRoleBinding(realItem *v19.RoleBinding, namespace string, client *v110.RbacV1Client, resultItems *[]OpenshiftItem) {
	updated, err := client.RoleBindings(namespace).Update(realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	*resultItems = append(*resultItems, NewOpenshiftItem(updated.Name, RoleBindingKey, updated.String()))
}
