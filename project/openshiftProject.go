package project

import (
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const Delimiter = ","
const AllItemsKey = "all"

const ConfigMapKey = "config-map"
const DeploymentConfigKey = "deployment-config"
const EventKey = "event"
const KubeStatefulSetKey = "k-stateful-set"
const PodKey = "pod"
const PvKey = "pv"
const PvClaimKey = "pv-claim"
const ReplicationControllerKey = "replication-controller"
const RoleKey = "role"
const RoleBindingKey = "role-binding"
const RouteKey string = "route"
const SecretKey = "secret"
const ServiceKey = "service"
const ServiceAccountKey = "service-account"
const StatefulSetKey = "stateful-set"

const AccountItemsKey = ServiceAccountKey + Delimiter +
	SecretKey

const StaticItemsKey = ConfigMapKey + Delimiter +
	AccountItemsKey + Delimiter +
	RoleKey + Delimiter +
	RoleBindingKey + Delimiter +
	PvKey + Delimiter +
	PvClaimKey

const ProjectItemsKey = DeploymentConfigKey + Delimiter +
	StaticItemsKey + Delimiter +
	ReplicationControllerKey + Delimiter +
	RouteKey + Delimiter +
	ServiceKey + Delimiter +
	StatefulSetKey + Delimiter +
	KubeStatefulSetKey + Delimiter

const MonitorItemsKey = PodKey + Delimiter +
	ReplicationControllerKey + Delimiter +
	DeploymentConfigKey + Delimiter +
	StatefulSetKey

const WatchItemsKey = EventKey

const ScalableItemsKey = DeploymentConfigKey + Delimiter +
	StatefulSetKey

type OpenshiftProject struct {
	namespace string
	name      string
	items     []OpenshiftItem
}

func NewFromLocal(namespace string, folderPath string) (OpenshiftProject, error) {
	localProject := OpenshiftProject{
		namespace: namespace,
	}
	parsedProject, err := parseProjectFile(folderPath)
	if err != nil {
		return localProject, err
	}
	localProject.name = parsedProject.Name
	items, err := parseLocalItemFiles(localProject.namespace, folderPath) // ./items
	if err != nil {
		return localProject, err
	}
	localProject.items = items
	return localProject, nil
}

func NewFromRemote(namespace string, name string, options v1.ListOptions) (OpenshiftProject, error) {
	remoteProject := OpenshiftProject{
		namespace: namespace,
		name:      name,
	}
	items, errs := loadAllItemsFromServer(remoteProject.namespace, options)
	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Printf("error in loading from server: %s \n", err.Error())
		}
	}
	remoteProject.items = items
	return remoteProject, nil
}

func NewFromRemoteByTypes(namespace string, name string, itemTypes string, options v1.ListOptions) (OpenshiftProject, error) {
	remoteProject := OpenshiftProject{
		namespace: namespace,
		name:      name,
	}
	items, errs := loadItemsByTypeFromServer(remoteProject.namespace, itemTypes, options)
	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Printf("error in loading from server: %s \n", err.Error())
		}
	}
	remoteProject.items = items
	return remoteProject, nil
}

func (op OpenshiftProject) GetNamespace() string {
	return op.namespace
}

func (op OpenshiftProject) GetName() string {
	return op.name
}

func (op OpenshiftProject) GetItems() []OpenshiftItem {
	return op.items
}

func (op OpenshiftProject) appendItem(item OpenshiftItem) {
	op.items = append(op.items, item)
}
