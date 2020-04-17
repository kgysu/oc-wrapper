package templates

import (
	v1 "github.com/openshift/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetTemplateDeploymentConfig(name string) v1.DeploymentConfig {
	podTemplateSpec := GetPodTemplateSpec(name)
	return v1.DeploymentConfig{
		TypeMeta: metav1.TypeMeta{
			Kind:       "DeploymentConfig",
			APIVersion: "apps.openshift.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        name,
			Labels:      map[string]string{"app": name},
			Annotations: map[string]string{"app": name},
		},
		Spec: v1.DeploymentConfigSpec{
			Strategy: v1.DeploymentStrategy{},
			Triggers: v1.DeploymentTriggerPolicies{
				{
					Type: v1.DeploymentTriggerOnConfigChange,
				},
			},
			Replicas: 0,
			Template: &podTemplateSpec,
		},
	}
}
