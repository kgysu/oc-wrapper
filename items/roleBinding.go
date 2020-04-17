package items

import (
	"fmt"
	"github.com/kgysu/oc-wrapper/client"
	"github.com/kgysu/oc-wrapper/converter"
	"github.com/kgysu/oc-wrapper/fileutils"
	v1 "github.com/openshift/api/authorization/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"strings"
)

var OpRoleBindingTypeMeta = v12.TypeMeta{
	Kind:       "RoleBinding",
	APIVersion: "authorization.openshift.io/v1",
}

type OpRoleBinding struct {
	RoleBinding *v1.RoleBinding
}

func NewOpRoleBinding(RoleBinding v1.RoleBinding) *OpRoleBinding {
	if RoleBinding.TypeMeta.Kind != OpRoleBindingTypeMeta.Kind {
		RoleBinding.TypeMeta = OpRoleBindingTypeMeta
	}
	return &OpRoleBinding{
		RoleBinding: &RoleBinding,
	}
}

// Methods

func (oRoleBinding *OpRoleBinding) GetFileName() string {
	return fmt.Sprintf("%s-%s.yaml", oRoleBinding.GetName(), oRoleBinding.GetKind())
}

func (oRoleBinding *OpRoleBinding) WriteToFile(file string) error {
	yamlContent, err := oRoleBinding.ToYaml()
	if err != nil {
		return err
	}
	return fileutils.CreateFile(file, []byte(yamlContent))
}

func (oRoleBinding *OpRoleBinding) LoadFromFile(file string, envs map[string]string) error {
	tempData, err := fileutils.ReadFile(file)
	if err != nil {
		return err
	}
	data := fileutils.ReplaceEnvs(string(tempData), envs)
	err = oRoleBinding.FromData([]byte(data))
	if err != nil {
		return err
	}
	return nil
}

func (oRoleBinding *OpRoleBinding) Get(namespace string, restConf *rest.Config, name string) error {
	RoleBindingInterface, err := client.GetRoleBindingsInterface(namespace, restConf)
	if err != nil {
		return err
	}
	RoleBinding, err := RoleBindingInterface.Get(name, v12.GetOptions{})
	if err != nil {
		return err
	}
	oRoleBinding.RoleBinding = RoleBinding
	return nil
}

func (oRoleBinding *OpRoleBinding) Create(namespace string, restConf *rest.Config) error {
	RoleBindingInterface, err := client.GetRoleBindingsInterface(namespace, restConf)
	if err != nil {
		return err
	}
	_, err = RoleBindingInterface.Create(oRoleBinding.RoleBinding)
	if err != nil {
		return err
	}
	return nil
}

func (oRoleBinding *OpRoleBinding) Delete(namespace string, restConf *rest.Config, options *v12.DeleteOptions) error {
	RoleBindingInterface, err := client.GetRoleBindingsInterface(namespace, restConf)
	if err != nil {
		return err
	}
	err = RoleBindingInterface.Delete(oRoleBinding.RoleBinding.Name, options)
	if err != nil {
		return err
	}
	return nil
}

func (oRoleBinding OpRoleBinding) Update(namespace string, restConf *rest.Config) error {
	RoleBindingInterface, err := client.GetRoleBindingsInterface(namespace, restConf)
	if err != nil {
		return err
	}
	toUpdate, err := RoleBindingInterface.Get(oRoleBinding.GetName(), v12.GetOptions{})
	if err != nil {
		return err
	}
	toUpdate.RoleRef = oRoleBinding.RoleBinding.RoleRef
	toUpdate.Subjects = oRoleBinding.RoleBinding.Subjects
	toUpdate.GroupNames = oRoleBinding.RoleBinding.GroupNames
	toUpdate.UserNames = oRoleBinding.RoleBinding.UserNames
	toUpdate.Name = oRoleBinding.RoleBinding.Name
	toUpdate.Labels = oRoleBinding.RoleBinding.Labels
	_, err = RoleBindingInterface.Update(toUpdate)
	if err != nil {
		return err
	}
	return nil
}

func (oRoleBinding *OpRoleBinding) UpdateScale(replicas int32, namespace string, restConf *rest.Config) error {
	return nil
}

func (oRoleBinding *OpRoleBinding) GetScale() int32 {
	return 0
}

func (oRoleBinding *OpRoleBinding) IsScalable() bool {
	return false
}

func (oRoleBinding *OpRoleBinding) String() string {
	return fmt.Sprintf("%s %s ", oRoleBinding.Info(), oRoleBinding.Status())
}

func (oRoleBinding *OpRoleBinding) Info() string {
	return fmt.Sprintf("[%s] %s ",
		oRoleBinding.GetKind(),
		oRoleBinding.GetName())
}

// TODO more infos
func (oRoleBinding *OpRoleBinding) Status() string {
	subjects := ""
	for _, subject := range oRoleBinding.RoleBinding.Subjects {
		subjects = subjects + ":" + subject.Name
	}
	return fmt.Sprintf("(%d) %s",
		len(oRoleBinding.RoleBinding.Subjects),
		subjects)

}

func (oRoleBinding OpRoleBinding) InfoStatusHtml() string {
	return fmt.Sprint(
		createInfo(oRoleBinding.GetKind(), oRoleBinding.GetName()),
		createLabelBadges(oRoleBinding.RoleBinding.Labels),
		createStatusButton("secondary", fmt.Sprint("RoleBindings ",
			createBadge("light", fmt.Sprintf("%d", len(oRoleBinding.RoleBinding.Subjects))))),
	)
}

func (oRoleBinding *OpRoleBinding) GetName() string {
	return oRoleBinding.RoleBinding.Name
}

func (oRoleBinding *OpRoleBinding) GetKind() string {
	return oRoleBinding.RoleBinding.Kind
}

func (oRoleBinding *OpRoleBinding) ToYaml() (string, error) {
	var contentBuilder strings.Builder
	err := converter.ObjToYaml(oRoleBinding.RoleBinding, &contentBuilder, true, true)
	if err != nil {
		return "", err
	}
	return contentBuilder.String(), nil
}

func (oRoleBinding *OpRoleBinding) FromData(data []byte) error {
	_, _, err := converter.YamlToObject(data, false, oRoleBinding.RoleBinding)
	if err != nil {
		return err
	}
	return nil
}
