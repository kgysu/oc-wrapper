package project

import (
	"github.com/kgysu/oc-wrapper/wrapper"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type OpenshiftService struct {
	name    string
	service v1.Service
}

func fromService(service v1.Service) OpenshiftService {
	return OpenshiftService{
		name:    service.Name,
		service: service,
	}
}

func (oService OpenshiftService) setService(service v1.Service) {
	oService.name = service.Name
	oService.service = service
}

func (oService OpenshiftService) GetName() string {
	return oService.name
}

func (oService OpenshiftService) GetKind() string {
	return ServiceKey
}

func (oService OpenshiftService) GetStatus() string {
	return oService.service.CreationTimestamp.String()
}

func (oService OpenshiftService) GetService() v1.Service {
	return oService.service
}

func (oService OpenshiftService) Create(namespace string) error {
	_, err := wrapper.CreateService(namespace, &oService.service)
	if err != nil {
		return err
	}
	//oService.setService(createdService)
	return nil
}

func (oService OpenshiftService) Update(namespace string) error {
	_, err := wrapper.UpdateService(namespace, &oService.service)
	if err != nil {
		return err
	}
	//oService.setService(updatedService)
	return nil
}

func (oService OpenshiftService) Delete(namespace string, options v12.DeleteOptions) error {
	return wrapper.DeleteService(namespace, oService.name, options)
}
