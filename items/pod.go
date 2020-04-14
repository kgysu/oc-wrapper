package items

import (
	"fmt"
	"github.com/kgysu/oc-wrapper/client"
	"github.com/kgysu/oc-wrapper/converter"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"strings"
)

var OpPodTypeMeta = v12.TypeMeta{
	Kind:       "Pod",
	APIVersion: "v1",
}

type OpPod struct {
	Pod *v1.Pod
}

func NewOpPod(Pod v1.Pod) *OpPod {
	if Pod.TypeMeta.Kind != OpPodTypeMeta.Kind {
		Pod.TypeMeta = OpPodTypeMeta
	}
	return &OpPod{
		Pod: &Pod,
	}
}

// Methods

func (oPod *OpPod) GetFileName() string {
	return fmt.Sprintf("%s-%s.yaml", oPod.GetName(), oPod.GetKind())
}

func (oPod *OpPod) WriteToFile(file string) error {
	return nil
}

func (oPod *OpPod) LoadFromFile(file string, envs map[string]string) error {
	return nil
}

func (oPod *OpPod) Get(namespace string, restConf *rest.Config, name string) error {
	PodInterface, err := client.GetPodsInterface(namespace, restConf)
	if err != nil {
		return err
	}
	Pod, err := PodInterface.Get(name, v12.GetOptions{})
	if err != nil {
		return err
	}
	oPod.Pod = Pod
	return nil
}

func (oPod *OpPod) Create(namespace string, restConf *rest.Config) error {
	PodInterface, err := client.GetPodsInterface(namespace, restConf)
	if err != nil {
		return err
	}
	_, err = PodInterface.Create(oPod.Pod)
	if err != nil {
		return err
	}
	return nil
}

func (oPod *OpPod) Delete(namespace string, restConf *rest.Config, options *v12.DeleteOptions) error {
	PodInterface, err := client.GetPodsInterface(namespace, restConf)
	if err != nil {
		return err
	}
	err = PodInterface.Delete(oPod.Pod.Name, options)
	if err != nil {
		return err
	}
	return nil
}

func (oPod *OpPod) UpdateScale(replicas int32, namespace string, restConf *rest.Config) error {
	return fmt.Errorf("")
}

func (oPod *OpPod) GetScale() int32 {
	return 0
}

func (oPod *OpPod) String() string {
	return fmt.Sprintf("%s %s ", oPod.Info(), oPod.Status())
}

func (oPod *OpPod) Info() string {
	return fmt.Sprintf("[%s] %s ",
		oPod.GetKind(),
		oPod.GetName())
}

// TODO more infos
func (oPod *OpPod) Status() string {
	return fmt.Sprintf("%s [%s][%s][%s] (%s)",
		oPod.Pod.Status.Phase,
		oPod.Pod.Status.StartTime,
		oPod.Pod.Status.HostIP,
		oPod.Pod.Status.PodIP,
		oPod.Pod.Status.Message)
}

func (oPod OpPod) InfoStatusHtml() string {
	phaseStatus := "danger"
	if v1.PodRunning == oPod.Pod.Status.Phase {
		phaseStatus = "success"
	}
	if v1.PodSucceeded == oPod.Pod.Status.Phase {
		phaseStatus = "success"
	}
	if v1.PodPending == oPod.Pod.Status.Phase {
		phaseStatus = "warning"
	}
	if v1.PodFailed == oPod.Pod.Status.Phase {
		phaseStatus = "danger"
	}
	if v1.PodUnknown == oPod.Pod.Status.Phase {
		phaseStatus = "secondary"
	}
	return fmt.Sprint(
		createInfo(oPod.GetKind(), oPod.GetName()),
		createLabelBadges(oPod.Pod.Labels),
		createStatusButton(phaseStatus, fmt.Sprintf("%s", oPod.Pod.Status.Phase)),
		createStatusButton("secondary", fmt.Sprint("StartTime ",
			createBadge("light", fmt.Sprintf("%s", oPod.Pod.Status.StartTime)))),
		createStatusButton("secondary", fmt.Sprint("HostIP ",
			createBadge("light", oPod.Pod.Status.HostIP))),
		createStatusButton("secondary", fmt.Sprint("PodIP ",
			createBadge("light", oPod.Pod.Status.PodIP))),
		createStatusButton("secondary", fmt.Sprint("Message ",
			createBadge("light", oPod.Pod.Status.Message))),
	)
}

func (oPod *OpPod) GetName() string {
	return oPod.Pod.Name
}

func (oPod *OpPod) GetKind() string {
	return oPod.Pod.Kind
}

func (oPod *OpPod) ToYaml() (string, error) {
	var contentBuilder strings.Builder
	err := converter.ObjToYaml(oPod.Pod, &contentBuilder, true, true)
	if err != nil {
		return "", err
	}
	return contentBuilder.String(), nil
}

func (oPod *OpPod) FromData(data []byte) error {
	_, _, err := converter.YamlToObject(data, false, oPod.Pod)
	if err != nil {
		return err
	}
	return nil
}
