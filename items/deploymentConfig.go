package items

import (
	"fmt"
	"github.com/kgysu/oc-wrapper/converter"
	"github.com/kgysu/oc-wrapper/files"
	v3 "github.com/kgysu/oc-wrapper/v3"
	v1 "github.com/openshift/api/apps/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"strings"
)

var OpDeploymentConfigTypeMeta = v12.TypeMeta{
	Kind:       "DeploymentConfig",
	APIVersion: "apps.openshift.io/v1",
}

type OpDeploymentConfig struct {
	dc *v1.DeploymentConfig
}

func NewOpDeploymentConfig(dc *v1.DeploymentConfig) *OpDeploymentConfig {
	dc.TypeMeta = OpDeploymentConfigTypeMeta
	return &OpDeploymentConfig{
		dc: dc,
	}
}

// Methods

func (odc *OpDeploymentConfig) GetFileName() string {
	return fmt.Sprintf("%s-%s.yaml", odc.dc.Name, odc.dc.Kind)
}

func (odc *OpDeploymentConfig) WriteToFile(file string) error {
	var sb strings.Builder
	err := converter.ObjToYaml(odc.dc, &sb, true, false)
	if err != nil {
		return err
	}
	fileData := []byte(sb.String())
	return files.CreateFile(file, fileData)
}

func (odc *OpDeploymentConfig) LoadFromFile(file string) error {
	data, err := files.ReadFile(file)
	if err != nil {
		return err
	}
	_, _, err = converter.YamlToObject(data, false, odc.dc)
	if err != nil {
		return err
	}
	return nil
}

func (odc *OpDeploymentConfig) Get(namespace string, restConf *rest.Config, name string) error {
	dcInterface, err := v3.GetDeploymentConfigsInterface(namespace, restConf)
	if err != nil {
		return err
	}
	dc, err := dcInterface.Get(name, v12.GetOptions{})
	if err != nil {
		return err
	}
	odc.dc = dc
	return nil
}

func (odc *OpDeploymentConfig) Create(namespace string, restConf *rest.Config) error {
	dcInterface, err := v3.GetDeploymentConfigsInterface(namespace, restConf)
	if err != nil {
		return err
	}
	_, err = dcInterface.Create(odc.dc)
	if err != nil {
		return err
	}
	return nil
}

func (odc *OpDeploymentConfig) Delete(namespace string, restConf *rest.Config, options *v12.DeleteOptions) error {
	dcInterface, err := v3.GetDeploymentConfigsInterface(namespace, restConf)
	if err != nil {
		return err
	}
	err = dcInterface.Delete(odc.dc.Name, options)
	if err != nil {
		return err
	}
	return nil
}

func (odc *OpDeploymentConfig) String() string {
	return fmt.Sprintf("%s %s \n", odc.Info(), odc.Status())
}

func (odc *OpDeploymentConfig) Info() string {
	return fmt.Sprintf("[%s] %s ",
		odc.dc.Kind,
		odc.dc.Name)
}

func (odc *OpDeploymentConfig) Status() string {
	return fmt.Sprintf("%d (%d/%d) ",
		odc.dc.Spec.Replicas,
		odc.dc.Status.ReadyReplicas,
		odc.dc.Status.AvailableReplicas)
}
