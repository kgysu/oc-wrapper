package wrapper

import (
	v12 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ServiceAccountList []v12.ServiceAccount

func ListServiceAccounts(ns string, options v1.ListOptions) (ServiceAccountList, error) {
	serviceAccountsApi, err := GetServiceAccountApi(ns)
	if err != nil {
		return nil, err
	}
	serviceAccounts, err := serviceAccountsApi.List(options)
	if err != nil {
		return nil, err
	}
	return serviceAccounts.Items, nil
}

func GetServiceAccountByName(ns string, name string, options v1.GetOptions) (*v12.ServiceAccount, error) {
	serviceAccountsApi, err := GetServiceAccountApi(ns)
	if err != nil {
		return nil, err
	}
	return serviceAccountsApi.Get(name, options)
}

func UpdateServiceAccount(ns string, serviceAccount *v12.ServiceAccount) (*v12.ServiceAccount, error) {
	serviceAccountsApi, err := GetServiceAccountApi(ns)
	if err != nil {
		return nil, err
	}
	return serviceAccountsApi.Update(serviceAccount)
}

func CreateServiceAccount(ns string, serviceAccount *v12.ServiceAccount) (*v12.ServiceAccount, error) {
	serviceAccountsApi, err := GetServiceAccountApi(ns)
	if err != nil {
		return nil, err
	}
	return serviceAccountsApi.Create(serviceAccount)
}

func DeleteServiceAccount(ns string, name string, options v1.DeleteOptions) error {
	serviceAccountsApi, err := GetServiceAccountApi(ns)
	if err != nil {
		return err
	}
	return serviceAccountsApi.Delete(name, &options)
}

func GetServiceAccountJson(ns string, name string, options v1.GetOptions) (string, error) {
	serviceAccount, err := GetServiceAccountByName(ns, name, options)
	if err != nil {
		return "", err
	}
	serviceAccountData, err := ObjectToJsonString(serviceAccount)
	if err != nil {
		return "", err
	}
	return string(serviceAccountData), nil
}
