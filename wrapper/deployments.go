package wrapper

import (
	"k8s.io/api/apps/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
)

type DeploymentList []v1beta1.Deployment

func ListDeployments(ns string, options metav1.ListOptions) (DeploymentList, error) {
	deploymentsApi, err := GetDeploymentsApi(ns)
	if err != nil {
		return nil, err
	}
	deployments, err := deploymentsApi.List(options)
	if err != nil {
		return nil, err
	}
	return deployments.Items, nil
}

func GetDeploymentByName(ns string, name string, options metav1.GetOptions) (*v1beta1.Deployment, error) {
	deploymentsApi, err := GetDeploymentsApi(ns)
	if err != nil {
		return nil, err
	}
	return deploymentsApi.Get(name, options)
}

func UpdateDeployment(ns string, dc *v1beta1.Deployment) (*v1beta1.Deployment, error) {
	deploymentsApi, err := GetDeploymentsApi(ns)
	if err != nil {
		return nil, err
	}
	return deploymentsApi.Update(dc)
}

func CreateDeployment(ns string, dc *v1beta1.Deployment) (*v1beta1.Deployment, error) {
	deploymentsApi, err := GetDeploymentsApi(ns)
	if err != nil {
		return nil, err
	}
	return deploymentsApi.Create(dc)
}

func DeleteDeployment(ns string, name string, options metav1.DeleteOptions) error {
	deploymentsApi, err := GetDeploymentsApi(ns)
	if err != nil {
		return err
	}
	return deploymentsApi.Delete(name, &options)
}

func WatchDeployment(ns string, options metav1.ListOptions) (watch.Interface, error) {
	deploymentsApi, err := GetDeploymentsApi(ns)
	if err != nil {
		return nil, err
	}
	return deploymentsApi.Watch(options)
}

func GetDeploymentJson(ns string, name string, options metav1.GetOptions) (string, error) {
	deployment, err := GetDeploymentByName(ns, name, options)
	if err != nil {
		return "", err
	}
	deploymentData, err := ObjectToJsonString(deployment)
	if err != nil {
		return "", err
	}
	return string(deploymentData), nil
}
