package project

import v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type OpenshiftItem interface {
	GetName() string
	GetKind() string
	GetStatus() string
	Create(string) error
	Update(string) error
	Delete(string, v1.DeleteOptions) error
}
