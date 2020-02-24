package project

import (
	"github.com/kgysu/oc-wrapper/wrapper"
	v1 "k8s.io/api/rbac/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type OpenshiftRole struct {
	name string
	role v1.Role
}

func fromRole(role v1.Role) OpenshiftRole {
	return OpenshiftRole{
		name: role.Name,
		role: role,
	}
}

func (oRole OpenshiftRole) setRole(role v1.Role) {
	oRole.name = role.Name
	oRole.role = role
}

func (oRole OpenshiftRole) GetName() string {
	return oRole.name
}

func (oRole OpenshiftRole) GetKind() string {
	return RoleKey
}

func (oRole OpenshiftRole) GetStatus() string {
	return oRole.role.CreationTimestamp.String()
}

func (oRole OpenshiftRole) GetRole() v1.Role {
	return oRole.role
}

func (oRole OpenshiftRole) Create(namespace string) error {
	_, err := wrapper.CreateRole(namespace, &oRole.role)
	if err != nil {
		return err
	}
	//oRole.setRole(createdRole)
	return nil
}

func (oRole OpenshiftRole) Update(namespace string) error {
	_, err := wrapper.UpdateRole(namespace, &oRole.role)
	if err != nil {
		return err
	}
	//oRole.setRole(updatedRole)
	return nil
}

func (oRole OpenshiftRole) Delete(namespace string, options v12.DeleteOptions) error {
	return wrapper.DeleteRole(namespace, oRole.name, options)
}
