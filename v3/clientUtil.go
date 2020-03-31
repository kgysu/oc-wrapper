package v3

import (
	"fmt"
	appsv1client "github.com/openshift/client-go/apps/clientset/versioned/typed/apps/v1"
	routev1client "github.com/openshift/client-go/route/clientset/versioned/typed/route/v1"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	rbacv1client "k8s.io/client-go/kubernetes/typed/rbac/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

func GetAppsV1Client(restConf *rest.Config) (*appsv1client.AppsV1Client, error) {
	client, err := appsv1client.NewForConfig(restConf)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetCoreV1Client(restConf *rest.Config) (*corev1client.CoreV1Client, error) {
	client, err := corev1client.NewForConfig(restConf)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetRouteV1Client(restConf *rest.Config) (*routev1client.RouteV1Client, error) {
	client, err := routev1client.NewForConfig(restConf)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetRbacV1Client(restConf *rest.Config) (*rbacv1client.RbacV1Client, error) {
	client, err := rbacv1client.NewForConfig(restConf)
	if err != nil {
		return nil, err
	}
	return client, nil
}

const NamespaceEnvName = "NAMESPACE"

func GetNamespace(fromEnv bool) string {
	if fromEnv {
		namespace := os.Getenv(NamespaceEnvName)
		if namespace == "" {
			panic("NAMESPACE was not set")
		}
		return namespace
	} else {
		kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
			clientcmd.NewDefaultClientConfigLoadingRules(),
			&clientcmd.ConfigOverrides{},
		)

		namespace, _, err := kubeconfig.Namespace()
		if err != nil {
			fmt.Println(err.Error())
		}
		if namespace == "" {
			panic("NAMESPACE was not set")
		}
		return namespace
	}
}

func GetRestConfig(inCluster bool) (*rest.Config, error) {
	if inCluster {
		// Build a rest.Config from configuration injected into the Pod by
		// Kubernetes.  Clients will use the Pod's ServiceAccount principal.
		restconfig, err := rest.InClusterConfig()
		if err != nil {
			return nil, err
		}
		return restconfig, nil
	} else {
		// Instantiate loader for kubeconfig file.
		kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
			clientcmd.NewDefaultClientConfigLoadingRules(),
			&clientcmd.ConfigOverrides{},
		)
		restConfig, err := kubeconfig.ClientConfig()
		if err != nil {
			return nil, err
		}
		return restConfig, nil
	}
}
