package project

import (
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

type OpenshiftItemInterface interface {
	WriteToFile(file string) error
	LoadFromFile(file string, envs map[string]string) error
	GetFileName() string
	String() string
	Info() string
	Create(namespace string, restConf *rest.Config) error
	Delete(namespace string, restConf *rest.Config, options *v12.DeleteOptions) error
	GetName() string
	GetKind() string
	ToYaml() (string, error)
	FromData(data []byte) error
}
