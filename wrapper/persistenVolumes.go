package wrapper

import (
	v12 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PersistentVolumeList []v12.PersistentVolume

func ListPersistentVolumes(ns string, options v1.ListOptions) (PersistentVolumeList, error) {
	persistentVolumesApi, err := GetPersistentVolumeApi(ns)
	if err != nil {
		return nil, err
	}
	persistentVolumes, err := persistentVolumesApi.List(options)
	if err != nil {
		return nil, err
	}
	return persistentVolumes.Items, nil
}

func GetPersistentVolumeByName(ns string, name string, options v1.GetOptions) (*v12.PersistentVolume, error) {
	persistentVolumesApi, err := GetPersistentVolumeApi(ns)
	if err != nil {
		return nil, err
	}
	return persistentVolumesApi.Get(name, options)
}

func UpdatePersistentVolume(ns string, persistentVolume *v12.PersistentVolume) (*v12.PersistentVolume, error) {
	persistentVolumesApi, err := GetPersistentVolumeApi(ns)
	if err != nil {
		return nil, err
	}
	return persistentVolumesApi.Update(persistentVolume)
}

func CreatePersistentVolume(ns string, persistentVolume *v12.PersistentVolume) (*v12.PersistentVolume, error) {
	persistentVolumesApi, err := GetPersistentVolumeApi(ns)
	if err != nil {
		return nil, err
	}
	return persistentVolumesApi.Create(persistentVolume)
}

func DeletePersistentVolume(ns string, name string, options v1.DeleteOptions) error {
	persistentVolumesApi, err := GetPersistentVolumeApi(ns)
	if err != nil {
		return err
	}
	return persistentVolumesApi.Delete(name, &options)
}

func GetPersistentVolumeJson(ns string, name string, options v1.GetOptions) (string, error) {
	persistentVolume, err := GetPersistentVolumeByName(ns, name, options)
	if err != nil {
		return "", err
	}
	persistentVolumeData, err := ObjectToJsonString(persistentVolume)
	if err != nil {
		return "", err
	}
	return string(persistentVolumeData), nil
}
