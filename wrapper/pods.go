package wrapper

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
)

type PodList []v1.Pod

// Get all pods in namespace
func ListPods(ns string, options metav1.ListOptions) (PodList, error) {
	podsApi, err := GetPodApi(ns)
	if err != nil {
		return nil, err
	}
	pods, err := podsApi.List(options)
	if err != nil {
		return nil, err
	}
	return pods.Items, nil
}

func GetPodByName(ns string, name string, options metav1.GetOptions) (*v1.Pod, error) {
	podsApi, err := GetPodApi(ns)
	if err != nil {
		return nil, err
	}
	return podsApi.Get(name, options)
}

func UpdatePod(ns string, pod *v1.Pod) (*v1.Pod, error) {
	podsApi, err := GetPodApi(ns)
	if err != nil {
		return nil, err
	}
	return podsApi.Update(pod)
}

func DeletePod(ns string, name string, options *metav1.DeleteOptions) error {
	podsApi, err := GetPodApi(ns)
	if err != nil {
		return err
	}
	return podsApi.Delete(name, options)
}

func CreatePod(ns string, pod *v1.Pod) (*v1.Pod, error) {
	podsApi, err := GetPodApi(ns)
	if err != nil {
		return nil, err
	}
	return podsApi.Create(pod)
}

func GetPodLogsRequest(ns string, name string, options *v1.PodLogOptions) (*rest.Request, error) {
	podsApi, err := GetPodApi(ns)
	if err != nil {
		return nil, err
	}
	return podsApi.GetLogs(name, options), nil
}

func WatchPod(ns string, options metav1.ListOptions) (watch.Interface, error) {
	watchApi, err := GetPodApi(ns)
	if err != nil {
		return nil, err
	}
	return watchApi.Watch(options)
}

func GetPodJson(ns string, name string, options metav1.GetOptions) (string, error) {
	pod, err := GetPodByName(ns, name, options)
	if err != nil {
		return "", err
	}
	podData, err := ObjectToJsonString(pod)
	if err != nil {
		return "", err
	}
	return string(podData), nil
}
