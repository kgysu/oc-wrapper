package wrapper

import (
	v1 "github.com/openshift/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
)

type DeploymentConfigList []v1.DeploymentConfig

func ListDeploymentConfigs(ns string, options metav1.ListOptions) (DeploymentConfigList, error) {
	deploymentsApi, err := GetDeploymentConfigApi(ns)
	if err != nil {
		return nil, err
	}
	deployments, err := deploymentsApi.List(options)
	if err != nil {
		return nil, err
	}
	return deployments.Items, nil
}

func GetDeploymentConfigByName(ns string, name string, options metav1.GetOptions) (*v1.DeploymentConfig, error) {
	deploymentsApi, err := GetDeploymentConfigApi(ns)
	if err != nil {
		return nil, err
	}
	return deploymentsApi.Get(name, options)
}

func UpdateDeploymentConfig(ns string, dc *v1.DeploymentConfig) (*v1.DeploymentConfig, error) {
	deploymentsApi, err := GetDeploymentConfigApi(ns)
	if err != nil {
		return nil, err
	}
	return deploymentsApi.Update(dc)
}

func CreateDeploymentConfig(ns string, dc *v1.DeploymentConfig) (*v1.DeploymentConfig, error) {
	deploymentsApi, err := GetDeploymentConfigApi(ns)
	if err != nil {
		return nil, err
	}
	return deploymentsApi.Create(dc)
}

func DeleteDeploymentConfig(ns string, name string, options metav1.DeleteOptions) error {
	deploymentsApi, err := GetDeploymentConfigApi(ns)
	if err != nil {
		return err
	}
	return deploymentsApi.Delete(name, &options)
}

func WatchDeploymentConfig(ns string, options metav1.ListOptions) (watch.Interface, error) {
	deploymentsApi, err := GetDeploymentConfigApi(ns)
	if err != nil {
		return nil, err
	}
	return deploymentsApi.Watch(options)
}

func GetDeploymentConfigJson(ns string, name string, options metav1.GetOptions) (string, error) {
	deployment, err := GetDeploymentConfigByName(ns, name, options)
	if err != nil {
		return "", err
	}
	deploymentData, err := ObjectToJsonString(deployment)
	if err != nil {
		return "", err
	}
	return string(deploymentData), nil
}
