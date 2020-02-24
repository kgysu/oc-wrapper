package wrapper

import (
	v12 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NamespaceList []v12.Namespace

func ListNamespaces(ns string, options v1.ListOptions) (NamespaceList, error) {
	namespacesApi, err := GetNamespaceApi(ns)
	if err != nil {
		return nil, err
	}
	namespaces, err := namespacesApi.List(options)
	if err != nil {
		return nil, err
	}
	return namespaces.Items, nil
}

func GetNamespaceByName(ns string, name string, options v1.GetOptions) (*v12.Namespace, error) {
	namespacesApi, err := GetNamespaceApi(ns)
	if err != nil {
		return nil, err
	}
	return namespacesApi.Get(name, options)
}

func UpdateNamespace(ns string, namespace *v12.Namespace) (*v12.Namespace, error) {
	namespacesApi, err := GetNamespaceApi(ns)
	if err != nil {
		return nil, err
	}
	return namespacesApi.Update(namespace)
}

func CreateNamespace(ns string, namespace *v12.Namespace) (*v12.Namespace, error) {
	namespacesApi, err := GetNamespaceApi(ns)
	if err != nil {
		return nil, err
	}
	return namespacesApi.Create(namespace)
}

func DeleteNamespace(ns string, name string, options v1.DeleteOptions) error {
	namespacesApi, err := GetNamespaceApi(ns)
	if err != nil {
		return err
	}
	return namespacesApi.Delete(name, &options)
}

func GetNamespaceJson(ns string, name string, options v1.GetOptions) (string, error) {
	namespace, err := GetNamespaceByName(ns, name, options)
	if err != nil {
		return "", err
	}
	namespaceData, err := ObjectToJsonString(namespace)
	if err != nil {
		return "", err
	}
	return string(namespaceData), nil
}
