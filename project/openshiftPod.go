package project

import (
	"fmt"
	"github.com/kgysu/oc-wrapper/wrapper"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type OpenshiftPod struct {
	name string
	pod  v1.Pod
}

func fromPod(pod v1.Pod) OpenshiftPod {
	return OpenshiftPod{
		name: pod.Name,
		pod:  pod,
	}
}

func (oPod OpenshiftPod) setPod(pod v1.Pod) {
	oPod.name = pod.Name
	oPod.pod = pod
}

func (oPod OpenshiftPod) GetName() string {
	return oPod.name
}

func (oPod OpenshiftPod) GetKind() string {
	return PodKey
}

func (oPod OpenshiftPod) GetStatus() string {
	return fmt.Sprintf("%s (%s) %s [%s]", oPod.pod.Status.Phase,
		oPod.pod.Status.StartTime.String(),
		oPod.pod.Status.Reason,
		oPod.pod.Status.Message)
}

func (oPod OpenshiftPod) GetPod() v1.Pod {
	return oPod.pod
}

func (oPod OpenshiftPod) Create(namespace string) error {
	_, err := wrapper.CreatePod(namespace, &oPod.pod)
	if err != nil {
		return err
	}
	//oPod.setPod(createdPod)
	return nil
}

func (oPod OpenshiftPod) Update(namespace string) error {
	_, err := wrapper.UpdatePod(namespace, &oPod.pod)
	if err != nil {
		return err
	}
	//oPod.setPod(updatedPod)
	return nil
}

func (oPod OpenshiftPod) Delete(namespace string, options v12.DeleteOptions) error {
	return wrapper.DeletePod(namespace, oPod.name, &options)
}
