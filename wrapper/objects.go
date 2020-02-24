package wrapper

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type MyProjectItems struct {
	deployments  DeploymentConfigList
	pods         PodList
	routes       RouteList
	services     ServiceList
	configMaps   ConfigMapList
	statefulSets StatefulSetList
}

func ListProjectItems(ns string, options v1.ListOptions) (*MyProjectItems, error) {
	dcs, err := ListDeploymentConfigs(ns, options)
	if err != nil {
		return nil, err
	}

	pods, err := ListPods(ns, options)
	if err != nil {
		return nil, err
	}

	routes, err := ListRoutes(ns, options)
	if err != nil {
		return nil, err
	}

	services, err := ListServices(ns, options)
	if err != nil {
		return nil, err
	}

	configMaps, err := ListConfigMaps(ns, options)
	if err != nil {
		return nil, err
	}

	statefulSets, err := ListStatefulSets(ns, options)
	if err != nil {
		return nil, err
	}

	return &MyProjectItems{
		deployments:  dcs,
		pods:         pods,
		routes:       routes,
		services:     services,
		configMaps:   configMaps,
		statefulSets: statefulSets,
	}, nil
}

func UpdateProjectItems(ns string, items MyProjectItems) []error {
	var errs []error

	if items.deployments != nil {
		if len(items.deployments) > 0 {
			for _, dc := range items.deployments {
				_, err := UpdateDeploymentConfig(ns, &dc)
				if err != nil {
					errs = append(errs, err)
				}
			}
		}
	}
	if items.services != nil {
		if len(items.services) > 0 {
			for _, svc := range items.services {
				_, err := UpdateService(ns, &svc)
				if err != nil {
					errs = append(errs, err)
				}
			}
		}
	}
	if items.routes != nil {
		if len(items.routes) > 0 {
			for _, route := range items.routes {
				_, err := UpdateRoute(ns, &route)
				if err != nil {
					errs = append(errs, err)
				}
			}
		}
	}
	if items.statefulSets != nil {
		if len(items.statefulSets) > 0 {
			for _, ss := range items.statefulSets {
				_, err := UpdateStatefulSet(ns, &ss)
				if err != nil {
					errs = append(errs, err)
				}
			}
		}
	}
	if items.configMaps != nil {
		if len(items.configMaps) > 0 {
			for _, cm := range items.configMaps {
				_, err := UpdateConfigMap(ns, &cm)
				if err != nil {
					errs = append(errs, err)
				}
			}
		}
	}

	return errs
}
