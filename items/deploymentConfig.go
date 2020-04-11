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
	DeploymentConfig *v1.DeploymentConfig
}

func NewOpDeploymentConfig(DeploymentConfig v1.DeploymentConfig) *OpDeploymentConfig {
	DeploymentConfig.TypeMeta = OpDeploymentConfigTypeMeta
	return &OpDeploymentConfig{
		DeploymentConfig: &DeploymentConfig,
	}
}

// Methods

func (oDeploymentConfig OpDeploymentConfig) GetFileName() string {
	return fmt.Sprintf("%s-%s.yaml", oDeploymentConfig.GetName(), oDeploymentConfig.GetKind())
}

func (oDeploymentConfig OpDeploymentConfig) WriteToFile(file string) error {
	yamlContent, err := oDeploymentConfig.ToYaml()
	if err != nil {
		return err
	}
	return files.CreateFile(file, []byte(yamlContent))
}

func (oDeploymentConfig OpDeploymentConfig) LoadFromFile(file string, envs map[string]string) error {
	tempData, err := files.ReadFile(file)
	if err != nil {
		return err
	}
	data := files.ReplaceEnvs(string(tempData), envs)
	err = oDeploymentConfig.FromData([]byte(data))
	if err != nil {
		return err
	}
	return nil
}

func (oDeploymentConfig *OpDeploymentConfig) Get(namespace string, restConf *rest.Config, name string) error {
	DeploymentConfigInterface, err := v3.GetDeploymentConfigsInterface(namespace, restConf)
	if err != nil {
		return err
	}
	DeploymentConfig, err := DeploymentConfigInterface.Get(name, v12.GetOptions{})
	if err != nil {
		return err
	}
	oDeploymentConfig.DeploymentConfig = DeploymentConfig
	return nil
}

func (oDeploymentConfig OpDeploymentConfig) Create(namespace string, restConf *rest.Config) error {
	DeploymentConfigInterface, err := v3.GetDeploymentConfigsInterface(namespace, restConf)
	if err != nil {
		return err
	}
	_, err = DeploymentConfigInterface.Create(oDeploymentConfig.DeploymentConfig)
	if err != nil {
		return err
	}
	return nil
}

func (oDeploymentConfig OpDeploymentConfig) Delete(namespace string, restConf *rest.Config, options *v12.DeleteOptions) error {
	DeploymentConfigInterface, err := v3.GetDeploymentConfigsInterface(namespace, restConf)
	if err != nil {
		return err
	}
	err = DeploymentConfigInterface.Delete(oDeploymentConfig.DeploymentConfig.Name, options)
	if err != nil {
		return err
	}
	return nil
}

func (oDeploymentConfig OpDeploymentConfig) String() string {
	return fmt.Sprintf("%s %s \n", oDeploymentConfig.Info(), oDeploymentConfig.Status())
}

func (oDeploymentConfig OpDeploymentConfig) Info() string {
	return fmt.Sprintf("[%s] %s ",
		oDeploymentConfig.GetKind(),
		oDeploymentConfig.GetName())
}

func (oDeploymentConfig OpDeploymentConfig) Status() string {
	return fmt.Sprintf("%d (%d/%d) ",
		oDeploymentConfig.DeploymentConfig.Spec.Replicas,
		oDeploymentConfig.DeploymentConfig.Status.ReadyReplicas,
		oDeploymentConfig.DeploymentConfig.Status.AvailableReplicas)
}

func (oDeploymentConfig OpDeploymentConfig) GetName() string {
	return oDeploymentConfig.DeploymentConfig.Name
}

func (oDeploymentConfig OpDeploymentConfig) GetKind() string {
	return oDeploymentConfig.DeploymentConfig.Kind
}

func (oDeploymentConfig *OpDeploymentConfig) ToYaml() (string, error) {
	var contentBuilder strings.Builder
	err := converter.ObjToYaml(oDeploymentConfig.DeploymentConfig, &contentBuilder, true, true)
	if err != nil {
		return "", err
	}
	return contentBuilder.String(), nil
}

func (oDeploymentConfig *OpDeploymentConfig) FromData(data []byte) error {
	_, _, err := converter.YamlToObject(data, false, oDeploymentConfig.DeploymentConfig)
	if err != nil {
		return err
	}
	return nil
}
