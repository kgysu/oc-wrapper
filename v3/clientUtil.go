package v3

import (
	appsv1client "github.com/openshift/client-go/apps/clientset/versioned/typed/apps/v1"
	routev1client "github.com/openshift/client-go/route/clientset/versioned/typed/route/v1"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	rbacv1client "k8s.io/client-go/kubernetes/typed/rbac/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func GetAppsV1Client(fromLocal bool) (*appsv1client.AppsV1Client, error) {
	restConfig, err := GetConfig(fromLocal)
	if err != nil {
		return nil, err
	}
	client, err := appsv1client.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetCoreV1Client(fromLocal bool) (*corev1client.CoreV1Client, error) {
	restConfig, err := GetConfig(fromLocal)
	if err != nil {
		return nil, err
	}
	client, err := corev1client.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetRouteV1Client(fromLocal bool) (*routev1client.RouteV1Client, error) {
	restConfig, err := GetConfig(fromLocal)
	if err != nil {
		return nil, err
	}
	client, err := routev1client.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetRbacV1Client(fromLocal bool) (*rbacv1client.RbacV1Client, error) {
	restConfig, err := GetConfig(fromLocal)
	if err != nil {
		return nil, err
	}
	client, err := rbacv1client.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetConfig(fromLocal bool) (*rest.Config, error) {
	if fromLocal {
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
	} else {
		// Build a rest.Config from configuration injected into the Pod by
		// Kubernetes.  Clients will use the Pod's ServiceAccount principal.
		restconfig, err := rest.InClusterConfig()
		if err != nil {
			return nil, err
		}
		return restconfig, nil
	}
}
