package items

import (
	"fmt"
	"github.com/kgysu/oc-wrapper/converter"
	"github.com/kgysu/oc-wrapper/files"
	v3 "github.com/kgysu/oc-wrapper/v3"
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
	Service.TypeMeta = OpServiceTypeMeta
	return &OpService{
		Service: &Service,
	}
}

// Methods

func (oService *OpService) GetFileName() string {
	return fmt.Sprintf("%s-%s.yaml", oService.Service.Name, oService.Service.Kind)
}

func (oService *OpService) WriteToFile(file string) error {
	var sb strings.Builder
	err := converter.ObjToYaml(oService.Service, &sb, true, false)
	if err != nil {
		return err
	}
	fileData := []byte(sb.String())
	return files.CreateFile(file, fileData)
}

func (oService *OpService) LoadFromFile(file string, envs map[string]string) error {
	tempData, err := files.ReadFile(file)
	if err != nil {
		return err
	}
	data := files.ReplaceEnvs(string(tempData), envs)
	_, _, err = converter.YamlToObject([]byte(data), false, oService.Service)
	if err != nil {
		return err
	}
	return nil
}

func (oService *OpService) Get(namespace string, restConf *rest.Config, name string) error {
	ServiceInterface, err := v3.GetServicesInterface(namespace, restConf)
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
	ServiceInterface, err := v3.GetServicesInterface(namespace, restConf)
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
	ServiceInterface, err := v3.GetServicesInterface(namespace, restConf)
	if err != nil {
		return err
	}
	err = ServiceInterface.Delete(oService.Service.Name, options)
	if err != nil {
		return err
	}
	return nil
}

func (oService *OpService) String() string {
	return fmt.Sprintf("%s %s \n", oService.Info(), oService.Status())
}

func (oService *OpService) Info() string {
	return fmt.Sprintf("[%s] %s ",
		oService.Service.Kind,
		oService.Service.Name)
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
