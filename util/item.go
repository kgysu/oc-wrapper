package util

import (
	"fmt"
	"github.com/kgysu/oc-wrapper/items"
	"github.com/kgysu/oc-wrapper/project"
	v1 "github.com/openshift/api/apps/v1"
	v13 "github.com/openshift/api/route/v1"
	v12 "k8s.io/api/core/v1"
	"strings"
)

// TODO add new Types
func NewOpenshiftItemFromFile(file string, envs map[string]string) (project.OpenshiftItemInterface, error) {

	if strings.HasSuffix(file, "DeploymentConfig.yaml") {
		item := items.NewOpDeploymentConfig(v1.DeploymentConfig{})
		err := item.LoadFromFile(file, envs)
		return item, err
	}
	if strings.HasSuffix(file, "Service.yaml") {
		item := items.NewOpService(v12.Service{})
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
