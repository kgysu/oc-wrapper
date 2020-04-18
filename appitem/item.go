package appitem

import (
	"fmt"
	"github.com/kgysu/oc-wrapper/items"
	v1 "github.com/openshift/api/apps/v1"
	v13 "github.com/openshift/api/route/v1"
	v15 "k8s.io/api/apps/v1"
	v14 "k8s.io/api/core/v1"
	v17 "k8s.io/api/rbac/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"strings"
)

type AppItem interface {
	WriteToFile(file string) error
	LoadFromFile(file string, envs map[string]string) error
	GetFileName() string
	String() string
	Info() string
	Status() string
	InfoStatusHtml() string
	Create(namespace string, restConf *rest.Config) error
	Delete(namespace string, restConf *rest.Config, options *v12.DeleteOptions) error
	Update(namespace string, restConf *rest.Config) error
	UpdateScale(replicas int32, namespace string, restConf *rest.Config) error
	GetScale() int32
	IsScalable() bool
	GetName() string
	GetKind() string
	ToYaml() (string, error)
	FromData(data []byte) error
}

// Hint: add more types here
func NewAppItemFromFile(file string, envs map[string]string) (AppItem, error) {
	if strings.HasSuffix(file, "DeploymentConfig.yaml") {
		item := items.NewOpDeploymentConfig(v1.DeploymentConfig{})
		err := item.LoadFromFile(file, envs)
		return item, err
	}
	if strings.HasSuffix(file, "Service.yaml") {
		item := items.NewOpService(v14.Service{})
		err := item.LoadFromFile(file, envs)
		return item, err
	}
	if strings.HasSuffix(file, "Route.yaml") {
		item := items.NewOpRoute(v13.Route{})
		err := item.LoadFromFile(file, envs)
		return item, err
	}
	if strings.HasSuffix(file, "ServiceAccount.yaml") {
		item := items.NewOpServiceAccount(v14.ServiceAccount{})
		err := item.LoadFromFile(file, envs)
		return item, err
	}
	if strings.HasSuffix(file, "StatefulSet.yaml") {
		item := items.NewOpStatefulSet(v15.StatefulSet{})
		err := item.LoadFromFile(file, envs)
		return item, err
	}
	if strings.HasSuffix(file, "Role.yaml") {
		item := items.NewOpRole(v17.Role{})
		err := item.LoadFromFile(file, envs)
		return item, err
	}
	if strings.HasSuffix(file, "RoleBinding.yaml") {
		item := items.NewOpRoleBinding(v17.RoleBinding{})
		err := item.LoadFromFile(file, envs)
		return item, err
	}
	if strings.HasSuffix(file, "ConfigMap.yaml") {
		item := items.NewOpConfigMap(v14.ConfigMap{})
		err := item.LoadFromFile(file, envs)
		return item, err
	}
	if strings.HasSuffix(file, "PersistentVolumeClaim.yaml") {
		item := items.NewOpPersistentVolumeClaim(v14.PersistentVolumeClaim{})
		err := item.LoadFromFile(file, envs)
		return item, err
	}
	return nil, fmt.Errorf("unknown kind in file [%s]\n", file)
}
