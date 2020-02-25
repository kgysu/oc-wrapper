package v2

import (
	"encoding/json"
	appsv1client "github.com/openshift/client-go/apps/clientset/versioned/typed/apps/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/typed/apps/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// List

func ListAllFromRemote(namespace string) ([]OpenshiftItem, error) {
	kubeClient, err := GetKubeAppsV1Client()
	if err != nil {
		return nil, err
	}
	appsClient, err := GetAppsV1Client()
	if err != nil {
		return nil, err
	}
	var results []OpenshiftItem
	dcs, err := kubeClient.Deployments(namespace).List(v12.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, dc := range dcs.Items {
		results = append(results, New(dc.Name, dc.Kind, dc.String()))
	}
	statefulSets, err := kubeClient.StatefulSets(namespace).List(v12.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, ss := range statefulSets.Items {
		results = append(results, New(ss.Name, ss.Kind, ss.String()))
	}
	replicaSets, err := kubeClient.ReplicaSets(namespace).List(v12.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, rs := range replicaSets.Items {
		results = append(results, New(rs.Name, rs.Kind, rs.String()))
	}
	return results, nil
}

// Util

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

// json

func ParseJsonString(jsonContent string, v interface{}) error {
	return json.Unmarshal([]byte(jsonContent), v)
}

func ObjectToJsonString(v interface{}) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
}
