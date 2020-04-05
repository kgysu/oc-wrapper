package samples

import (
	"github.com/kgysu/oc-wrapper/items"
	"github.com/kgysu/oc-wrapper/project"
	v1 "github.com/openshift/api/apps/v1"
	v13 "github.com/openshift/api/route/v1"
	v12 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func GetTemplateProject(name string) project.OpenshiftProject {
	return project.OpenshiftProject{
		Name:         name,
		Environments: nil,
		Items: []project.OpenshiftItemInterface{
			items.NewOpDeploymentConfig(GetTemplateDeploymentConfig(name)),
			items.NewOpService(GetTemplateService(name)),
			items.NewOpRoute(GetTemplateRoute(name)),
		},
	}
}

func GetTemplateDeploymentConfig(name string) v1.DeploymentConfig {
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
			Template: &v12.PodTemplateSpec{
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
			},
		},
	}
}

func GetTemplateService(name string) v12.Service {
	return v12.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        name,
			Labels:      map[string]string{"app": name},
			Annotations: map[string]string{"app": name},
		},
		Spec: v12.ServiceSpec{
			Ports: []v12.ServicePort{
				{
					Name:       "basic",
					Protocol:   v12.ProtocolTCP,
					Port:       8080,
					TargetPort: intstr.FromInt(8080),
				},
			},
			Selector:                 map[string]string{"app": name},
			Type:                     v12.ServiceTypeClusterIP,
			SessionAffinity:          v12.ServiceAffinityNone,
			PublishNotReadyAddresses: false,
		},
		Status: v12.ServiceStatus{},
	}
}

func GetTemplateRoute(name string) v13.Route {
	var defaultWeight = int32(100)

	return v13.Route{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Route",
			APIVersion: "route.openshift.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        name,
			Labels:      map[string]string{"app": name},
			Annotations: map[string]string{"app": name},
		},
		Spec: v13.RouteSpec{
			Host: name + "-route",
			Path: "",
			To: v13.RouteTargetReference{
				Kind:   "Service",
				Name:   name,
				Weight: &defaultWeight,
			},
			Port: &v13.RoutePort{TargetPort: intstr.FromInt(8080)},
			TLS: &v13.TLSConfig{
				Termination:                   v13.TLSTerminationEdge,
				InsecureEdgeTerminationPolicy: v13.InsecureEdgeTerminationPolicyRedirect,
			},
			WildcardPolicy: v13.WildcardPolicyNone,
		},
		Status: v13.RouteStatus{},
	}

}