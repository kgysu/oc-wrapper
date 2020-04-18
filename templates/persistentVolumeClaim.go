package templates

import (
	v12 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetTemplatePersistentVolumeClaim(name string) v12.PersistentVolumeClaim {
	volumeMode := v12.PersistentVolumeFilesystem
	return v12.PersistentVolumeClaim{
		TypeMeta: v1.TypeMeta{
			Kind:       "PersistentVolumeClaim",
			APIVersion: "v1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:        name,
			Labels:      map[string]string{"app": name},
			Annotations: map[string]string{"app": name},
		},
		Spec: v12.PersistentVolumeClaimSpec{
			AccessModes: []v12.PersistentVolumeAccessMode{v12.ReadWriteOnce},
			//Selector: &v1.LabelSelector{
			//	MatchLabels: map[string]string{"app": name},
			//},
			Resources: v12.ResourceRequirements{
				Requests: v12.ResourceList{
					v12.ResourceStorage: resource.MustParse("2G"),
				},
			},
			VolumeName: name,
			VolumeMode: &volumeMode,
		},
		Status: v12.PersistentVolumeClaimStatus{},
	}
}
