package v2

import (
	"fmt"
	v1 "github.com/openshift/api/apps/v1"
	v18 "github.com/openshift/api/route/v1"
	v12 "github.com/openshift/client-go/apps/clientset/versioned/typed/apps/v1"
	v17 "github.com/openshift/client-go/route/clientset/versioned/typed/route/v1"
	v15 "k8s.io/api/apps/v1"
	v16 "k8s.io/api/core/v1"
	v19 "k8s.io/api/rbac/v1"
	v14 "k8s.io/client-go/kubernetes/typed/apps/v1"
	v13 "k8s.io/client-go/kubernetes/typed/core/v1"
	rbacv1client "k8s.io/client-go/kubernetes/typed/rbac/v1"
)

// Create
func CreateAllItemsRemote(namespace string, items []OpenshiftItem) ([]OpenshiftItem, error) {
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
			createDeploymentConfig(namespace, appsClient, item, &resultItems)
		}
		if item.kind == StatefulSetKey {
			createStatefulSet(namespace, kubeClient, item, &resultItems)
		}
		if item.kind == ServiceKey {
			createService(namespace, coreClient, item, &resultItems)
		}
		if item.kind == ServiceAccountKey {
			createServiceAccount(namespace, coreClient, item, &resultItems)
		}
		if item.kind == SecretKey {
			createSecret(namespace, coreClient, item, &resultItems)
		}
		if item.kind == RouteKey {
			createRoute(namespace, routeClient, item, &resultItems)
		}
		if item.kind == ConfigMapKey {
			createConfigMap(namespace, coreClient, item, &resultItems)
		}
		if item.kind == RoleKey {
			createRole(namespace, rbacClient, item, &resultItems)
		}
		if item.kind == RoleBindingKey {
			createRoleBinding(namespace, rbacClient, item, &resultItems)
		}
	}

	return resultItems, nil
}

func createDeploymentConfig(namespace string, client *v12.AppsV1Client, item OpenshiftItem, resultItems *[]OpenshiftItem) {
	var realItem v1.DeploymentConfig
	err := item.ParseTo(realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	created, err := client.DeploymentConfigs(namespace).Create(&realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	*resultItems = append(*resultItems, NewOpenshiftItem(created.Name, DeploymentConfigKey, created.String()))
}

func createStatefulSet(namespace string, client v14.AppsV1Interface, item OpenshiftItem, resultItems *[]OpenshiftItem) {
	var realItem v15.StatefulSet
	err := item.ParseTo(realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	created, err := client.StatefulSets(namespace).Create(&realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	*resultItems = append(*resultItems, NewOpenshiftItem(created.Name, StatefulSetKey, created.String()))
}

func createService(namespace string, client *v13.CoreV1Client, item OpenshiftItem, resultItems *[]OpenshiftItem) {
	var realItem v16.Service
	err := item.ParseTo(realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	created, err := client.Services(namespace).Create(&realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	*resultItems = append(*resultItems, NewOpenshiftItem(created.Name, ServiceKey, created.String()))
}

func createServiceAccount(namespace string, client *v13.CoreV1Client, item OpenshiftItem, resultItems *[]OpenshiftItem) {
	var realItem v16.ServiceAccount
	err := item.ParseTo(realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	created, err := client.ServiceAccounts(namespace).Create(&realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	*resultItems = append(*resultItems, NewOpenshiftItem(created.Name, ServiceAccountKey, created.String()))
}

func createSecret(namespace string, client *v13.CoreV1Client, item OpenshiftItem, resultItems *[]OpenshiftItem) {
	var realItem v16.Secret
	err := item.ParseTo(realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	created, err := client.Secrets(namespace).Create(&realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	*resultItems = append(*resultItems, NewOpenshiftItem(created.Name, SecretKey, created.String()))
}

func createConfigMap(namespace string, client *v13.CoreV1Client, item OpenshiftItem, resultItems *[]OpenshiftItem) {
	var realItem v16.ConfigMap
	err := item.ParseTo(realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	created, err := client.ConfigMaps(namespace).Create(&realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	*resultItems = append(*resultItems, NewOpenshiftItem(created.Name, ConfigMapKey, created.String()))
}

func createRoute(namespace string, client *v17.RouteV1Client, item OpenshiftItem, resultItems *[]OpenshiftItem) {
	var realItem v18.Route
	err := item.ParseTo(realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	created, err := client.Routes(namespace).Create(&realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	*resultItems = append(*resultItems, NewOpenshiftItem(created.Name, RouteKey, created.String()))
}

func createRole(namespace string, client *rbacv1client.RbacV1Client, item OpenshiftItem, resultItems *[]OpenshiftItem) {
	var realItem v19.Role
	err := item.ParseTo(realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	created, err := client.Roles(namespace).Create(&realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	*resultItems = append(*resultItems, NewOpenshiftItem(created.Name, RouteKey, created.String()))
}

func createRoleBinding(namespace string, client *rbacv1client.RbacV1Client, item OpenshiftItem, resultItems *[]OpenshiftItem) {
	var realItem v19.RoleBinding
	err := item.ParseTo(realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	created, err := client.RoleBindings(namespace).Create(&realItem)
	if err != nil {
		onlyLogOnError(err)
		return
	}
	*resultItems = append(*resultItems, NewOpenshiftItem(created.Name, RouteKey, created.String()))
}
