package project

import (
	"fmt"
	"github.com/kgysu/oc-wrapper/wrapper"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type OpenshiftPersistentVolumeClaim struct {
	name                  string
	persistentVolumeClaim v1.PersistentVolumeClaim
}

func fromPersistentVolumeClaim(persistentVolumeClaim v1.PersistentVolumeClaim) OpenshiftPersistentVolumeClaim {
	return OpenshiftPersistentVolumeClaim{
		name:                  persistentVolumeClaim.Name,
		persistentVolumeClaim: persistentVolumeClaim,
	}
}

func (oPersistentVolumeClaim OpenshiftPersistentVolumeClaim) setPersistentVolumeClaim(persistentVolumeClaim v1.PersistentVolumeClaim) {
	oPersistentVolumeClaim.name = persistentVolumeClaim.Name
	oPersistentVolumeClaim.persistentVolumeClaim = persistentVolumeClaim
}

func (oPersistentVolumeClaim OpenshiftPersistentVolumeClaim) GetName() string {
	return oPersistentVolumeClaim.name
}

func (oPersistentVolumeClaim OpenshiftPersistentVolumeClaim) GetKind() string {
	return PvClaimKey
}

func (oPersistentVolumeClaim OpenshiftPersistentVolumeClaim) GetStatus() string {
	return fmt.Sprintf("%s", oPersistentVolumeClaim.persistentVolumeClaim.Status.Phase)
}

func (oPersistentVolumeClaim OpenshiftPersistentVolumeClaim) GetPersistentVolumeClaim() v1.PersistentVolumeClaim {
	return oPersistentVolumeClaim.persistentVolumeClaim
}

func (oPersistentVolumeClaim OpenshiftPersistentVolumeClaim) Create(namespace string) error {
	_, err := wrapper.CreatePersistentVolumeClaim(namespace, &oPersistentVolumeClaim.persistentVolumeClaim)
	if err != nil {
		return err
	}
	//oPersistentVolumeClaim.setPersistentVolumeClaim(createdPersistentVolumeClaim)
	return nil
}

func (oPersistentVolumeClaim OpenshiftPersistentVolumeClaim) Update(namespace string) error {
	_, err := wrapper.UpdatePersistentVolumeClaim(namespace, &oPersistentVolumeClaim.persistentVolumeClaim)
	if err != nil {
		return err
	}
	//oPersistentVolumeClaim.setPersistentVolumeClaim(updatedPersistentVolumeClaim)
	return nil
}

func (oPersistentVolumeClaim OpenshiftPersistentVolumeClaim) Delete(namespace string, options v12.DeleteOptions) error {
	return wrapper.DeletePersistentVolumeClaim(namespace, oPersistentVolumeClaim.name, options)
}
