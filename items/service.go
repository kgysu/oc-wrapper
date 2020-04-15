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

var OpServiceTypeMeta = v12.TypeMeta{
	Kind:       "Service",
	APIVersion: "v1",
}

type OpService struct {
	Service *v1.Service
}

func NewOpService(Service v1.Service) *OpService {
	if Service.TypeMeta.Kind != OpServiceTypeMeta.Kind {
		Service.TypeMeta = OpServiceTypeMeta
	}
	return &OpService{
		Service: &Service,
	}
}

// Methods

func (oService *OpService) GetFileName() string {
	return fmt.Sprintf("%s-%s.yaml", oService.GetName(), oService.GetKind())
}

func (oService *OpService) WriteToFile(file string) error {
	yamlContent, err := oService.ToYaml()
	if err != nil {
		return err
	}
	return fileutils.CreateFile(file, []byte(yamlContent))
}

func (oService *OpService) LoadFromFile(file string, envs map[string]string) error {
	tempData, err := fileutils.ReadFile(file)
	if err != nil {
		return err
	}
	data := fileutils.ReplaceEnvs(string(tempData), envs)
	err = oService.FromData([]byte(data))
	if err != nil {
		return err
	}
	return nil
}

func (oService *OpService) Get(namespace string, restConf *rest.Config, name string) error {
	ServiceInterface, err := client.GetServicesInterface(namespace, restConf)
	if err != nil {
		return err
	}
	Service, err := ServiceInterface.Get(name, v12.GetOptions{})
	if err != nil {
		return err
	}
	oService.Service = Service
	return nil
}

func (oService *OpService) Create(namespace string, restConf *rest.Config) error {
	ServiceInterface, err := client.GetServicesInterface(namespace, restConf)
	if err != nil {
		return err
	}
	_, err = ServiceInterface.Create(oService.Service)
	if err != nil {
		return err
	}
	return nil
}

func (oService *OpService) Delete(namespace string, restConf *rest.Config, options *v12.DeleteOptions) error {
	ServiceInterface, err := client.GetServicesInterface(namespace, restConf)
	if err != nil {
		return err
	}
	err = ServiceInterface.Delete(oService.Service.Name, options)
	if err != nil {
		return err
	}
	return nil
}

func (oService OpService) Update(namespace string, restConf *rest.Config) error {
	ServiceInterface, err := client.GetServicesInterface(namespace, restConf)
	if err != nil {
		return err
	}
	toUpdate, err := ServiceInterface.Get(oService.GetName(), v12.GetOptions{})
	if err != nil {
		return err
	}
	toUpdate.Spec = oService.Service.Spec
	_, err = ServiceInterface.Update(toUpdate)
	if err != nil {
		return err
	}
	return nil
}

func (oService *OpService) UpdateScale(replicas int32, namespace string, restConf *rest.Config) error {
	return nil
}

func (oService *OpService) GetScale() int32 {
	return 0
}

func (oService *OpService) IsScalable() bool {
	return false
}

func (oService *OpService) String() string {
	return fmt.Sprintf("%s %s ", oService.Info(), oService.Status())
}

func (oService *OpService) Info() string {
	return fmt.Sprintf("[%s] %s ",
		oService.GetKind(),
		oService.GetName())
}

// TODO more infos
func (oService *OpService) Status() string {
	ports := ""
	for _, port := range oService.Service.Spec.Ports {
		ports = ports + ":" + port.Name
	}
	return fmt.Sprintf("%s %s (%v) [%s]",
		oService.Service.Spec.ClusterIP,
		oService.Service.Spec.Type,
		oService.Service.Spec.PublishNotReadyAddresses,
		ports)

}

func (oService OpService) InfoStatusHtml() string {
	ports := ""
	for _, port := range oService.Service.Spec.Ports {
		ports = ports + createBadge("secondary", fmt.Sprintf("%s (%d)", port.Name, port.Port))
	}
	return fmt.Sprint(
		createInfo(oService.GetKind(), oService.GetName()),
		createLabelBadges(oService.Service.Labels),
		createStatusButton("secondary", fmt.Sprint("ClusterIP ",
			createBadge("light", oService.Service.Spec.ClusterIP))),
		createStatusButton("secondary", fmt.Sprint("LoadBalancerIP ",
			createBadge("light", oService.Service.Spec.LoadBalancerIP))),
		createStatusButton("secondary", fmt.Sprint("Type ",
			createBadge("light", fmt.Sprint(oService.Service.Spec.Type)))),
		createStatusButton("secondary", fmt.Sprint("PublishNotReadyAddresses ",
			createBadge("light", fmt.Sprintf("%v", oService.Service.Spec.PublishNotReadyAddresses)))),
	)
}

func (oService *OpService) GetName() string {
	return oService.Service.Name
}

func (oService *OpService) GetKind() string {
	return oService.Service.Kind
}

func (oService *OpService) ToYaml() (string, error) {
	var contentBuilder strings.Builder
	err := converter.ObjToYaml(oService.Service, &contentBuilder, true, true)
	if err != nil {
		return "", err
	}
	return contentBuilder.String(), nil
}

func (oService *OpService) FromData(data []byte) error {
	_, _, err := converter.YamlToObject(data, false, oService.Service)
	if err != nil {
		return err
	}
	return nil
}
