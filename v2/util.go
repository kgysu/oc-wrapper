package v2

import (
	"encoding/json"
	"fmt"
	appsv1client "github.com/openshift/client-go/apps/clientset/versioned/typed/apps/v1"
	routev1client "github.com/openshift/client-go/route/clientset/versioned/typed/route/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/typed/apps/v1"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	rbacv1client "k8s.io/client-go/kubernetes/typed/rbac/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

func GetConfigWithinPod() (*corev1client.CoreV1Client, string, error) {
	// Build a rest.Config from configuration injected into the Pod by
	// Kubernetes.  Clients will use the Pod's ServiceAccount principal.
	restconfig, err := rest.InClusterConfig()
	if err != nil {
		return nil, "", nil
	}

	// If you need to know the Pod's Namespace, adjust the Pod's spec to pass
	// the information into an environment variable in advance via the downward
	// API.
	namespace := os.Getenv("NAMESPACE")
	if namespace == "" {
		panic("NAMESPACE was not set")
	}

	// Create a Kubernetes core/v1 client.
	coreclient, err := corev1client.NewForConfig(restconfig)
	if err != nil {
		return nil, "", err
	}

	return coreclient, namespace, nil
}

func GetKubeAppsV1Client() (v1.AppsV1Interface, error) {
	clientSet, err := GetKubeClientSet()
	if err != nil {
		return nil, err
	}
	return clientSet.AppsV1(), nil
}

func GetAppsV1Client() (*appsv1client.AppsV1Client, error) {
	restConfig, err := GetConfigs()
	if err != nil {
		return nil, err
	}
	client, err := appsv1client.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetCoreV1Client() (*corev1client.CoreV1Client, error) {
	restConfig, err := GetConfigs()
	if err != nil {
		return nil, err
	}
	client, err := corev1client.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetRouteV1Client() (*routev1client.RouteV1Client, error) {
	restConfig, err := GetConfigs()
	if err != nil {
		return nil, err
	}
	client, err := routev1client.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetRbacV1Client() (*rbacv1client.RbacV1Client, error) {
	restConfig, err := GetConfigs()
	if err != nil {
		return nil, err
	}
	client, err := rbacv1client.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetConfigs() (*rest.Config, error) {
	// Instantiate loader for kubeconfig file.
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{},
	)

	// TODO check: should take namespace from config or dynamic?
	_, _, err := kubeconfig.Namespace()
	if err != nil {
		return nil, err
	}
	restConfig, err := kubeconfig.ClientConfig()
	if err != nil {
		return nil, err
	}
	return restConfig, nil
}

func GetKubeClientSet() (*kubernetes.Clientset, error) {
	restConfig, err := GetConfigs()
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(restConfig)
	return clientset, nil
}

// Other

func onlyLogOnError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func ParseJsonString(jsonContent string, v interface{}) error {
	return json.Unmarshal([]byte(jsonContent), v)
}

func ObjectToJsonString(v interface{}) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
}
