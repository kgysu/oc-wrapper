package wrapper

import (
	v12 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ConfigMapList []v12.ConfigMap

func ListConfigMaps(ns string, options v1.ListOptions) (ConfigMapList, error) {
	configMapsApi, err := GetConfigMapApi(ns)
	if err != nil {
		return nil, err
	}
	configMaps, err := configMapsApi.List(options)
	if err != nil {
		return nil, err
	}
	return configMaps.Items, nil
}

func GetConfigMapByName(ns string, name string, options v1.GetOptions) (*v12.ConfigMap, error) {
	configMapsApi, err := GetConfigMapApi(ns)
	if err != nil {
		return nil, err
	}
	return configMapsApi.Get(name, options)
}

func UpdateConfigMap(ns string, configMap *v12.ConfigMap) (*v12.ConfigMap, error) {
	configMapsApi, err := GetConfigMapApi(ns)
	if err != nil {
		return nil, err
	}
	return configMapsApi.Update(configMap)
}

func CreateConfigMap(ns string, configMap *v12.ConfigMap) (*v12.ConfigMap, error) {
	configMapsApi, err := GetConfigMapApi(ns)
	if err != nil {
		return nil, err
	}
	return configMapsApi.Create(configMap)
}

func DeleteConfigMap(ns string, name string, options v1.DeleteOptions) error {
	configMapsApi, err := GetConfigMapApi(ns)
	if err != nil {
		return err
	}
	return configMapsApi.Delete(name, &options)
}

func GetConfigMapJson(ns string, name string, options v1.GetOptions) (string, error) {
	configMap, err := GetConfigMapByName(ns, name, options)
	if err != nil {
		return "", err
	}
	configMapData, err := ObjectToJsonString(configMap)
	if err != nil {
		return "", err
	}
	return string(configMapData), nil
}
