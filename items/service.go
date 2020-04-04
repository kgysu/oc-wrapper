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
	svc *v1.Service
}

func NewOpService(svc *v1.Service) *OpService {
	svc.TypeMeta = OpServiceTypeMeta
	return &OpService{
		svc: svc,
	}
}

// Methods

func (oSvc *OpService) GetFileName() string {
	return fmt.Sprintf("%s-%s.yaml", oSvc.svc.Name, oSvc.svc.Kind)
}

func (oSvc *OpService) WriteToFile(file string) error {
	var sb strings.Builder
	err := converter.ObjToYaml(oSvc.svc, &sb, true, false)
	if err != nil {
		return err
	}
	fileData := []byte(sb.String())
	return files.CreateFile(file, fileData)
}

func (oSvc *OpService) LoadFromFile(file string) error {
	data, err := files.ReadFile(file)
	if err != nil {
		return err
	}
	_, _, err = converter.YamlToObject(data, false, oSvc.svc)
	if err != nil {
		return err
	}
	return nil
}

func (oSvc *OpService) Get(namespace string, restConf *rest.Config, name string) error {
	svcInterface, err := v3.GetServicesInterface(namespace, restConf)
	if err != nil {
		return err
	}
	svc, err := svcInterface.Get(name, v12.GetOptions{})
	if err != nil {
		return err
	}
	oSvc.svc = svc
	return nil
}

func (oSvc *OpService) Create(namespace string, restConf *rest.Config) error {
	svcInterface, err := v3.GetServicesInterface(namespace, restConf)
	if err != nil {
		return err
	}
	_, err = svcInterface.Create(oSvc.svc)
	if err != nil {
		return err
	}
	return nil
}

func (oSvc *OpService) Delete(namespace string, restConf *rest.Config, options *v12.DeleteOptions) error {
	svcInterface, err := v3.GetServicesInterface(namespace, restConf)
	if err != nil {
		return err
	}
	err = svcInterface.Delete(oSvc.svc.Name, options)
	if err != nil {
		return err
	}
	return nil
}

func (oSvc *OpService) String() string {
	return fmt.Sprintf("%s %s \n", oSvc.Info(), oSvc.Status())
}

func (oSvc *OpService) Info() string {
	return fmt.Sprintf("[%s] %s ",
		oSvc.svc.Kind,
		oSvc.svc.Name)
}

// TODO more infos
func (oSvc *OpService) Status() string {
	ports := ""
	for _, port := range oSvc.svc.Spec.Ports {
		ports = ports + ":" + port.Name
	}
	return fmt.Sprintf("%s %s (%v) [%s]",
		oSvc.svc.Spec.ClusterIP,
		oSvc.svc.Spec.Type,
		oSvc.svc.Spec.PublishNotReadyAddresses,
		ports)

}
