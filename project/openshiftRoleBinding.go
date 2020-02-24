package project

import (
	"github.com/kgysu/oc-wrapper/wrapper"
	v1 "k8s.io/api/rbac/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type OpenshiftRoleBinding struct {
	name        string
	roleBinding v1.RoleBinding
}

func fromRoleBinding(roleBinding v1.RoleBinding) OpenshiftRoleBinding {
	return OpenshiftRoleBinding{
		name:        roleBinding.Name,
		roleBinding: roleBinding,
	}
}

func (oRoleBinding OpenshiftRoleBinding) setRoleBinding(roleBinding v1.RoleBinding) {
	oRoleBinding.name = roleBinding.Name
	oRoleBinding.roleBinding = roleBinding
}

func (oRoleBinding OpenshiftRoleBinding) GetName() string {
	return oRoleBinding.name
}

func (oRoleBinding OpenshiftRoleBinding) GetKind() string {
	return RoleBindingKey
}

func (oRoleBinding OpenshiftRoleBinding) GetStatus() string {
	return oRoleBinding.roleBinding.CreationTimestamp.String()
}

func (oRoleBinding OpenshiftRoleBinding) GetRoleBinding() v1.RoleBinding {
	return oRoleBinding.roleBinding
}

func (oRoleBinding OpenshiftRoleBinding) Create(namespace string) error {
	_, err := wrapper.CreateRoleBinding(namespace, &oRoleBinding.roleBinding)
	if err != nil {
		return err
	}
	//oRoleBinding.setRoleBinding(createdRoleBinding)
	return nil
}

func (oRoleBinding OpenshiftRoleBinding) Update(namespace string) error {
	_, err := wrapper.UpdateRoleBinding(namespace, &oRoleBinding.roleBinding)
	if err != nil {
		return err
	}
	//oRoleBinding.setRoleBinding(updatedRoleBinding)
	return nil
}

func (oRoleBinding OpenshiftRoleBinding) Delete(namespace string, options v12.DeleteOptions) error {
	return wrapper.DeleteRoleBinding(namespace, oRoleBinding.name, options)
}
