package templates

import (
	v15 "github.com/openshift/api/authorization/v1"
	v12 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetTemplateRoleBinding(name string) v15.RoleBinding {
	return v15.RoleBinding{
		TypeMeta: metav1.TypeMeta{
			Kind:       "RoleBinding",
			APIVersion: "authorization.openshift.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        name,
			Labels:      map[string]string{"app": name},
			Annotations: map[string]string{"app": name},
		},
		UserNames:  []string{name},
		GroupNames: nil,
		Subjects: []v12.ObjectReference{
			{
				Kind: "ServiceAccount",
				Name: "sa",
			},
		},
		RoleRef: v12.ObjectReference{
			Kind: "Role",
			Name: "rolename",
		},
	}
}
