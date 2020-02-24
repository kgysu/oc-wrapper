package wrapper

import (
	v12 "github.com/openshift/api/route/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type RouteList []v12.Route

func ListRoutes(ns string, options v1.ListOptions) (RouteList, error) {
	routesApi, err := GetRouteApi(ns)
	if err != nil {
		return nil, err
	}
	routes, err := routesApi.List(options)
	if err != nil {
		return nil, err
	}
	return routes.Items, nil
}

func GetRouteByName(ns string, name string, options v1.GetOptions) (*v12.Route, error) {
	routesApi, err := GetRouteApi(ns)
	if err != nil {
		return nil, err
	}
	return routesApi.Get(name, options)
}

func UpdateRoute(ns string, route *v12.Route) (*v12.Route, error) {
	routesApi, err := GetRouteApi(ns)
	if err != nil {
		return nil, err
	}
	return routesApi.Update(route)
}

func CreateRoute(ns string, route *v12.Route) (*v12.Route, error) {
	routesApi, err := GetRouteApi(ns)
	if err != nil {
		return nil, err
	}
	return routesApi.Create(route)
}

func DeleteRoute(ns string, name string, options v1.DeleteOptions) error {
	routesApi, err := GetRouteApi(ns)
	if err != nil {
		return err
	}
	return routesApi.Delete(name, &options)
}

func GetRouteJson(ns string, name string, options v1.GetOptions) (string, error) {
	route, err := GetRouteByName(ns, name, options)
	if err != nil {
		return "", err
	}
	routeData, err := ObjectToJsonString(route)
	if err != nil {
		return "", err
	}
	return string(routeData), nil
}
