package wrapper

import (
	v12 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ServiceList []v12.Service

func ListServices(ns string, options v1.ListOptions) (ServiceList, error) {
	servicesApi, err := GetServiceApi(ns)
	if err != nil {
		return nil, err
	}
	services, err := servicesApi.List(options)
	if err != nil {
		return nil, err
	}
	return services.Items, nil
}

func GetServiceByName(ns string, name string, options v1.GetOptions) (*v12.Service, error) {
	servicesApi, err := GetServiceApi(ns)
	if err != nil {
		return nil, err
	}
	return servicesApi.Get(name, options)
}

func UpdateService(ns string, service *v12.Service) (*v12.Service, error) {
	servicesApi, err := GetServiceApi(ns)
	if err != nil {
		return nil, err
	}
	return servicesApi.Update(service)
}

func CreateService(ns string, service *v12.Service) (*v12.Service, error) {
	servicesApi, err := GetServiceApi(ns)
	if err != nil {
		return nil, err
	}
	return servicesApi.Create(service)
}

func DeleteService(ns string, name string, options v1.DeleteOptions) error {
	servicesApi, err := GetServiceApi(ns)
	if err != nil {
		return err
	}
	return servicesApi.Delete(name, &options)
}

func GetServiceJson(ns string, name string, options v1.GetOptions) (string, error) {
	service, err := GetServiceByName(ns, name, options)
	if err != nil {
		return "", err
	}
	serviceData, err := ObjectToJsonString(service)
	if err != nil {
		return "", err
	}
	return string(serviceData), nil
}
