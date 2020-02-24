package project

import (
	"github.com/kgysu/oc-wrapper/wrapper"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type OpenshiftServiceAccount struct {
	name           string
	serviceAccount v1.ServiceAccount
}

func fromServiceAccount(serviceAccount v1.ServiceAccount) OpenshiftServiceAccount {
	return OpenshiftServiceAccount{
		name:           serviceAccount.Name,
		serviceAccount: serviceAccount,
	}
}

func (oServiceAccount OpenshiftServiceAccount) setServiceAccount(serviceAccount v1.ServiceAccount) {
	oServiceAccount.name = serviceAccount.Name
	oServiceAccount.serviceAccount = serviceAccount
}

func (oServiceAccount OpenshiftServiceAccount) GetName() string {
	return oServiceAccount.name
}

func (oServiceAccount OpenshiftServiceAccount) GetKind() string {
	return ServiceAccountKey
}

func (oServiceAccount OpenshiftServiceAccount) GetStatus() string {
	return oServiceAccount.serviceAccount.CreationTimestamp.String()
}

func (oServiceAccount OpenshiftServiceAccount) GetServiceAccount() v1.ServiceAccount {
	return oServiceAccount.serviceAccount
}

func (oServiceAccount OpenshiftServiceAccount) Create(namespace string) error {
	_, err := wrapper.CreateServiceAccount(namespace, &oServiceAccount.serviceAccount)
	if err != nil {
		return err
	}
	//oServiceAccount.setServiceAccount(createdServiceAccount)
	return nil
}

func (oServiceAccount OpenshiftServiceAccount) Update(namespace string) error {
	_, err := wrapper.UpdateServiceAccount(namespace, &oServiceAccount.serviceAccount)
	if err != nil {
		return err
	}
	//oServiceAccount.setServiceAccount(updatedServiceAccount)
	return nil
}

func (oServiceAccount OpenshiftServiceAccount) Delete(namespace string, options v12.DeleteOptions) error {
	return wrapper.DeleteServiceAccount(namespace, oServiceAccount.name, options)
}
