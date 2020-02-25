package v2

import (
	"fmt"
	v14 "github.com/openshift/client-go/apps/clientset/versioned/typed/apps/v1"
	v1 "github.com/openshift/client-go/route/clientset/versioned/typed/route/v1"
	v13 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v12 "k8s.io/client-go/kubernetes/typed/apps/v1"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	rbacv1client "k8s.io/client-go/kubernetes/typed/rbac/v1"
)

// List
func CreateOrUpdateAllRemoteItems(namespace string, items []OpenshiftItem) ([]OpenshiftItem, error) {
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
		fmt.Printf("Update %s: %s", item.kind, item.name)
		if item.kind == DeploymentConfigKey {
			createOrUpdateDeploymentConfig(namespace, appsClient, item, &resultItems)
		}
		if item.kind == StatefulSetKey {
			createOrUpdateStatefulSet(namespace, kubeClient, item, &resultItems)
		}
		if item.kind == ServiceKey {
			createOrUpdateService(namespace, coreClient, item, &resultItems)
		}
		if item.kind == ServiceAccountKey {
			createOrUpdateServiceAccount(namespace, coreClient, item, &resultItems)
		}
		if item.kind == SecretKey {
			createOrUpdateSecret(namespace, coreClient, item, &resultItems)
		}
		if item.kind == ConfigMapKey {
			createOrUpdateConfigMap(namespace, coreClient, item, &resultItems)
		}
		if item.kind == RouteKey {
			createOrUpdateRoute(namespace, routeClient, item, &resultItems)
		}
		if item.kind == RoleKey {
			createOrUpdateRole(namespace, rbacClient, item, &resultItems)
		}
		if item.kind == RoleBindingKey {
			createOrUpdateRoleBinding(namespace, rbacClient, item, &resultItems)
		}
	}

	return resultItems, nil
}

func createOrUpdateDeploymentConfig(namespace string, appsClient *v14.AppsV1Client, item OpenshiftItem, resultItems *[]OpenshiftItem) {
	fromRemote, err := appsClient.DeploymentConfigs(namespace).Get(item.name, v13.GetOptions{})
	if err != nil {
		createDeploymentConfig(namespace, appsClient, item, resultItems)
	} else {
		if fromRemote != nil {
			updateDeploymentConfigFromRemote(fromRemote, item, namespace, appsClient, resultItems)
		} else {
			updateDeploymentConfigItem(item, namespace, appsClient, resultItems)
		}
	}
}

func createOrUpdateStatefulSet(namespace string, client v12.AppsV1Interface, item OpenshiftItem, resultItems *[]OpenshiftItem) {
	fromRemote, err := client.StatefulSets(namespace).Get(item.name, v13.GetOptions{})
	if err != nil {
		createStatefulSet(namespace, client, item, resultItems)
	} else {
		if fromRemote != nil {
			updateStatefulSetFromRemote(fromRemote, item, namespace, client, resultItems)
		} else {
			updateStatefulSetItem(item, namespace, client, resultItems)
		}
	}
}

func createOrUpdateService(namespace string, client *corev1client.CoreV1Client, item OpenshiftItem, resultItems *[]OpenshiftItem) {
	fromRemote, err := client.Services(namespace).Get(item.name, v13.GetOptions{})
	if err != nil {
		createService(namespace, client, item, resultItems)
	} else {
		if fromRemote != nil {
			updateServiceFromRemote(fromRemote, item, namespace, client, resultItems)
		} else {
			updateServiceItem(item, namespace, client, resultItems)
		}
	}
}

func createOrUpdateServiceAccount(namespace string, client *corev1client.CoreV1Client, item OpenshiftItem, resultItems *[]OpenshiftItem) {
	fromRemote, err := client.ServiceAccounts(namespace).Get(item.name, v13.GetOptions{})
	if err != nil {
		createServiceAccount(namespace, client, item, resultItems)
	} else {
		if fromRemote != nil {
			updateServiceAccountFromRemote(fromRemote, item, namespace, client, resultItems)
		} else {
			updateServiceAccountItem(item, namespace, client, resultItems)
		}
	}
}

func createOrUpdateSecret(namespace string, client *corev1client.CoreV1Client, item OpenshiftItem, resultItems *[]OpenshiftItem) {
	fromRemote, err := client.Secrets(namespace).Get(item.name, v13.GetOptions{})
	if err != nil {
		createSecret(namespace, client, item, resultItems)
	} else {
		if fromRemote != nil {
			updateSecretFromRemote(fromRemote, item, namespace, client, resultItems)
		} else {
			updateSecretItem(item, namespace, client, resultItems)
		}
	}
}

func createOrUpdateConfigMap(namespace string, client *corev1client.CoreV1Client, item OpenshiftItem, resultItems *[]OpenshiftItem) {
	fromRemote, err := client.ConfigMaps(namespace).Get(item.name, v13.GetOptions{})
	if err != nil {
		createConfigMap(namespace, client, item, resultItems)
	} else {
		if fromRemote != nil {
			updateConfigMapFromRemote(fromRemote, item, namespace, client, resultItems)
		} else {
			updateConfigMapItem(item, namespace, client, resultItems)
		}
	}
}

func createOrUpdateRoute(namespace string, client *v1.RouteV1Client, item OpenshiftItem, resultItems *[]OpenshiftItem) {
	fromRemote, err := client.Routes(namespace).Get(item.name, v13.GetOptions{})
	if err != nil {
		createRoute(namespace, client, item, resultItems)
	} else {
		if fromRemote != nil {
			updateRouteFromRemote(fromRemote, item, namespace, client, resultItems)
		} else {
			updateRouteItem(item, namespace, client, resultItems)
		}
	}
}

func createOrUpdateRole(namespace string, client *rbacv1client.RbacV1Client, item OpenshiftItem, resultItems *[]OpenshiftItem) {
	fromRemote, err := client.Roles(namespace).Get(item.name, v13.GetOptions{})
	if err != nil {
		createRole(namespace, client, item, resultItems)
	} else {
		if fromRemote != nil {
			updateRoleFromRemote(fromRemote, item, namespace, client, resultItems)
		} else {
			updateRoleItem(item, namespace, client, resultItems)
		}
	}
}

func createOrUpdateRoleBinding(namespace string, client *rbacv1client.RbacV1Client, item OpenshiftItem, resultItems *[]OpenshiftItem) {
	fromRemote, err := client.RoleBindings(namespace).Get(item.name, v13.GetOptions{})
	if err != nil {
		createRoleBinding(namespace, client, item, resultItems)
	} else {
		if fromRemote != nil {
			updateRoleBindingFromRemote(fromRemote, item, namespace, client, resultItems)
		} else {
			updateRoleBindingItem(item, namespace, client, resultItems)
		}
	}
}
