package templates

import (
	v15 "github.com/openshift/api/authorization/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

func GetTemplateRole(name string) v15.Role {
	return v15.Role{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Role",
			APIVersion: "authorization.openshift.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        name,
			Labels:      map[string]string{"app": name},
			Annotations: map[string]string{"app": name},
		},
		Rules: []v15.PolicyRule{
			{
				Verbs:                 []string{"get"},
				AttributeRestrictions: runtime.RawExtension{},
				APIGroups:             []string{""},
				Resources:             []string{"services"},
			},
		},
	}
}
