package project

import (
	"fmt"
	"github.com/kgysu/oc-wrapper/wrapper"
	"k8s.io/api/apps/v1beta1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type OpenshiftStatefulSet struct {
	name        string
	statefulSet v1beta1.StatefulSet
}

func fromStatefulSet(statefulSet v1beta1.StatefulSet) OpenshiftStatefulSet {
	return OpenshiftStatefulSet{
		name:        statefulSet.Name,
		statefulSet: statefulSet,
	}
}

func (oStatefulSet OpenshiftStatefulSet) setStatefulSet(statefulSet v1beta1.StatefulSet) {
	oStatefulSet.name = statefulSet.Name
	oStatefulSet.statefulSet = statefulSet
}

func (oStatefulSet OpenshiftStatefulSet) GetName() string {
	return oStatefulSet.name
}

func (oStatefulSet OpenshiftStatefulSet) GetKind() string {
	return StatefulSetKey
}

func (oStatefulSet OpenshiftStatefulSet) GetStatus() string {
	return fmt.Sprintf("%d (%d/%d)", oStatefulSet.statefulSet.Status.Replicas,
		oStatefulSet.statefulSet.Status.ReadyReplicas,
		oStatefulSet.statefulSet.Status.CurrentReplicas)
}

func (oStatefulSet OpenshiftStatefulSet) GetStatefulSet() v1beta1.StatefulSet {
	return oStatefulSet.statefulSet
}

func (oStatefulSet OpenshiftStatefulSet) Create(namespace string) error {
	_, err := wrapper.CreateStatefulSet(namespace, &oStatefulSet.statefulSet)
	if err != nil {
		return err
	}
	//oStatefulSet.setStatefulSet(createdStatefulSet)
	return nil
}

func (oStatefulSet OpenshiftStatefulSet) Update(namespace string) error {
	_, err := wrapper.UpdateStatefulSet(namespace, &oStatefulSet.statefulSet)
	if err != nil {
		return err
	}
	//oStatefulSet.setStatefulSet(updatedStatefulSet)
	return nil
}

func (oStatefulSet OpenshiftStatefulSet) Delete(namespace string, options v12.DeleteOptions) error {
	return wrapper.DeleteStatefulSet(namespace, oStatefulSet.name, options)
}
