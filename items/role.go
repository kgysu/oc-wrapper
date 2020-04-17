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

var OpRoleTypeMeta = v12.TypeMeta{
	Kind:       "Role",
	APIVersion: "authorization.openshift.io/v1",
}

type OpRole struct {
	Role *v1.Role
}

func NewOpRole(Role v1.Role) *OpRole {
	if Role.TypeMeta.Kind != OpRoleTypeMeta.Kind {
		Role.TypeMeta = OpRoleTypeMeta
	}
	return &OpRole{
		Role: &Role,
	}
}

// Methods

func (oRole *OpRole) GetFileName() string {
	return fmt.Sprintf("%s-%s.yaml", oRole.GetName(), oRole.GetKind())
}

func (oRole *OpRole) WriteToFile(file string) error {
	yamlContent, err := oRole.ToYaml()
	if err != nil {
		return err
	}
	return fileutils.CreateFile(file, []byte(yamlContent))
}

func (oRole *OpRole) LoadFromFile(file string, envs map[string]string) error {
	tempData, err := fileutils.ReadFile(file)
	if err != nil {
		return err
	}
	data := fileutils.ReplaceEnvs(string(tempData), envs)
	err = oRole.FromData([]byte(data))
	if err != nil {
		return err
	}
	return nil
}

func (oRole *OpRole) Get(namespace string, restConf *rest.Config, name string) error {
	RoleInterface, err := client.GetRolesInterface(namespace, restConf)
	if err != nil {
		return err
	}
	Role, err := RoleInterface.Get(name, v12.GetOptions{})
	if err != nil {
		return err
	}
	oRole.Role = Role
	return nil
}

func (oRole *OpRole) Create(namespace string, restConf *rest.Config) error {
	RoleInterface, err := client.GetRolesInterface(namespace, restConf)
	if err != nil {
		return err
	}
	_, err = RoleInterface.Create(oRole.Role)
	if err != nil {
		return err
	}
	return nil
}

func (oRole *OpRole) Delete(namespace string, restConf *rest.Config, options *v12.DeleteOptions) error {
	RoleInterface, err := client.GetRolesInterface(namespace, restConf)
	if err != nil {
		return err
	}
	err = RoleInterface.Delete(oRole.Role.Name, options)
	if err != nil {
		return err
	}
	return nil
}

func (oRole OpRole) Update(namespace string, restConf *rest.Config) error {
	RoleInterface, err := client.GetRolesInterface(namespace, restConf)
	if err != nil {
		return err
	}
	toUpdate, err := RoleInterface.Get(oRole.GetName(), v12.GetOptions{})
	if err != nil {
		return err
	}
	toUpdate.Rules = oRole.Role.Rules
	toUpdate.Name = oRole.Role.Name
	toUpdate.Labels = oRole.Role.Labels
	_, err = RoleInterface.Update(toUpdate)
	if err != nil {
		return err
	}
	return nil
}

func (oRole *OpRole) UpdateScale(replicas int32, namespace string, restConf *rest.Config) error {
	return nil
}

func (oRole *OpRole) GetScale() int32 {
	return 0
}

func (oRole *OpRole) IsScalable() bool {
	return false
}

func (oRole *OpRole) String() string {
	return fmt.Sprintf("%s %s ", oRole.Info(), oRole.Status())
}

func (oRole *OpRole) Info() string {
	return fmt.Sprintf("[%s] %s ",
		oRole.GetKind(),
		oRole.GetName())
}

// TODO more infos
func (oRole *OpRole) Status() string {
	rules := ""
	for _, rule := range oRole.Role.Rules {
		rules = rules + ":" + rule.String()
	}
	return fmt.Sprintf("(%d) %s",
		len(oRole.Role.Rules),
		rules)

}

func (oRole OpRole) InfoStatusHtml() string {
	return fmt.Sprint(
		createInfo(oRole.GetKind(), oRole.GetName()),
		createLabelBadges(oRole.Role.Labels),
		createStatusButton("secondary", fmt.Sprint("Roles ",
			createBadge("light", fmt.Sprintf("%d", len(oRole.Role.Rules))))),
	)
}

func (oRole *OpRole) GetName() string {
	return oRole.Role.Name
}

func (oRole *OpRole) GetKind() string {
	return oRole.Role.Kind
}

func (oRole *OpRole) ToYaml() (string, error) {
	var contentBuilder strings.Builder
	err := converter.ObjToYaml(oRole.Role, &contentBuilder, true, true)
	if err != nil {
		return "", err
	}
	return contentBuilder.String(), nil
}

func (oRole *OpRole) FromData(data []byte) error {
	_, _, err := converter.YamlToObject(data, false, oRole.Role)
	if err != nil {
		return err
	}
	return nil
}
