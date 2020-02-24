package wrapper

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
)

type ReplicationControllerList []v1.ReplicationController

func ListReplicationControllers(ns string, options metav1.ListOptions) (ReplicationControllerList, error) {
	rcApi, err := GetReplicationControllerApi(ns)
	if err != nil {
		return nil, err
	}
	rcs, err := rcApi.List(options)
	if err != nil {
		return nil, err
	}
	return rcs.Items, nil
}

func GetReplicationControllerByName(ns string, name string, options metav1.GetOptions) (*v1.ReplicationController, error) {
	rcApi, err := GetReplicationControllerApi(ns)
	if err != nil {
		return nil, err
	}
	return rcApi.Get(name, options)
}

func UpdateReplicationController(ns string, dc *v1.ReplicationController) (*v1.ReplicationController, error) {
	rcApi, err := GetReplicationControllerApi(ns)
	if err != nil {
		return nil, err
	}
	return rcApi.Update(dc)
}

func CreateReplicationController(ns string, dc *v1.ReplicationController) (*v1.ReplicationController, error) {
	rcApi, err := GetReplicationControllerApi(ns)
	if err != nil {
		return nil, err
	}
	return rcApi.Create(dc)
}

func DeleteReplicationController(ns string, name string, options metav1.DeleteOptions) error {
	rcApi, err := GetReplicationControllerApi(ns)
	if err != nil {
		return err
	}
	return rcApi.Delete(name, &options)
}

func WatchReplicationController(ns string, options metav1.ListOptions) (watch.Interface, error) {
	rcApi, err := GetReplicationControllerApi(ns)
	if err != nil {
		return nil, err
	}
	watcher, err := rcApi.Watch(options)
	if err != nil {
		return nil, err
	}
	return watcher, nil
}

func GetReplicationControllerJson(ns string, name string, options metav1.GetOptions) (string, error) {
	rc, err := GetReplicationControllerByName(ns, name, options)
	if err != nil {
		return "", err
	}
	rcData, err := ObjectToJsonString(rc)
	if err != nil {
		return "", err
	}
	return string(rcData), nil
}
