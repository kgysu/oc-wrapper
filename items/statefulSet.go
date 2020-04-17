package items

import (
	"fmt"
	"github.com/kgysu/oc-wrapper/client"
	"github.com/kgysu/oc-wrapper/converter"
	"github.com/kgysu/oc-wrapper/fileutils"
	v1 "k8s.io/api/apps/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"strconv"
	"strings"
)

var OpStatefulSetTypeMeta = v12.TypeMeta{
	Kind:       "StatefulSet",
	APIVersion: "apps/v1",
}

type OpStatefulSet struct {
	StatefulSet *v1.StatefulSet
}

func NewOpStatefulSet(StatefulSet v1.StatefulSet) *OpStatefulSet {
	StatefulSet.TypeMeta = OpStatefulSetTypeMeta
	return &OpStatefulSet{
		StatefulSet: &StatefulSet,
	}
}

// Methods

func (oStatefulSet OpStatefulSet) GetFileName() string {
	return fmt.Sprintf("%s-%s.yaml", oStatefulSet.GetName(), oStatefulSet.GetKind())
}

func (oStatefulSet OpStatefulSet) WriteToFile(file string) error {
	yamlContent, err := oStatefulSet.ToYaml()
	if err != nil {
		return err
	}
	return fileutils.CreateFile(file, []byte(yamlContent))
}

func (oStatefulSet *OpStatefulSet) LoadFromFile(file string, envs map[string]string) error {
	tempData, err := fileutils.ReadFile(file)
	if err != nil {
		return err
	}
	data := fileutils.ReplaceEnvs(string(tempData), envs)
	err = oStatefulSet.FromData([]byte(data))
	if err != nil {
		return err
	}
	return nil
}

func (oStatefulSet *OpStatefulSet) Get(namespace string, restConf *rest.Config, name string) error {
	StatefulSetInterface, err := client.GetStatefulSetsInterface(namespace, restConf)
	if err != nil {
		return err
	}
	StatefulSet, err := StatefulSetInterface.Get(name, v12.GetOptions{})
	if err != nil {
		return err
	}
	oStatefulSet.StatefulSet = StatefulSet
	return nil
}

func (oStatefulSet OpStatefulSet) Create(namespace string, restConf *rest.Config) error {
	StatefulSetInterface, err := client.GetStatefulSetsInterface(namespace, restConf)
	if err != nil {
		return err
	}
	_, err = StatefulSetInterface.Create(oStatefulSet.StatefulSet)
	if err != nil {
		return err
	}
	return nil
}

func (oStatefulSet OpStatefulSet) Delete(namespace string, restConf *rest.Config, options *v12.DeleteOptions) error {
	StatefulSetInterface, err := client.GetStatefulSetsInterface(namespace, restConf)
	if err != nil {
		return err
	}
	err = StatefulSetInterface.Delete(oStatefulSet.StatefulSet.Name, options)
	if err != nil {
		return err
	}
	return nil
}

func (oStatefulSet OpStatefulSet) Update(namespace string, restConf *rest.Config) error {
	StatefulSetInterface, err := client.GetStatefulSetsInterface(namespace, restConf)
	if err != nil {
		return err
	}
	toUpdate, err := StatefulSetInterface.Get(oStatefulSet.GetName(), v12.GetOptions{})
	if err != nil {
		return err
	}
	toUpdate.Spec = oStatefulSet.StatefulSet.Spec
	toUpdate.Labels = oStatefulSet.StatefulSet.Labels
	toUpdate.Name = oStatefulSet.StatefulSet.Name
	_, err = StatefulSetInterface.Update(toUpdate)
	if err != nil {
		return err
	}
	return nil
}

func (oStatefulSet OpStatefulSet) UpdateScale(replicas int32, namespace string, restConf *rest.Config) error {
	StatefulSetInterface, err := client.GetStatefulSetsInterface(namespace, restConf)
	if err != nil {
		return err
	}
	toUpdate, err := StatefulSetInterface.Get(oStatefulSet.GetName(), v12.GetOptions{})
	if err != nil {
		return err
	}
	toUpdate.Spec.Replicas = &replicas
	_, err = StatefulSetInterface.Update(toUpdate)
	if err != nil {
		return err
	}
	return nil
}

func (oStatefulSet OpStatefulSet) GetScale() int32 {
	value := fmt.Sprintf("%d", oStatefulSet.StatefulSet.Spec.Replicas)
	val, _ := strconv.ParseInt(value, 10, 32)
	return int32(val)
}

func (oStatefulSet OpStatefulSet) IsScalable() bool {
	return true
}

func (oStatefulSet OpStatefulSet) String() string {
	return fmt.Sprintf("%s %s ", oStatefulSet.Info(), oStatefulSet.Status())
}

func (oStatefulSet OpStatefulSet) Info() string {
	return fmt.Sprintf("[%s] %s",
		oStatefulSet.GetKind(),
		oStatefulSet.GetName())
}

func (oStatefulSet OpStatefulSet) Status() string {
	return fmt.Sprintf("%d (%d/%d/%d)",
		oStatefulSet.StatefulSet.Spec.Replicas,
		oStatefulSet.StatefulSet.Status.ReadyReplicas,
		oStatefulSet.StatefulSet.Status.CurrentReplicas,
		oStatefulSet.StatefulSet.Status.UpdatedReplicas)
}

func (oStatefulSet OpStatefulSet) InfoStatusHtml() string {
	replicasStatus := "warning"
	value := fmt.Sprintf("%d", oStatefulSet.StatefulSet.Spec.Replicas)
	val, _ := strconv.ParseInt(value, 10, 32)
	if int32(val) == oStatefulSet.StatefulSet.Status.ReadyReplicas {
		replicasStatus = "success"
	}
	readyStatus := "warning"
	if oStatefulSet.StatefulSet.Status.CurrentReplicas == oStatefulSet.StatefulSet.Status.ReadyReplicas {
		readyStatus = "success"
	}
	return fmt.Sprint(
		createInfo(oStatefulSet.GetKind(), oStatefulSet.GetName()),
		createLabelBadges(oStatefulSet.StatefulSet.Labels),
		createStatusButton(replicasStatus, fmt.Sprint("Replicas ",
			createBadge("light", fmt.Sprintf("%d", oStatefulSet.StatefulSet.Spec.Replicas)))),
		createStatusButton(readyStatus, fmt.Sprint("Status ",
			createBadge("light", fmt.Sprintf("(%d/%d/%d)", oStatefulSet.StatefulSet.Status.ReadyReplicas,
				oStatefulSet.StatefulSet.Status.CurrentReplicas, oStatefulSet.StatefulSet.Status.UpdatedReplicas)))),
	)
}

func (oStatefulSet OpStatefulSet) GetName() string {
	return oStatefulSet.StatefulSet.Name
}

func (oStatefulSet OpStatefulSet) GetKind() string {
	return oStatefulSet.StatefulSet.Kind
}

func (oStatefulSet *OpStatefulSet) ToYaml() (string, error) {
	var contentBuilder strings.Builder
	err := converter.ObjToYaml(oStatefulSet.StatefulSet, &contentBuilder, true, true)
	if err != nil {
		return "", err
	}
	return contentBuilder.String(), nil
}

func (oStatefulSet *OpStatefulSet) FromData(data []byte) error {
	_, _, err := converter.YamlToObject(data, false, oStatefulSet.StatefulSet)
	if err != nil {
		return err
	}
	return nil
}
