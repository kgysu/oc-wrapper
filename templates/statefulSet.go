package templates

import (
	v14 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetTemplateStatefulSet(name string) v14.StatefulSet {
	replicas := int32(0)
	return v14.StatefulSet{
		TypeMeta: metav1.TypeMeta{
			Kind: "StatefulSet",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        name,
			Labels:      map[string]string{"app": name},
			Annotations: map[string]string{"app": name},
		},
		Spec: v14.StatefulSetSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"app": name},
			},
			Template:            GetPodTemplateSpec(name),
			ServiceName:         name,
			PodManagementPolicy: v14.OrderedReadyPodManagement,
			UpdateStrategy: v14.StatefulSetUpdateStrategy{
				Type: v14.RollingUpdateStatefulSetStrategyType,
			},
		},
		Status: v14.StatefulSetStatus{},
	}
}
