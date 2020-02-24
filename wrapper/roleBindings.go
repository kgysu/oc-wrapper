package wrapper

import (
	v12 "k8s.io/api/rbac/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type RoleBindingList []v12.RoleBinding

func ListRoleBindings(ns string, options v1.ListOptions) (RoleBindingList, error) {
	roleBindingsApi, err := GetRoleBindingApi(ns)
	if err != nil {
		return nil, err
	}
	roleBindings, err := roleBindingsApi.List(options)
	if err != nil {
		return nil, err
	}
	return roleBindings.Items, nil
}

func GetRoleBindingByName(ns string, name string, options v1.GetOptions) (*v12.RoleBinding, error) {
	roleBindingsApi, err := GetRoleBindingApi(ns)
	if err != nil {
		return nil, err
	}
	return roleBindingsApi.Get(name, options)
}

func UpdateRoleBinding(ns string, roleBinding *v12.RoleBinding) (*v12.RoleBinding, error) {
	roleBindingsApi, err := GetRoleBindingApi(ns)
	if err != nil {
		return nil, err
	}
	return roleBindingsApi.Update(roleBinding)
}

func CreateRoleBinding(ns string, roleBinding *v12.RoleBinding) (*v12.RoleBinding, error) {
	roleBindingsApi, err := GetRoleBindingApi(ns)
	if err != nil {
		return nil, err
	}
	return roleBindingsApi.Create(roleBinding)
}

func DeleteRoleBinding(ns string, name string, options v1.DeleteOptions) error {
	roleBindingsApi, err := GetRoleBindingApi(ns)
	if err != nil {
		return err
	}
	return roleBindingsApi.Delete(name, &options)
}

func GetRoleBindingJson(ns string, name string, options v1.GetOptions) (string, error) {
	roleBinding, err := GetRoleBindingByName(ns, name, options)
	if err != nil {
		return "", err
	}
	roleBindingData, err := ObjectToJsonString(roleBinding)
	if err != nil {
		return "", err
	}
	return string(roleBindingData), nil
}
