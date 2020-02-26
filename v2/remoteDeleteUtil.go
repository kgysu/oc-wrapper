package v2

import (
	"fmt"
	v14 "github.com/openshift/client-go/apps/clientset/versioned/typed/apps/v1"
	v1 "github.com/openshift/client-go/route/clientset/versioned/typed/route/v1"
	v13 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v15 "k8s.io/client-go/kubernetes/typed/apps/v1"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	v12 "k8s.io/client-go/kubernetes/typed/rbac/v1"
)

// List
func DeleteAllRemoteItems(namespace string, items []OpenshiftItem, options v13.DeleteOptions) error {
	kubeClient, err := GetKubeAppsV1Client()
	if err != nil {
		return err
	}
	appsClient, err := GetAppsV1Client()
	if err != nil {
		return err
	}
	coreClient, err := GetCoreV1Client()
	if err != nil {
		return err
	}
	routeClient, err := GetRouteV1Client()
	if err != nil {
		return err
	}
	rbacClient, err := GetRbacV1Client()
	if err != nil {
		return err
	}

	for _, item := range items {
		fmt.Printf("Delete %s: %s \n", item.kind, item.name)
		if item.kind == DeploymentConfigKey {
			onlyLogOnError(deleteDeploymentConfig(namespace, appsClient, item, options))
		}
		if item.kind == StatefulSetKey {
			onlyLogOnError(deleteStatefulSet(namespace, kubeClient, item, options))
		}

		if item.kind == ServiceKey {
			onlyLogOnError(deleteService(namespace, coreClient, item, options))
		}
		if item.kind == ServiceAccountKey {
			onlyLogOnError(deleteServiceAccount(namespace, coreClient, item, options))
		}
		if item.kind == SecretKey {
			onlyLogOnError(deleteSecret(namespace, coreClient, item, options))
		}
		if item.kind == ConfigMapKey {
			onlyLogOnError(deleteConfigMap(namespace, coreClient, item, options))
		}
		if item.kind == RouteKey {
			onlyLogOnError(deleteRoute(namespace, routeClient, item, options))
		}
		if item.kind == RoleKey {
			onlyLogOnError(deleteRole(namespace, rbacClient, item, options))
		}
		if item.kind == RoleBindingKey {
			onlyLogOnError(deleteRoleBinding(namespace, rbacClient, item, options))
		}
	}
	return nil
}

func deleteDeploymentConfig(namespace string, client *v14.AppsV1Client, item OpenshiftItem, options v13.DeleteOptions) error {
	return client.DeploymentConfigs(namespace).Delete(item.name, &options)
}

func deleteStatefulSet(namespace string, client v15.AppsV1Interface, item OpenshiftItem, options v13.DeleteOptions) error {
	return client.StatefulSets(namespace).Delete(item.name, &options)
}

func deleteService(namespace string, client *corev1client.CoreV1Client, item OpenshiftItem, options v13.DeleteOptions) error {
	return client.Services(namespace).Delete(item.name, &options)
}

func deleteServiceAccount(namespace string, client *corev1client.CoreV1Client, item OpenshiftItem, options v13.DeleteOptions) error {
	return client.ServiceAccounts(namespace).Delete(item.name, &options)
}

func deleteSecret(namespace string, client *corev1client.CoreV1Client, item OpenshiftItem, options v13.DeleteOptions) error {
	return client.Secrets(namespace).Delete(item.name, &options)
}

func deleteConfigMap(namespace string, client *corev1client.CoreV1Client, item OpenshiftItem, options v13.DeleteOptions) error {
	return client.ConfigMaps(namespace).Delete(item.name, &options)
}

func deleteRoute(namespace string, client *v1.RouteV1Client, item OpenshiftItem, options v13.DeleteOptions) error {
	return client.Routes(namespace).Delete(item.name, &options)
}

func deleteRole(namespace string, client *v12.RbacV1Client, item OpenshiftItem, options v13.DeleteOptions) error {
	return client.Roles(namespace).Delete(item.name, &options)
}

func deleteRoleBinding(namespace string, client *v12.RbacV1Client, item OpenshiftItem, options v13.DeleteOptions) error {
	return client.RoleBindings(namespace).Delete(item.name, &options)
}
