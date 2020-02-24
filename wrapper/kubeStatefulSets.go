package wrapper

import (
	v1 "k8s.io/api/apps/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type KubeStatefulSetList []v1.StatefulSet

func ListKubeStatefulSets(ns string, options v12.ListOptions) ([]v1.StatefulSet, error) {
	statefulSetsApi, err := GetKubeStatefulSetApi(ns)
	if err != nil {
		return nil, err
	}
	statefulSets, err := statefulSetsApi.List(options)
	if err != nil {
		return nil, err
	}
	return statefulSets.Items, nil
}

func GetKubeStatefulSetByName(ns string, name string, options v12.GetOptions) (*v1.StatefulSet, error) {
	statefulSetsApi, err := GetKubeStatefulSetApi(ns)
	if err != nil {
		return nil, err
	}
	return statefulSetsApi.Get(name, options)
}

func UpdateKubeStatefulSet(ns string, statefulSet *v1.StatefulSet) (*v1.StatefulSet, error) {
	statefulSetsApi, err := GetKubeStatefulSetApi(ns)
	if err != nil {
		return nil, err
	}
	return statefulSetsApi.Update(statefulSet)
}

func CreateKubeStatefulSet(ns string, statefulSet *v1.StatefulSet) (*v1.StatefulSet, error) {
	statefulSetsApi, err := GetKubeStatefulSetApi(ns)
	if err != nil {
		return nil, err
	}
	return statefulSetsApi.Create(statefulSet)
}

func DeleteKubeStatefulSet(ns string, name string, options v12.DeleteOptions) error {
	statefulSetsApi, err := GetKubeStatefulSetApi(ns)
	if err != nil {
		return err
	}
	return statefulSetsApi.Delete(name, &options)
}

func GetKubeStatefulSetJson(ns string, name string, options v12.GetOptions) (string, error) {
	statefulSet, err := GetKubeStatefulSetByName(ns, name, options)
	if err != nil {
		return "", err
	}
	statefulSetData, err := ObjectToJsonString(statefulSet)
	if err != nil {
		return "", err
	}
	return string(statefulSetData), nil
}
