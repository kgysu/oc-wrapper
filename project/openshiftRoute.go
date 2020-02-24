package project

import (
	"github.com/kgysu/oc-wrapper/wrapper"
	v1 "github.com/openshift/api/route/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type OpenshiftRoute struct {
	name  string
	route v1.Route
}

func fromRoute(route v1.Route) OpenshiftRoute {
	return OpenshiftRoute{
		name:  route.Name,
		route: route,
	}
}

func (oRoute OpenshiftRoute) setRoute(route v1.Route) {
	oRoute.name = route.Name
	oRoute.route = route
}

func (oRoute OpenshiftRoute) GetName() string {
	return oRoute.name
}

func (oRoute OpenshiftRoute) GetKind() string {
	return RouteKey
}

func (oRoute OpenshiftRoute) GetStatus() string {
	return oRoute.route.Spec.To.Name
}

func (oRoute OpenshiftRoute) GetRoute() v1.Route {
	return oRoute.route
}

func (oRoute OpenshiftRoute) Create(namespace string) error {
	_, err := wrapper.CreateRoute(namespace, &oRoute.route)
	if err != nil {
		return err
	}
	//oRoute.setRoute(createdRoute)
	return nil
}

func (oRoute OpenshiftRoute) Update(namespace string) error {
	_, err := wrapper.UpdateRoute(namespace, &oRoute.route)
	if err != nil {
		return err
	}
	//oRoute.setRoute(updatedRoute)
	return nil
}

func (oRoute OpenshiftRoute) Delete(namespace string, options v12.DeleteOptions) error {
	return wrapper.DeleteRoute(namespace, oRoute.name, options)
}
