package v2

import (
	"fmt"
	v1 "github.com/openshift/client-go/apps/clientset/versioned/typed/apps/v1"
	v13 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v14 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"strings"
)

const Delimiter = ","
const AllItemsKey = "all"

const ConfigMapKey = "config-map"
const DeploymentConfigKey = "deployment-config"
const DeploymentsKey = "Deployment"
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

func ListAllRemoteItems(namespace string, types string, options v13.ListOptions) ([]OpenshiftItem, error) {
	kubeClient, err := GetKubeAppsV1Client()
	if err != nil {
		return nil, err
	}
	appsClient, err := GetAppsV1Client()
	if err != nil {
		return nil, err
	}
	var resultItems []OpenshiftItem

	if strings.Contains(types, DeploymentConfigKey) {
		appendDeploymentConfigs(resultItems, appsClient, namespace, options)
	}
	if strings.Contains(types, DeploymentsKey) {
		appendDeployments(resultItems, kubeClient, namespace, options)
	}
	if strings.Contains(types, StatefulSetKey) {
		appendStatefulSets(resultItems, kubeClient, namespace, options)
	}

	return resultItems, nil
}

func appendDeploymentConfigs(resultItems []OpenshiftItem, appsClient *v1.AppsV1Client, namespace string,
	options v13.ListOptions) {
	list, err := appsClient.DeploymentConfigs(namespace).List(options)
	onlyLogOnError(err)
	for _, item := range list.Items {
		resultItems = append(resultItems, New(item.Name, item.Kind, item.String()))
	}
}

func appendDeployments(resultItems []OpenshiftItem, appsClient v14.AppsV1Interface, namespace string,
	options v13.ListOptions) {
	list, err := appsClient.Deployments(namespace).List(options)
	onlyLogOnError(err)
	for _, item := range list.Items {
		resultItems = append(resultItems, New(item.Name, item.Kind, item.String()))
	}
}

func appendStatefulSets(resultItems []OpenshiftItem, appsClient v14.AppsV1Interface, namespace string,
	options v13.ListOptions) {
	list, err := appsClient.StatefulSets(namespace).List(options)
	onlyLogOnError(err)
	for _, item := range list.Items {
		resultItems = append(resultItems, New(item.Name, item.Kind, item.String()))
	}
}

func onlyLogOnError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
