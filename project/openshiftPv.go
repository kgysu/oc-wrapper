package project

import (
	"fmt"
	"github.com/kgysu/oc-wrapper/wrapper"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type OpenshiftPersistentVolume struct {
	name             string
	persistentVolume v1.PersistentVolume
}

func fromPersistentVolume(persistentVolume v1.PersistentVolume) OpenshiftPersistentVolume {
	return OpenshiftPersistentVolume{
		name:             persistentVolume.Name,
		persistentVolume: persistentVolume,
	}
}

func (oPersistentVolume OpenshiftPersistentVolume) setPersistentVolume(persistentVolume v1.PersistentVolume) {
	oPersistentVolume.name = persistentVolume.Name
	oPersistentVolume.persistentVolume = persistentVolume
}

func (oPersistentVolume OpenshiftPersistentVolume) GetName() string {
	return oPersistentVolume.name
}

func (oPersistentVolume OpenshiftPersistentVolume) GetKind() string {
	return PvKey
}

func (oPersistentVolume OpenshiftPersistentVolume) GetStatus() string {
	return fmt.Sprintf("%s (%s) [%s]", oPersistentVolume.persistentVolume.Status.Phase,
		oPersistentVolume.persistentVolume.Status.Message,
		oPersistentVolume.persistentVolume.Status.Reason)
}

func (oPersistentVolume OpenshiftPersistentVolume) GetPersistentVolume() v1.PersistentVolume {
	return oPersistentVolume.persistentVolume
}

func (oPersistentVolume OpenshiftPersistentVolume) Create(namespace string) error {
	_, err := wrapper.CreatePersistentVolume(namespace, &oPersistentVolume.persistentVolume)
	if err != nil {
		return err
	}
	//oPersistentVolume.setPersistentVolume(createdPersistentVolume)
	return nil
}

func (oPersistentVolume OpenshiftPersistentVolume) Update(namespace string) error {
	_, err := wrapper.UpdatePersistentVolume(namespace, &oPersistentVolume.persistentVolume)
	if err != nil {
		return err
	}
	//oPersistentVolume.setPersistentVolume(updatedPersistentVolume)
	return nil
}

func (oPersistentVolume OpenshiftPersistentVolume) Delete(namespace string, options v12.DeleteOptions) error {
	return wrapper.DeletePersistentVolume(namespace, oPersistentVolume.name, options)
}
