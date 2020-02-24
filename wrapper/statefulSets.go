package wrapper

import (
	"k8s.io/api/apps/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type StatefulSetList []v1beta1.StatefulSet

func ListStatefulSets(ns string, options v1.ListOptions) (StatefulSetList, error) {
	statefulSetsApi, err := GetStatefulSetApi(ns)
	if err != nil {
		return nil, err
	}
	statefulSets, err := statefulSetsApi.List(options)
	if err != nil {
		return nil, err
	}
	return statefulSets.Items, nil
}

func GetStatefulSetByName(ns string, name string, options v1.GetOptions) (*v1beta1.StatefulSet, error) {
	statefulSetsApi, err := GetStatefulSetApi(ns)
	if err != nil {
		return nil, err
	}
	return statefulSetsApi.Get(name, options)
}

func UpdateStatefulSet(ns string, statefulSet *v1beta1.StatefulSet) (*v1beta1.StatefulSet, error) {
	statefulSetsApi, err := GetStatefulSetApi(ns)
	if err != nil {
		return nil, err
	}
	return statefulSetsApi.Update(statefulSet)
}

func CreateStatefulSet(ns string, statefulSet *v1beta1.StatefulSet) (*v1beta1.StatefulSet, error) {
	statefulSetsApi, err := GetStatefulSetApi(ns)
	if err != nil {
		return nil, err
	}
	return statefulSetsApi.Create(statefulSet)
}

func DeleteStatefulSet(ns string, name string, options v1.DeleteOptions) error {
	statefulSetsApi, err := GetStatefulSetApi(ns)
	if err != nil {
		return err
	}
	return statefulSetsApi.Delete(name, &options)
}

func GetStatefulSetJson(ns string, name string, options v1.GetOptions) (string, error) {
	statefulSet, err := GetStatefulSetByName(ns, name, options)
	if err != nil {
		return "", err
	}
	statefulSetData, err := ObjectToJsonString(statefulSet)
	if err != nil {
		return "", err
	}
	return string(statefulSetData), nil
}
