package client

import (
	"fmt"
	appsv1client "github.com/openshift/client-go/apps/clientset/versioned/typed/apps/v1"
	authorizationv1client "github.com/openshift/client-go/authorization/clientset/versioned/typed/authorization/v1"
	routev1client "github.com/openshift/client-go/route/clientset/versioned/typed/route/v1"
	kubeappsv1client "k8s.io/client-go/kubernetes/typed/apps/v1"
	appsv1beta1client "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
	appsv1beta2client "k8s.io/client-go/kubernetes/typed/apps/v1beta2"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	rbacv1client "k8s.io/client-go/kubernetes/typed/rbac/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

func GetAuthorizationV1Client(restConf *rest.Config) (*authorizationv1client.AuthorizationV1Client, error) {
	client, err := authorizationv1client.NewForConfig(restConf)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetKubeAppsV1Client(restConf *rest.Config) (*kubeappsv1client.AppsV1Client, error) {
	client, err := kubeappsv1client.NewForConfig(restConf)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetKubeAppsV1Beta1Client(restConf *rest.Config) (*appsv1beta1client.AppsV1beta1Client, error) {
	client, err := appsv1beta1client.NewForConfig(restConf)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetKubeAppsV1Beta2Client(restConf *rest.Config) (*appsv1beta2client.AppsV1beta2Client, error) {
	client, err := appsv1beta2client.NewForConfig(restConf)
	if err != nil {
		return nil, err
	}
	return client, nil
}

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

func GetNamespace(fromEnv bool, envVarName string) string {
	if fromEnv {
		namespace := os.Getenv(envVarName)
		if namespace == "" {
			panic(envVarName + ", namespace is not defined")
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
