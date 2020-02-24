package wrapper

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type SecretList []v1.Secret

func ListSecrets(ns string, options metav1.ListOptions) (SecretList, error) {
	secretsApi, err := GetSecretApi(ns)
	if err != nil {
		return nil, err
	}
	secrets, err := secretsApi.List(options)
	if err != nil {
		return nil, err
	}
	return secrets.Items, nil
}

func GetSecretByName(ns string, name string, options metav1.GetOptions) (*v1.Secret, error) {
	secretsApi, err := GetSecretApi(ns)
	if err != nil {
		return nil, err
	}
	return secretsApi.Get(name, options)
}

func UpdateSecret(ns string, dc *v1.Secret) (*v1.Secret, error) {
	secretsApi, err := GetSecretApi(ns)
	if err != nil {
		return nil, err
	}
	return secretsApi.Update(dc)
}

func CreateSecret(ns string, dc *v1.Secret) (*v1.Secret, error) {
	secretsApi, err := GetSecretApi(ns)
	if err != nil {
		return nil, err
	}
	return secretsApi.Create(dc)
}

func DeleteSecret(ns string, name string, options metav1.DeleteOptions) error {
	secretsApi, err := GetSecretApi(ns)
	if err != nil {
		return err
	}
	return secretsApi.Delete(name, &options)
}

func GetSecretJson(ns string, name string, options metav1.GetOptions) (string, error) {
	secret, err := GetSecretByName(ns, name, options)
	if err != nil {
		return "", err
	}
	secretData, err := ObjectToJsonString(secret)
	if err != nil {
		return "", err
	}
	return string(secretData), nil
}
