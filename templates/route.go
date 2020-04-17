package templates

import (
	v13 "github.com/openshift/api/route/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

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
