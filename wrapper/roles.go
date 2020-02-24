package wrapper

import (
	v12 "k8s.io/api/rbac/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type RoleList []v12.Role

func ListRoles(ns string, options v1.ListOptions) (RoleList, error) {
	rolesApi, err := GetRoleApi(ns)
	if err != nil {
		return nil, err
	}
	roles, err := rolesApi.List(options)
	if err != nil {
		return nil, err
	}
	return roles.Items, nil
}

func GetRoleByName(ns string, name string, options v1.GetOptions) (*v12.Role, error) {
	rolesApi, err := GetRoleApi(ns)
	if err != nil {
		return nil, err
	}
	return rolesApi.Get(name, options)
}

func UpdateRole(ns string, role *v12.Role) (*v12.Role, error) {
	rolesApi, err := GetRoleApi(ns)
	if err != nil {
		return nil, err
	}
	return rolesApi.Update(role)
}

func CreateRole(ns string, role *v12.Role) (*v12.Role, error) {
	rolesApi, err := GetRoleApi(ns)
	if err != nil {
		return nil, err
	}
	return rolesApi.Create(role)
}

func DeleteRole(ns string, name string, options v1.DeleteOptions) error {
	rolesApi, err := GetRoleApi(ns)
	if err != nil {
		return err
	}
	return rolesApi.Delete(name, &options)
}

func GetRoleJson(ns string, name string, options v1.GetOptions) (string, error) {
	role, err := GetRoleByName(ns, name, options)
	if err != nil {
		return "", err
	}
	roleData, err := ObjectToJsonString(role)
	if err != nil {
		return "", err
	}
	return string(roleData), nil
}
