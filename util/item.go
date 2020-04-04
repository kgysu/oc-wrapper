package util

import (
	"fmt"
	"github.com/kgysu/oc-wrapper/items"
	"github.com/kgysu/oc-wrapper/project"
	v1 "github.com/openshift/api/apps/v1"
	v12 "k8s.io/api/core/v1"
	"strings"
)

// TODO add new Types
func NewOpenshiftItemFromFile(file string) (project.OpenshiftItemInterface, error) {
	if strings.HasSuffix(file, "DeploymentConfig.yaml") {
		odc := items.NewOpDeploymentConfig(&v1.DeploymentConfig{})
		err := odc.LoadFromFile(file)
		return odc, err
	}
	if strings.HasSuffix(file, "Service.yaml") {
		oSvc := items.NewOpService(&v12.Service{})
		err := oSvc.LoadFromFile(file)
		return oSvc, err
	}
	return nil, fmt.Errorf("unknown kind in file [%s]", file)
}
