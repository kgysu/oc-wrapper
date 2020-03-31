package v3

import (
	v1 "github.com/openshift/client-go/apps/clientset/versioned/typed/apps/v1"
	v13 "github.com/openshift/client-go/route/clientset/versioned/typed/route/v1"
	v12 "k8s.io/client-go/kubernetes/typed/core/v1"
	v14 "k8s.io/client-go/kubernetes/typed/rbac/v1"
	"k8s.io/client-go/rest"
)

func GetDeploymentConfigsInterface(namespace string, restConf *rest.Config) (v1.DeploymentConfigInterface, error) {
	appsClient, err := GetAppsV1Client(restConf)
	if err != nil {
		return nil, err
	}
	return appsClient.DeploymentConfigs(namespace), nil
}

func GetServicesInterface(namespace string, restConf *rest.Config) (v12.ServiceInterface, error) {
	coreClient, err := GetCoreV1Client(restConf)
	if err != nil {
		return nil, err
	}
	return coreClient.Services(namespace), nil
}

func GetConfigMapsInterface(namespace string, restConf *rest.Config) (v12.ConfigMapInterface, error) {
	coreClient, err := GetCoreV1Client(restConf)
	if err != nil {
		return nil, err
	}
	return coreClient.ConfigMaps(namespace), nil
}

func GetRoutesInterface(namespace string, restConf *rest.Config) (v13.RouteInterface, error) {
	routeClient, err := GetRouteV1Client(restConf)
	if err != nil {
		return nil, err
	}
	return routeClient.Routes(namespace), nil
}

func GetRolesInterface(namespace string, restConf *rest.Config) (v14.RoleInterface, error) {
	rbacClient, err := GetRbacV1Client(restConf)
	if err != nil {
		return nil, err
	}
	return rbacClient.Roles(namespace), nil
}
func GetRoleBindingsInterface(namespace string, restConf *rest.Config) (v14.RoleBindingInterface, error) {
	rbacClient, err := GetRbacV1Client(restConf)
	if err != nil {
		return nil, err
	}
	return rbacClient.RoleBindings(namespace), nil
}

func GetEventsInterface(namespace string, restConf *rest.Config) (v12.EventInterface, error) {
	coreClient, err := GetCoreV1Client(restConf)
	if err != nil {
		return nil, err
	}
	return coreClient.Events(namespace), nil
}
func GetReplicationControllersInterface(namespace string, restConf *rest.Config) (v12.ReplicationControllerInterface, error) {
	coreClient, err := GetCoreV1Client(restConf)
	if err != nil {
		return nil, err
	}
	return coreClient.ReplicationControllers(namespace), nil
}
func GetEndpointsInterface(namespace string, restConf *rest.Config) (v12.EndpointsInterface, error) {
	coreClient, err := GetCoreV1Client(restConf)
	if err != nil {
		return nil, err
	}
	return coreClient.Endpoints(namespace), nil
}
func GetSecretsInterface(namespace string, restConf *rest.Config) (v12.SecretInterface, error) {
	coreClient, err := GetCoreV1Client(restConf)
	if err != nil {
		return nil, err
	}
	return coreClient.Secrets(namespace), nil
}
func GetPersistentVolumeClaimsInterface(namespace string, restConf *rest.Config) (v12.PersistentVolumeClaimInterface, error) {
	coreClient, err := GetCoreV1Client(restConf)
	if err != nil {
		return nil, err
	}
	return coreClient.PersistentVolumeClaims(namespace), nil
}
func GetPodsInterface(namespace string, restConf *rest.Config) (v12.PodInterface, error) {
	coreClient, err := GetCoreV1Client(restConf)
	if err != nil {
		return nil, err
	}
	return coreClient.Pods(namespace), nil
}
func GetServiceAccountsInterface(namespace string, restConf *rest.Config) (v12.ServiceAccountInterface, error) {
	coreClient, err := GetCoreV1Client(restConf)
	if err != nil {
		return nil, err
	}
	return coreClient.ServiceAccounts(namespace), nil
}

func GetResourceQuotasInterface(namespace string, restConf *rest.Config) (v12.ResourceQuotaInterface, error) {
	coreClient, err := GetCoreV1Client(restConf)
	if err != nil {
		return nil, err
	}
	return coreClient.ResourceQuotas(namespace), nil
}
func GetLimitRangesInterface(namespace string, restConf *rest.Config) (v12.LimitRangeInterface, error) {
	coreClient, err := GetCoreV1Client(restConf)
	if err != nil {
		return nil, err
	}
	return coreClient.LimitRanges(namespace), nil
}
