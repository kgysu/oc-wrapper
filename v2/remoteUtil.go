package v2

import (
	v1 "github.com/openshift/client-go/apps/clientset/versioned/typed/apps/v1"
	v15 "github.com/openshift/client-go/route/clientset/versioned/typed/route/v1"
	v13 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v14 "k8s.io/client-go/kubernetes/typed/apps/v1"
	v12 "k8s.io/client-go/kubernetes/typed/core/v1"
	v16 "k8s.io/client-go/kubernetes/typed/rbac/v1"
	"strings"
)

const Delimiter = ","
const AllItemsKey = ConfigMapKey + Delimiter +
	DeploymentConfigKey + Delimiter +
	EventKey + Delimiter +
	PodKey + Delimiter +
	PvKey + Delimiter +
	PvClaimKey + Delimiter +
	ReplicationControllerKey + Delimiter +
	RoleKey + Delimiter +
	RoleBindingKey + Delimiter +
	RouteKey + Delimiter +
	SecretKey + Delimiter +
	ServiceKey + Delimiter +
	ServiceAccountKey + Delimiter +
	StatefulSetKey

const ConfigMapKey = "ConfigMap"
const DeploymentConfigKey = "DeploymentConfig"
const DeploymentsKey = "Deployment"
const EventKey = "Event"
const KubeStatefulSetKey = "k-stateful-set"
const PodKey = "Pod"
const PvKey = "PersistentVolume" // notAllowed
const PvClaimKey = "PersistentVolumeClaim"
const ReplicationControllerKey = "ReplicationController"
const RoleKey = "Role"
const RoleBindingKey = "RoleBinding"
const RouteKey string = "Route"
const SecretKey = "Secret"
const ServiceKey = "Service"
const ServiceAccountKey = "ServiceAccount"
const StatefulSetKey = "StatefulSet"

const ProjectConfigObjectsKey = ConfigMapKey + Delimiter +
	//PvKey + Delimiter +
	PvClaimKey + Delimiter +
	RoleKey + Delimiter +
	RoleBindingKey + Delimiter +
	SecretKey + Delimiter +
	ServiceAccountKey

const ProjectObjectsKey = DeploymentConfigKey + Delimiter +
	StatefulSetKey + Delimiter +
	DeploymentsKey + Delimiter +
	ServiceKey + Delimiter +
	RouteKey

const AccountItemsKey = ServiceAccountKey + Delimiter +
	SecretKey + Delimiter +
	RoleKey + Delimiter +
	RoleBindingKey

const PvItemsKey = PvKey + Delimiter +
	PvClaimKey

const ProjectMonitorItemsKey = PodKey + Delimiter +
	ReplicationControllerKey + Delimiter +
	DeploymentConfigKey + Delimiter +
	StatefulSetKey

const WatchItemsKey = EventKey

const ScalableItemsKey = DeploymentConfigKey + Delimiter +
	StatefulSetKey

