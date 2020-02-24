package project

import (
	"fmt"
	"github.com/kgysu/oc-wrapper/wrapper"
	v1 "github.com/openshift/api/apps/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type OpenshiftDeploymentConfig struct {
	name             string
	deploymentConfig v1.DeploymentConfig
}

func fromDeploymentConfig(deploymentConfig v1.DeploymentConfig) OpenshiftDeploymentConfig {
	return OpenshiftDeploymentConfig{
		name:             deploymentConfig.Name,
		deploymentConfig: deploymentConfig,
	}
}

func (oDeploymentConfig OpenshiftDeploymentConfig) setDeploymentConfig(deploymentConfig v1.DeploymentConfig) {
	oDeploymentConfig.name = deploymentConfig.Name
	oDeploymentConfig.deploymentConfig = deploymentConfig
}

func (oDeploymentConfig OpenshiftDeploymentConfig) setDeploymentConfigRef(deploymentConfig *v1.DeploymentConfig) {
	oDeploymentConfig.name = deploymentConfig.Name
	oDeploymentConfig.deploymentConfig.Spec = deploymentConfig.Spec
	oDeploymentConfig.deploymentConfig.Name = deploymentConfig.Name
	oDeploymentConfig.deploymentConfig.Status = deploymentConfig.Status
	oDeploymentConfig.deploymentConfig.CreationTimestamp = deploymentConfig.CreationTimestamp
	oDeploymentConfig.deploymentConfig.Namespace = deploymentConfig.Namespace
	oDeploymentConfig.deploymentConfig.Annotations = deploymentConfig.Annotations
	oDeploymentConfig.deploymentConfig.Generation = deploymentConfig.Generation
	oDeploymentConfig.deploymentConfig.Labels = deploymentConfig.Labels
	oDeploymentConfig.deploymentConfig.ResourceVersion = deploymentConfig.ResourceVersion
	oDeploymentConfig.deploymentConfig.OwnerReferences = deploymentConfig.OwnerReferences
	oDeploymentConfig.deploymentConfig.UID = deploymentConfig.UID
}

func (oDeploymentConfig OpenshiftDeploymentConfig) GetName() string {
	return oDeploymentConfig.name
}

func (oDeploymentConfig OpenshiftDeploymentConfig) GetKind() string {
	return DeploymentConfigKey
}

func (oDeploymentConfig OpenshiftDeploymentConfig) GetStatus() string {
	return fmt.Sprintf("%d (%d/%d)",
		oDeploymentConfig.deploymentConfig.Status.Replicas,
		oDeploymentConfig.deploymentConfig.Status.ReadyReplicas,
		oDeploymentConfig.deploymentConfig.Status.AvailableReplicas)
}

func (oDeploymentConfig OpenshiftDeploymentConfig) GetDeploymentConfig() v1.DeploymentConfig {
	return oDeploymentConfig.deploymentConfig
}

func (oDeploymentConfig OpenshiftDeploymentConfig) Create(namespace string) error {
	createdDeploymentConfig, err := wrapper.CreateDeploymentConfig(namespace, &oDeploymentConfig.deploymentConfig)
	if err != nil {
		return err
	}
	oDeploymentConfig.setDeploymentConfigRef(createdDeploymentConfig)
	return nil
}

func (oDeploymentConfig OpenshiftDeploymentConfig) Update(namespace string) error {
	updatedDeploymentConfig, err := wrapper.UpdateDeploymentConfig(namespace, &oDeploymentConfig.deploymentConfig)
	if err != nil {
		return err
	}
	oDeploymentConfig.setDeploymentConfigRef(updatedDeploymentConfig)
	return nil
}

func (oDeploymentConfig OpenshiftDeploymentConfig) Delete(namespace string, options v12.DeleteOptions) error {
	return wrapper.DeleteDeploymentConfig(namespace, oDeploymentConfig.name, options)
}
