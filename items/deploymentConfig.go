package items

import (
	"fmt"
	"github.com/kgysu/oc-wrapper/client"
	"github.com/kgysu/oc-wrapper/converter"
	"github.com/kgysu/oc-wrapper/fileutils"
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
	return fileutils.CreateFile(file, []byte(yamlContent))
}

func (oDeploymentConfig OpDeploymentConfig) LoadFromFile(file string, envs map[string]string) error {
	tempData, err := fileutils.ReadFile(file)
	if err != nil {
		return err
	}
	data := fileutils.ReplaceEnvs(string(tempData), envs)
	err = oDeploymentConfig.FromData([]byte(data))
	if err != nil {
		return err
	}
	return nil
}

func (oDeploymentConfig *OpDeploymentConfig) Get(namespace string, restConf *rest.Config, name string) error {
	DeploymentConfigInterface, err := client.GetDeploymentConfigsInterface(namespace, restConf)
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
	DeploymentConfigInterface, err := client.GetDeploymentConfigsInterface(namespace, restConf)
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
	DeploymentConfigInterface, err := client.GetDeploymentConfigsInterface(namespace, restConf)
	if err != nil {
		return err
	}
	err = DeploymentConfigInterface.Delete(oDeploymentConfig.DeploymentConfig.Name, options)
	if err != nil {
		return err
	}
	return nil
}

func (oDeploymentConfig OpDeploymentConfig) UpdateScale(replicas int32, namespace string, restConf *rest.Config) error {
	DeploymentConfigInterface, err := client.GetDeploymentConfigsInterface(namespace, restConf)
	if err != nil {
		return err
	}
	toUpdate, err := DeploymentConfigInterface.Get(oDeploymentConfig.GetName(), v12.GetOptions{})
	if err != nil {
		return err
	}
	toUpdate.Spec.Replicas = replicas
	_, err = DeploymentConfigInterface.Update(toUpdate)
	if err != nil {
		return err
	}
	return nil
}

func (oDeploymentConfig OpDeploymentConfig) GetScale() int32 {
	return oDeploymentConfig.DeploymentConfig.Spec.Replicas
}

func (oDeploymentConfig OpDeploymentConfig) String() string {
	return fmt.Sprintf("%s %s", oDeploymentConfig.Info(), oDeploymentConfig.Status())
}

func (oDeploymentConfig OpDeploymentConfig) Info() string {
	return fmt.Sprintf("[%s] %s",
		oDeploymentConfig.GetKind(),
		oDeploymentConfig.GetName())
}

func (oDeploymentConfig OpDeploymentConfig) Status() string {
	return fmt.Sprintf("%d (%d/%d)",
		oDeploymentConfig.DeploymentConfig.Spec.Replicas,
		oDeploymentConfig.DeploymentConfig.Status.ReadyReplicas,
		oDeploymentConfig.DeploymentConfig.Status.AvailableReplicas)
}

func (oDeploymentConfig OpDeploymentConfig) InfoStatusHtml() string {
	replicasStatus := "warning"
	if oDeploymentConfig.DeploymentConfig.Spec.Replicas == oDeploymentConfig.DeploymentConfig.Status.ReadyReplicas {
		replicasStatus = "success"
	}
	readyStatus := "warning"
	if oDeploymentConfig.DeploymentConfig.Status.AvailableReplicas == oDeploymentConfig.DeploymentConfig.Status.ReadyReplicas {
		readyStatus = "success"
	}
	return fmt.Sprintf(`<b>%s</b> <span class="badge badge-info">%s</span> <button type="button" class="btn btn-sm btn-%s">
  Replicas <span class="badge badge-light">%d</span>
</button> <button type="button" class="btn btn-sm btn-%s">
  Status <span class="badge badge-light">(%d/%d)</span>
</button>
`,
		oDeploymentConfig.GetName(),
		oDeploymentConfig.GetKind(),
		replicasStatus,
		oDeploymentConfig.DeploymentConfig.Spec.Replicas,
		readyStatus,
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