// List
func ListAllRemoteItems(namespace string, types string, options v13.ListOptions) ([]OpenshiftItem, error) {
	kubeClient, err := GetKubeAppsV1Client()
	if err != nil {
		return nil, err
	}
	appsClient, err := GetAppsV1Client()
	if err != nil {
		return nil, err
	}
	coreClient, err := GetCoreV1Client()
	if err != nil {
		return nil, err
	}
	routeClient, err := GetRouteV1Client()
	if err != nil {
		return nil, err
	}
	rbacClient, err := GetRbacV1Client()
	if err != nil {
		return nil, err
	}
	var resultItems []OpenshiftItem

	if strings.Contains(types, DeploymentConfigKey) {
		appendDeploymentConfigs(&resultItems, appsClient, namespace, options)
	}
	if strings.Contains(types, DeploymentsKey) {
		appendDeployments(&resultItems, kubeClient, namespace, options)
	}
	if strings.Contains(types, StatefulSetKey) {
		appendStatefulSets(&resultItems, kubeClient, namespace, options)
	}
	if strings.Contains(types, EventKey) {
		appendEvents(&resultItems, coreClient, namespace, options)
	}
	if strings.Contains(types, PodKey) {
		appendPods(&resultItems, coreClient, namespace, options)
	}
	if strings.Contains(types, ReplicationControllerKey) {
		appendReplicationControllers(&resultItems, coreClient, namespace, options)
	}
	if strings.Contains(types, ServiceKey) {
		appendServices(&resultItems, coreClient, namespace, options)
	}
	if strings.Contains(types, ConfigMapKey) {
		appendConfigMaps(&resultItems, coreClient, namespace, options)
	}
	if strings.Contains(types, SecretKey) {
		appendSecrets(&resultItems, coreClient, namespace, options)
	}
	if strings.Contains(types, ServiceAccountKey) {
		appendServiceAccounts(&resultItems, coreClient, namespace, options)
	}
	if strings.Contains(types, PvClaimKey) {
		appendPersistentVolumeClaims(&resultItems, coreClient, namespace, options)
	}
	if strings.Contains(types, PvKey) {
		appendPersistentVolumes(&resultItems, coreClient, options)
	}
	if strings.Contains(types, RouteKey) {
		appendRoutes(&resultItems, routeClient, namespace, options)
	}
	if strings.Contains(types, RoleKey) {
		appendRoles(&resultItems, rbacClient, namespace, options)
	}
	if strings.Contains(types, RoleBindingKey) {
		appendRoleBindings(&resultItems, rbacClient, namespace, options)
	}
	return resultItems, nil
}

// CoreV1
func appendPods(resultItems *[]OpenshiftItem, appsClient *v12.CoreV1Client, namespace string,
	options v13.ListOptions) {
	list, err := appsClient.Pods(namespace).List(options)
	onlyLogOnError(err)
	for _, item := range list.Items {
		*resultItems = append(*resultItems, NewOpenshiftItem(item.Name, PodKey, item.String()))
	}
}

func appendServices(resultItems *[]OpenshiftItem, appsClient *v12.CoreV1Client, namespace string,
	options v13.ListOptions) {
	list, err := appsClient.Services(namespace).List(options)
	onlyLogOnError(err)
	for _, item := range list.Items {
		*resultItems = append(*resultItems, NewOpenshiftItem(item.Name, ServiceKey, item.String()))
	}
}

func appendConfigMaps(resultItems *[]OpenshiftItem, appsClient *v12.CoreV1Client, namespace string,
	options v13.ListOptions) {
	list, err := appsClient.ConfigMaps(namespace).List(options)
	onlyLogOnError(err)
	for _, item := range list.Items {
		*resultItems = append(*resultItems, NewOpenshiftItem(item.Name, ConfigMapKey, item.String()))
	}
}

func appendEvents(resultItems *[]OpenshiftItem, appsClient *v12.CoreV1Client, namespace string,
	options v13.ListOptions) {
	list, err := appsClient.Events(namespace).List(options)
	onlyLogOnError(err)
	for _, item := range list.Items {
		*resultItems = append(*resultItems, NewOpenshiftItem(item.Name, EventKey, item.String()))
	}
}

func appendPersistentVolumeClaims(resultItems *[]OpenshiftItem, appsClient *v12.CoreV1Client, namespace string,
	options v13.ListOptions) {
	list, err := appsClient.PersistentVolumeClaims(namespace).List(options)
	onlyLogOnError(err)
	for _, item := range list.Items {
		*resultItems = append(*resultItems, NewOpenshiftItem(item.Name, PvClaimKey, item.String()))
	}
}

func appendPersistentVolumes(resultItems *[]OpenshiftItem, appsClient *v12.CoreV1Client, options v13.ListOptions) {
	list, err := appsClient.PersistentVolumes().List(options)
	onlyLogOnError(err)
	for _, item := range list.Items {
		*resultItems = append(*resultItems, NewOpenshiftItem(item.Name, PvKey, item.String()))
	}
}

