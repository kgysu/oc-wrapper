package util

import (
	items2 "github.com/kgysu/oc-wrapper/items"
	"github.com/kgysu/oc-wrapper/project"
	"github.com/kgysu/oc-wrapper/v3"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

// Todo add more Types
func ListDeploymentConfigs(namespace string, restConf *rest.Config, options v12.ListOptions) ([]project.OpenshiftItemInterface, error) {
	dcInterface, err := v3.GetDeploymentConfigsInterface(namespace, restConf)
	if err != nil {
		return nil, err
	}
	dcs, err := dcInterface.List(options)
	if err != nil {
		return nil, err
	}
	var resultItems []project.OpenshiftItemInterface
	for _, dc := range dcs.Items {
		resultItems = append(resultItems, items2.NewOpDeploymentConfig(&dc))
	}
	return resultItems, nil
}

func ListServices(namespace string, restConf *rest.Config, options v12.ListOptions) ([]project.OpenshiftItemInterface, error) {
	svcInterface, err := v3.GetServicesInterface(namespace, restConf)
	if err != nil {
		return nil, err
	}
	dcs, err := svcInterface.List(options)
	if err != nil {
		return nil, err
	}
	var resultItems []project.OpenshiftItemInterface
	for _, svc := range dcs.Items {
		resultItems = append(resultItems, items2.NewOpService(&svc))
	}
	return resultItems, nil
}
