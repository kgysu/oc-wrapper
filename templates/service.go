package templates

import (
	v12 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

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
