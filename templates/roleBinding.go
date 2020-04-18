package templates

import (
	v1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetTemplateRoleBinding(name string) v1.RoleBinding {
	return v1.RoleBinding{
		TypeMeta: metav1.TypeMeta{
			Kind:       "RoleBinding",
			APIVersion: "rbac.authorization.k8s.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        name,
			Labels:      map[string]string{"app": name},
			Annotations: map[string]string{"app": name},
		},
		//UserNames:  []string{name},
		//GroupNames: []string{name},
		Subjects: []v1.Subject{
			{
				Kind: "ServiceAccount",
				Name: name,
			},
		},
		RoleRef: v1.RoleRef{
			Kind:     "Role",
			Name:     name,
			APIGroup: "rbac.authorization.k8s.io/v1",
		},
	}
}
