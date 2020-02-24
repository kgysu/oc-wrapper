package wrapper

import (
	v12 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PersistentVolumeClaimList []v12.PersistentVolumeClaim

func ListPersistentVolumeClaims(ns string, options v1.ListOptions) (PersistentVolumeClaimList, error) {
	persistentVolumeClaimsApi, err := GetPersistentVolumeClaimsApi(ns)
	if err != nil {
		return nil, err
	}
	persistentVolumeClaims, err := persistentVolumeClaimsApi.List(options)
	if err != nil {
		return nil, err
	}
	return persistentVolumeClaims.Items, nil
}

func GetPersistentVolumeClaimByName(ns string, name string, options v1.GetOptions) (*v12.PersistentVolumeClaim, error) {
	persistentVolumeClaimsApi, err := GetPersistentVolumeClaimsApi(ns)
	if err != nil {
		return nil, err
	}
	return persistentVolumeClaimsApi.Get(name, options)
}

func UpdatePersistentVolumeClaim(ns string, persistentVolumeClaim *v12.PersistentVolumeClaim) (*v12.PersistentVolumeClaim, error) {
	persistentVolumeClaimsApi, err := GetPersistentVolumeClaimsApi(ns)
	if err != nil {
		return nil, err
	}
	return persistentVolumeClaimsApi.Update(persistentVolumeClaim)
}

func CreatePersistentVolumeClaim(ns string, persistentVolumeClaim *v12.PersistentVolumeClaim) (*v12.PersistentVolumeClaim, error) {
	persistentVolumeClaimsApi, err := GetPersistentVolumeClaimsApi(ns)
	if err != nil {
		return nil, err
	}
	return persistentVolumeClaimsApi.Create(persistentVolumeClaim)
}

func DeletePersistentVolumeClaim(ns string, name string, options v1.DeleteOptions) error {
	persistentVolumeClaimsApi, err := GetPersistentVolumeClaimsApi(ns)
	if err != nil {
		return err
	}
	return persistentVolumeClaimsApi.Delete(name, &options)
}

func GetPersistentVolumeClaimJson(ns string, name string, options v1.GetOptions) (string, error) {
	persistentVolumeClaim, err := GetPersistentVolumeClaimByName(ns, name, options)
	if err != nil {
		return "", err
	}
	persistentVolumeClaimData, err := ObjectToJsonString(persistentVolumeClaim)
	if err != nil {
		return "", err
	}
	return string(persistentVolumeClaimData), nil
}
