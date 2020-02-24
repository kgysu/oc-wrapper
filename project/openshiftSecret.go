package project

import (
	"github.com/kgysu/oc-wrapper/wrapper"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type OpenshiftSecret struct {
	name   string
	secret v1.Secret
}

func fromSecret(secret v1.Secret) OpenshiftSecret {
	return OpenshiftSecret{
		name:   secret.Name,
		secret: secret,
	}
}

func (oSecret OpenshiftSecret) setSecret(secret v1.Secret) {
	oSecret.name = secret.Name
	oSecret.secret = secret
}

func (oSecret OpenshiftSecret) GetName() string {
	return oSecret.name
}

func (oSecret OpenshiftSecret) GetKind() string {
	return SecretKey
}

func (oSecret OpenshiftSecret) GetStatus() string {
	return oSecret.secret.CreationTimestamp.String()
}

func (oSecret OpenshiftSecret) GetSecret() v1.Secret {
	return oSecret.secret
}

func (oSecret OpenshiftSecret) Create(namespace string) error {
	_, err := wrapper.CreateSecret(namespace, &oSecret.secret)
	if err != nil {
		return err
	}
	//oSecret.setSecret(createdSecret)
	return nil
}

func (oSecret OpenshiftSecret) Update(namespace string) error {
	_, err := wrapper.UpdateSecret(namespace, &oSecret.secret)
	if err != nil {
		return err
	}
	//oSecret.setSecret(updatedSecret)
	return nil
}

func (oSecret OpenshiftSecret) Delete(namespace string, options v12.DeleteOptions) error {
	return wrapper.DeleteSecret(namespace, oSecret.name, options)
}
