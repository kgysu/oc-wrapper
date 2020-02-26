package v2

import (
	"errors"
	"fmt"
	v1 "github.com/openshift/api/apps/v1"
	v14 "github.com/openshift/client-go/apps/clientset/versioned/typed/apps/v1"
	v12 "k8s.io/api/apps/v1"
	v13 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v15 "k8s.io/client-go/kubernetes/typed/apps/v1"
)

// List
func ScaleDeploymentConfigsAndStatefulSets(namespace string, items []OpenshiftItem, replicas int) ([]OpenshiftItem, error) {
	kubeClient, err := GetKubeAppsV1Client()
	if err != nil {
		return nil, err
	}
	appsClient, err := GetAppsV1Client()
	if err != nil {
		return nil, err
	}
	var resultItems []OpenshiftItem

	for _, item := range items {
		fmt.Printf("Scale %s: %s", item.kind, item.name)
		if item.kind == DeploymentConfigKey {
			err := scaleDeploymentConfig(namespace, replicas, appsClient, item, &resultItems)
			onlyLogOnError(err)
		}
		if item.kind == StatefulSetKey {
			err := scaleStatefulSet(namespace, replicas, kubeClient, item, &resultItems)
			onlyLogOnError(err)
		}
	}
	return resultItems, nil
}

func scaleDeploymentConfig(namespace string, replicas int, appsClient *v14.AppsV1Client, item OpenshiftItem, resultItems *[]OpenshiftItem) error {
	var realItem v1.DeploymentConfig
	err := item.ParseTo(&realItem)
	if err != nil {
		return err
	}
	fromRemote, err := appsClient.DeploymentConfigs(namespace).Get(realItem.Name, v13.GetOptions{})
	if err != nil {
		return err
	}
	if fromRemote != nil {
		fromRemote.Spec.Replicas = int32(replicas)
		updated, err := appsClient.DeploymentConfigs(namespace).Update(fromRemote)
		if err != nil {
			return err
		}
		*resultItems = append(*resultItems, NewOpenshiftItem(updated.Name, DeploymentConfigKey, updated.String()))
		fmt.Printf("Scaled %s(%s) to %d", item.name, item.kind, replicas)
		return nil
	}
	return errors.New("could not update scale, no remote found")
}

func scaleStatefulSet(namespace string, replicas int, appsClient v15.AppsV1Interface, item OpenshiftItem, resultItems *[]OpenshiftItem) error {
	var realItem v12.StatefulSet
	err := item.ParseTo(&realItem)
	if err != nil {
		return err
	}
	fromRemote, err := appsClient.StatefulSets(namespace).Get(realItem.Name, v13.GetOptions{})
	if err != nil {
		return err
	}
	if fromRemote != nil {
		replicas32 := int32(replicas)
		fromRemote.Spec.Replicas = &replicas32
		updated, err := appsClient.StatefulSets(namespace).Update(fromRemote)
		if err != nil {
			return err
		}
		*resultItems = append(*resultItems, NewOpenshiftItem(updated.Name, StatefulSetKey, updated.String()))
		fmt.Printf("Scaled %s(%s) to %d", item.name, item.kind, replicas)
		return nil
	}
	return errors.New("could not update scale, no remote found")
}
