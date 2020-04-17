package items

import (
	"fmt"
	"github.com/kgysu/oc-wrapper/client"
	"github.com/kgysu/oc-wrapper/converter"
	"github.com/kgysu/oc-wrapper/fileutils"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"strings"
)

var OpConfigMapTypeMeta = v12.TypeMeta{
	Kind:       "ConfigMap",
	APIVersion: "v1",
}

type OpConfigMap struct {
	ConfigMap *v1.ConfigMap
}

func NewOpConfigMap(ConfigMap v1.ConfigMap) *OpConfigMap {
	if ConfigMap.TypeMeta.Kind != OpConfigMapTypeMeta.Kind {
		ConfigMap.TypeMeta = OpConfigMapTypeMeta
	}
	return &OpConfigMap{
		ConfigMap: &ConfigMap,
	}
}

// Methods

func (oConfigMap *OpConfigMap) GetFileName() string {
	return fmt.Sprintf("%s-%s.yaml", oConfigMap.GetName(), oConfigMap.GetKind())
}

func (oConfigMap *OpConfigMap) WriteToFile(file string) error {
	yamlContent, err := oConfigMap.ToYaml()
	if err != nil {
		return err
	}
	return fileutils.CreateFile(file, []byte(yamlContent))
}

func (oConfigMap *OpConfigMap) LoadFromFile(file string, envs map[string]string) error {
	tempData, err := fileutils.ReadFile(file)
	if err != nil {
		return err
	}
	data := fileutils.ReplaceEnvs(string(tempData), envs)
	err = oConfigMap.FromData([]byte(data))
	if err != nil {
		return err
	}
	return nil
}

func (oConfigMap *OpConfigMap) Get(namespace string, restConf *rest.Config, name string) error {
	ConfigMapInterface, err := client.GetConfigMapsInterface(namespace, restConf)
	if err != nil {
		return err
	}
	ConfigMap, err := ConfigMapInterface.Get(name, v12.GetOptions{})
	if err != nil {
		return err
	}
	oConfigMap.ConfigMap = ConfigMap
	return nil
}

func (oConfigMap *OpConfigMap) Create(namespace string, restConf *rest.Config) error {
	ConfigMapInterface, err := client.GetConfigMapsInterface(namespace, restConf)
	if err != nil {
		return err
	}
	_, err = ConfigMapInterface.Create(oConfigMap.ConfigMap)
	if err != nil {
		return err
	}
	return nil
}

func (oConfigMap *OpConfigMap) Delete(namespace string, restConf *rest.Config, options *v12.DeleteOptions) error {
	ConfigMapInterface, err := client.GetConfigMapsInterface(namespace, restConf)
	if err != nil {
		return err
	}
	err = ConfigMapInterface.Delete(oConfigMap.ConfigMap.Name, options)
	if err != nil {
		return err
	}
	return nil
}

func (oConfigMap OpConfigMap) Update(namespace string, restConf *rest.Config) error {
	ConfigMapInterface, err := client.GetConfigMapsInterface(namespace, restConf)
	if err != nil {
		return err
	}
	toUpdate, err := ConfigMapInterface.Get(oConfigMap.GetName(), v12.GetOptions{})
	if err != nil {
		return err
	}
	toUpdate.Data = oConfigMap.ConfigMap.Data
	toUpdate.Name = oConfigMap.ConfigMap.Name
	toUpdate.Labels = oConfigMap.ConfigMap.Labels
	_, err = ConfigMapInterface.Update(toUpdate)
	if err != nil {
		return err
	}
	return nil
}

func (oConfigMap *OpConfigMap) UpdateScale(replicas int32, namespace string, restConf *rest.Config) error {
	return nil
}

func (oConfigMap *OpConfigMap) GetScale() int32 {
	return 0
}

func (oConfigMap *OpConfigMap) IsScalable() bool {
	return false
}

func (oConfigMap *OpConfigMap) String() string {
	return fmt.Sprintf("%s %s ", oConfigMap.Info(), oConfigMap.Status())
}

func (oConfigMap *OpConfigMap) Info() string {
	return fmt.Sprintf("[%s] %s ",
		oConfigMap.GetKind(),
		oConfigMap.GetName())
}

// TODO more infos
func (oConfigMap *OpConfigMap) Status() string {
	return fmt.Sprintf("(%d)",
		len(oConfigMap.ConfigMap.Data),
	)

}

func (oConfigMap OpConfigMap) InfoStatusHtml() string {
	return fmt.Sprint(
		createInfo(oConfigMap.GetKind(), oConfigMap.GetName()),
		createLabelBadges(oConfigMap.ConfigMap.Labels),
		createStatusButton("secondary", fmt.Sprint("Files ",
			createBadge("light", fmt.Sprintf("%d", len(oConfigMap.ConfigMap.Data))))),
	)
}

func (oConfigMap *OpConfigMap) GetName() string {
	return oConfigMap.ConfigMap.Name
}

func (oConfigMap *OpConfigMap) GetKind() string {
	return oConfigMap.ConfigMap.Kind
}

func (oConfigMap *OpConfigMap) ToYaml() (string, error) {
	var contentBuilder strings.Builder
	err := converter.ObjToYaml(oConfigMap.ConfigMap, &contentBuilder, true, true)
	if err != nil {
		return "", err
	}
	return contentBuilder.String(), nil
}

func (oConfigMap *OpConfigMap) FromData(data []byte) error {
	_, _, err := converter.YamlToObject(data, false, oConfigMap.ConfigMap)
	if err != nil {
		return err
	}
	return nil
}