func appendReplicationControllers(resultItems *[]OpenshiftItem, appsClient *v12.CoreV1Client, namespace string,
	options v13.ListOptions) {
	list, err := appsClient.ReplicationControllers(namespace).List(options)
	onlyLogOnError(err)
	for _, item := range list.Items {
		*resultItems = append(*resultItems, NewOpenshiftItem(item.Name, ReplicationControllerKey, item.String()))
	}
}

func appendSecrets(resultItems *[]OpenshiftItem, appsClient *v12.CoreV1Client, namespace string,
	options v13.ListOptions) {
	list, err := appsClient.Secrets(namespace).List(options)
	onlyLogOnError(err)
	for _, item := range list.Items {
		*resultItems = append(*resultItems, NewOpenshiftItem(item.Name, SecretKey, item.String()))
	}
}

func appendServiceAccounts(resultItems *[]OpenshiftItem, appsClient *v12.CoreV1Client, namespace string,
	options v13.ListOptions) {
	list, err := appsClient.ServiceAccounts(namespace).List(options)
	onlyLogOnError(err)
	for _, item := range list.Items {
		*resultItems = append(*resultItems, NewOpenshiftItem(item.Name, ServiceAccountKey, item.String()))
	}
}

// RbacV1
func appendRoles(resultItems *[]OpenshiftItem, appsClient *v16.RbacV1Client, namespace string,
	options v13.ListOptions) {
	list, err := appsClient.Roles(namespace).List(options)
	onlyLogOnError(err)
	for _, item := range list.Items {
		*resultItems = append(*resultItems, NewOpenshiftItem(item.Name, RoleKey, item.String()))
	}
}

func appendRoleBindings(resultItems *[]OpenshiftItem, appsClient *v16.RbacV1Client, namespace string,
	options v13.ListOptions) {
	list, err := appsClient.RoleBindings(namespace).List(options)
	onlyLogOnError(err)
	for _, item := range list.Items {
		*resultItems = append(*resultItems, NewOpenshiftItem(item.Name, RoleBindingKey, item.String()))
	}
}

// RouteV1
func appendRoutes(resultItems *[]OpenshiftItem, appsClient *v15.RouteV1Client, namespace string,
	options v13.ListOptions) {
	list, err := appsClient.Routes(namespace).List(options)
	onlyLogOnError(err)
	for _, item := range list.Items {
		*resultItems = append(*resultItems, NewOpenshiftItem(item.Name, RouteKey, item.String()))
	}
}

// AppsV1
func appendDeploymentConfigs(resultItems *[]OpenshiftItem, appsClient *v1.AppsV1Client, namespace string,
	options v13.ListOptions) {
	list, err := appsClient.DeploymentConfigs(namespace).List(options)
	onlyLogOnError(err)
	for _, item := range list.Items {
		*resultItems = append(*resultItems, NewOpenshiftItem(item.Name, DeploymentConfigKey, item.String()))
	}
}

func appendDeployments(resultItems *[]OpenshiftItem, appsClient v14.AppsV1Interface, namespace string,
	options v13.ListOptions) {
	list, err := appsClient.Deployments(namespace).List(options)
	onlyLogOnError(err)
	for _, item := range list.Items {
		*resultItems = append(*resultItems, NewOpenshiftItem(item.Name, DeploymentsKey, item.String()))
	}
}

func appendStatefulSets(resultItems *[]OpenshiftItem, appsClient v14.AppsV1Interface, namespace string,
	options v13.ListOptions) {
	list, err := appsClient.StatefulSets(namespace).List(options)
	onlyLogOnError(err)
	for _, item := range list.Items {
		*resultItems = append(*resultItems, NewOpenshiftItem(item.Name, StatefulSetKey, item.String()))
	}
}
