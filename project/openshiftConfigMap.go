package project

import (
	"github.com/kgysu/oc-wrapper/wrapper"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type OpenshiftConfigMap struct {
	name      string
	configMap v1.ConfigMap
}

func fromConfigMap(configMap v1.ConfigMap) OpenshiftConfigMap {
	return OpenshiftConfigMap{
		name:      configMap.Name,
		configMap: configMap,
	}
}

func (ocm OpenshiftConfigMap) setConfigMap(configMap v1.ConfigMap) {
	ocm.name = configMap.Name
	ocm.configMap = configMap
}

func (ocm OpenshiftConfigMap) GetName() string {
	return ocm.name
}

func (ocm OpenshiftConfigMap) GetKind() string {
	return ConfigMapKey
}

func (ocm OpenshiftConfigMap) GetStatus() string {
	return ocm.configMap.CreationTimestamp.String()
}

func (ocm OpenshiftConfigMap) GetConfigMap() v1.ConfigMap {
	return ocm.configMap
}

func (ocm OpenshiftConfigMap) Create(namespace string) error {
	_, err := wrapper.CreateConfigMap(namespace, &ocm.configMap)
	if err != nil {
		return err
	}
	//ocm.setConfigMap(cm)
	return nil
}

func (ocm OpenshiftConfigMap) Update(namespace string) error {
	_, err := wrapper.UpdateConfigMap(namespace, &ocm.configMap)
	if err != nil {
		return err
	}
	//ocm.setConfigMap(updatedCm)
	return nil
}

func (ocm OpenshiftConfigMap) Delete(namespace string, options v12.DeleteOptions) error {
	return wrapper.DeletePod(namespace, ocm.name, &options)
}
