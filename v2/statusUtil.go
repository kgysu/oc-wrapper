package v2

import (
	"fmt"
	v1 "github.com/openshift/api/apps/v1"
	v13 "k8s.io/api/apps/v1"
	v12 "k8s.io/api/core/v1"
)

func GetStatusByType(item OpenshiftItem) string {
	if item.kind == DeploymentConfigKey {
		return getStatusByDeploymentConfig(item)
	}
	if item.kind == ReplicationControllerKey {
		return getStatusByReplicationController(item)
	}
	if item.kind == PodKey {
		return getStatusByPod(item)
	}
	if item.kind == StatefulSetKey {
		return getStatusByStatefulSet(item)
	}
	if item.kind == EventKey {
		return getStatusByEvent(item)
	}
	return ""
}

func getStatusByDeploymentConfig(item OpenshiftItem) string {
	var realItem v1.DeploymentConfig
	err := item.ParseTo(&realItem)
	if err != nil {
		onlyLogOnError(err)
		return "?"
	}
	return fmt.Sprintf("%d (%d/%d)",
		realItem.Spec.Replicas,
		realItem.Status.ReadyReplicas,
		realItem.Status.AvailableReplicas)
}

func getStatusByStatefulSet(item OpenshiftItem) string {
	var realItem v13.StatefulSet
	err := item.ParseTo(&realItem)
	if err != nil {
		onlyLogOnError(err)
		return "?"
	}
	return fmt.Sprintf("%d (%d/%d)",
		realItem.Spec.Replicas,
		realItem.Status.ReadyReplicas,
		realItem.Status.CurrentReplicas)
}

func getStatusByPod(item OpenshiftItem) string {
	var realItem v12.Pod
	err := item.ParseTo(&realItem)
	if err != nil {
		onlyLogOnError(err)
		return "?"
	}
	return fmt.Sprintf("[%s] %s (%s) %s",
		realItem.Status.StartTime,
		realItem.Status.Phase,
		realItem.Status.Reason,
		realItem.Status.Message)
}

func getStatusByReplicationController(item OpenshiftItem) string {
	var realItem v12.ReplicationController
	err := item.ParseTo(&realItem)
	if err != nil {
		onlyLogOnError(err)
		return "?"
	}
	return fmt.Sprintf("%d (%d/%d)",
		realItem.Status.Replicas,
		realItem.Status.ReadyReplicas,
		realItem.Status.AvailableReplicas)
}

func getStatusByEvent(item OpenshiftItem) string {
	var realItem v12.Event
	err := item.ParseTo(&realItem)
	if err != nil {
		onlyLogOnError(err)
		return "?"
	}
	return fmt.Sprintf("[%s][%s] %s (%s) %s",
		realItem.Kind,
		realItem.EventTime,
		realItem.Name,
		realItem.Reason,
		realItem.Message)
}
