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

var OpServiceAccountTypeMeta = v12.TypeMeta{
	Kind:       "ServiceAccount",
	APIVersion: "v1",
}

type OpServiceAccount struct {
	ServiceAccount *v1.ServiceAccount
}

func NewOpServiceAccount(ServiceAccount v1.ServiceAccount) *OpServiceAccount {
	if ServiceAccount.TypeMeta.Kind != OpServiceAccountTypeMeta.Kind {
		ServiceAccount.TypeMeta = OpServiceAccountTypeMeta
	}
	return &OpServiceAccount{
		ServiceAccount: &ServiceAccount,
	}
}

// Methods

func (oServiceAccount *OpServiceAccount) GetFileName() string {
	return fmt.Sprintf("%s-%s.yaml", oServiceAccount.GetName(), oServiceAccount.GetKind())
}

func (oServiceAccount *OpServiceAccount) WriteToFile(file string) error {
	yamlContent, err := oServiceAccount.ToYaml()
	if err != nil {
		return err
	}
	return fileutils.CreateFile(file, []byte(yamlContent))
}

func (oServiceAccount *OpServiceAccount) LoadFromFile(file string, envs map[string]string) error {
	tempData, err := fileutils.ReadFile(file)
	if err != nil {
		return err
	}
	data := fileutils.ReplaceEnvs(string(tempData), envs)
	err = oServiceAccount.FromData([]byte(data))
	if err != nil {
		return err
	}
	return nil
}

func (oServiceAccount *OpServiceAccount) Get(namespace string, restConf *rest.Config, name string) error {
	ServiceAccountInterface, err := client.GetServiceAccountsInterface(namespace, restConf)
	if err != nil {
		return err
	}
	ServiceAccount, err := ServiceAccountInterface.Get(name, v12.GetOptions{})
	if err != nil {
		return err
	}
	oServiceAccount.ServiceAccount = ServiceAccount
	return nil
}

func (oServiceAccount *OpServiceAccount) Create(namespace string, restConf *rest.Config) error {
	ServiceAccountInterface, err := client.GetServiceAccountsInterface(namespace, restConf)
	if err != nil {
		return err
	}
	_, err = ServiceAccountInterface.Create(oServiceAccount.ServiceAccount)
	if err != nil {
		return err
	}
	return nil
}

func (oServiceAccount *OpServiceAccount) Delete(namespace string, restConf *rest.Config, options *v12.DeleteOptions) error {
	ServiceAccountInterface, err := client.GetServiceAccountsInterface(namespace, restConf)
	if err != nil {
		return err
	}
	err = ServiceAccountInterface.Delete(oServiceAccount.ServiceAccount.Name, options)
	if err != nil {
		return err
	}
	return nil
}

func (oServiceAccount OpServiceAccount) Update(namespace string, restConf *rest.Config) error {
	ServiceAccountInterface, err := client.GetServiceAccountsInterface(namespace, restConf)
	if err != nil {
		return err
	}
	toUpdate, err := ServiceAccountInterface.Get(oServiceAccount.GetName(), v12.GetOptions{})
	if err != nil {
		return err
	}
	toUpdate.ImagePullSecrets = oServiceAccount.ServiceAccount.ImagePullSecrets
	toUpdate.Secrets = oServiceAccount.ServiceAccount.Secrets
	toUpdate.Name = oServiceAccount.ServiceAccount.Name
	toUpdate.Labels = oServiceAccount.ServiceAccount.Labels
	_, err = ServiceAccountInterface.Update(toUpdate)
	if err != nil {
		return err
	}
	return nil
}

func (oServiceAccount *OpServiceAccount) UpdateScale(replicas int32, namespace string, restConf *rest.Config) error {
	return nil
}

func (oServiceAccount *OpServiceAccount) GetScale() int32 {
	return 0
}

func (oServiceAccount *OpServiceAccount) IsScalable() bool {
	return false
}

func (oServiceAccount *OpServiceAccount) String() string {
	return fmt.Sprintf("%s %s ", oServiceAccount.Info(), oServiceAccount.Status())
}

func (oServiceAccount *OpServiceAccount) Info() string {
	return fmt.Sprintf("[%s] %s ",
		oServiceAccount.GetKind(),
		oServiceAccount.GetName())
}

// TODO more infos
func (oServiceAccount *OpServiceAccount) Status() string {
	return fmt.Sprintf("(%d)(%d))",
		len(oServiceAccount.ServiceAccount.Secrets),
		len(oServiceAccount.ServiceAccount.ImagePullSecrets))

}

func (oServiceAccount OpServiceAccount) InfoStatusHtml() string {
	return fmt.Sprint(
		createInfo(oServiceAccount.GetKind(), oServiceAccount.GetName()),
		createLabelBadges(oServiceAccount.ServiceAccount.Labels),
		createStatusButton("secondary", fmt.Sprint("Secrets ",
			createBadge("light", fmt.Sprintf("%d", len(oServiceAccount.ServiceAccount.Secrets))))),
		createStatusButton("secondary", fmt.Sprint("ImagePullSecrets ",
			createBadge("light", fmt.Sprintf("%d", len(oServiceAccount.ServiceAccount.ImagePullSecrets))))),
	)
}

func (oServiceAccount *OpServiceAccount) GetName() string {
	return oServiceAccount.ServiceAccount.Name
}

func (oServiceAccount *OpServiceAccount) GetKind() string {
	return oServiceAccount.ServiceAccount.Kind
}

func (oServiceAccount *OpServiceAccount) ToYaml() (string, error) {
	var contentBuilder strings.Builder
	err := converter.ObjToYaml(oServiceAccount.ServiceAccount, &contentBuilder, true, true)
	if err != nil {
		return "", err
	}
	return contentBuilder.String(), nil
}

func (oServiceAccount *OpServiceAccount) FromData(data []byte) error {
	_, _, err := converter.YamlToObject(data, false, oServiceAccount.ServiceAccount)
	if err != nil {
		return err
	}
	return nil
}
