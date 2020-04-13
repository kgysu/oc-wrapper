package project

import (
	"fmt"
	"github.com/kgysu/oc-wrapper/items"
	v1 "github.com/openshift/api/apps/v1"
	v13 "github.com/openshift/api/route/v1"
	v14 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"strings"
)

type OpenshiftItemInterface interface {
	WriteToFile(file string) error
	LoadFromFile(file string, envs map[string]string) error
	GetFileName() string
	String() string
	Info() string
	Status() string
	InfoStatusHtml() string
	Create(namespace string, restConf *rest.Config) error
	Delete(namespace string, restConf *rest.Config, options *v12.DeleteOptions) error
	UpdateScale(replicas int, namespace string, restConf *rest.Config) error
	GetScale() int32
	GetName() string
	GetKind() string
	ToYaml() (string, error)
	FromData(data []byte) error
}

func NewOpenshiftItemFromFile(file string, envs map[string]string) (OpenshiftItemInterface, error) {
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
	return nil, fmt.Errorf("unknown kind in file [%s]", file)
}
