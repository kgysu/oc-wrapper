package templates

import (
	v12 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func GetPodTemplateSpec(name string) v12.PodTemplateSpec {
	return v12.PodTemplateSpec{
		ObjectMeta: metav1.ObjectMeta{
			Name:        name,
			Labels:      map[string]string{"app": name},
			Annotations: map[string]string{"app": name},
		},
		Spec: v12.PodSpec{
			RestartPolicy: v12.RestartPolicyAlways,
			DNSPolicy:     v12.DNSClusterFirst,
			Containers: []v12.Container{
				{
					Name:  name,
					Image: "sample:1.0",
					Ports: []v12.ContainerPort{
						{
							Name:          "basic",
							ContainerPort: 8080,
						},
					},
					Env: []v12.EnvVar{
						{
							Name:  "SAMPLE",
							Value: "VALUE",
						},
					},
					Resources: v12.ResourceRequirements{
						Limits: v12.ResourceList{
							"cpu":    resource.MustParse("1"),
							"memory": resource.MustParse("1Gi"),
						},
						Requests: v12.ResourceList{
							"cpu":    resource.MustParse("100m"),
							"memory": resource.MustParse("100Mi"),
						},
					},
					LivenessProbe: &v12.Probe{
						Handler: v12.Handler{
							HTTPGet: &v12.HTTPGetAction{
								Path:   "/health",
								Port:   intstr.FromInt(8080),
								Scheme: v12.URISchemeHTTP,
							},
						},
						InitialDelaySeconds: 40,
						TimeoutSeconds:      5,
						PeriodSeconds:       10,
						SuccessThreshold:    1,
						FailureThreshold:    3,
					},
					ReadinessProbe: &v12.Probe{
						Handler: v12.Handler{
							HTTPGet: &v12.HTTPGetAction{
								Path:   "/info",
								Port:   intstr.FromInt(8080),
								Scheme: v12.URISchemeHTTP,
							},
						},
						InitialDelaySeconds: 40,
						TimeoutSeconds:      5,
						PeriodSeconds:       10,
						SuccessThreshold:    1,
						FailureThreshold:    3,
					},
					ImagePullPolicy: v12.PullIfNotPresent,
				},
			},
		},
	}
}
