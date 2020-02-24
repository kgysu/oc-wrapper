package project

import (
	"fmt"
	"github.com/kgysu/oc-wrapper/wrapper"
	v1 "k8s.io/api/apps/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type OpenshiftKubeStatefulSet struct {
	name            string
	kubeStatefulSet v1.StatefulSet
}

func fromKubeStatefulSet(kubeStatefulSet v1.StatefulSet) OpenshiftKubeStatefulSet {
	return OpenshiftKubeStatefulSet{
		name:            kubeStatefulSet.Name,
		kubeStatefulSet: kubeStatefulSet,
	}
}

func (oKubeStatefulSet OpenshiftKubeStatefulSet) setKubeStatefulSet(kubeStatefulSet v1.StatefulSet) {
	oKubeStatefulSet.name = kubeStatefulSet.Name
	oKubeStatefulSet.kubeStatefulSet = kubeStatefulSet
}

func (oKubeStatefulSet OpenshiftKubeStatefulSet) GetName() string {
	return oKubeStatefulSet.name
}

func (oKubeStatefulSet OpenshiftKubeStatefulSet) GetKind() string {
	return KubeStatefulSetKey
}

func (oKubeStatefulSet OpenshiftKubeStatefulSet) GetStatus() string {
	return fmt.Sprintf("%d (%d/%d)", oKubeStatefulSet.kubeStatefulSet.Status.Replicas,
		oKubeStatefulSet.kubeStatefulSet.Status.ReadyReplicas,
		oKubeStatefulSet.kubeStatefulSet.Status.CurrentReplicas)
}

func (oKubeStatefulSet OpenshiftKubeStatefulSet) GetKubeStatefulSet() v1.StatefulSet {
	return oKubeStatefulSet.kubeStatefulSet
}

func (oKubeStatefulSet OpenshiftKubeStatefulSet) Create(namespace string) error {
	_, err := wrapper.CreateKubeStatefulSet(namespace, &oKubeStatefulSet.kubeStatefulSet)
	if err != nil {
		return err
	}
	//oKubeStatefulSet.setKubeStatefulSet(createdKubeStatefulSet)
	return nil
}

func (oKubeStatefulSet OpenshiftKubeStatefulSet) Update(namespace string) error {
	_, err := wrapper.UpdateKubeStatefulSet(namespace, &oKubeStatefulSet.kubeStatefulSet)
	if err != nil {
		return err
	}
	//oKubeStatefulSet.setKubeStatefulSet(updatedKubeStatefulSet)
	return nil
}

func (oKubeStatefulSet OpenshiftKubeStatefulSet) Delete(namespace string, options v12.DeleteOptions) error {
	return wrapper.DeleteKubeStatefulSet(namespace, oKubeStatefulSet.name, options)
}
