package wrapper

import (
	appsv1client "github.com/openshift/client-go/apps/clientset/versioned/typed/apps/v1"
	authorizationv1client "github.com/openshift/client-go/authorization/clientset/versioned/typed/authorization/v1"
	buildv1client "github.com/openshift/client-go/build/clientset/versioned/typed/build/v1"
	imagev1client "github.com/openshift/client-go/image/clientset/versioned/typed/image/v1"
	networkv1client "github.com/openshift/client-go/network/clientset/versioned/typed/network/v1"
	oauthv1client "github.com/openshift/client-go/oauth/clientset/versioned/typed/oauth/v1"
	projectv1client "github.com/openshift/client-go/project/clientset/versioned/typed/project/v1"
	quotav1client "github.com/openshift/client-go/quota/clientset/versioned/typed/quota/v1"
	routev1client "github.com/openshift/client-go/route/clientset/versioned/typed/route/v1"
	securityv1client "github.com/openshift/client-go/security/clientset/versioned/typed/security/v1"
	templatev1client "github.com/openshift/client-go/template/clientset/versioned/typed/template/v1"
	userv1client "github.com/openshift/client-go/user/clientset/versioned/typed/user/v1"
	"k8s.io/client-go/kubernetes"
	admissionregistrationv1beta1client "k8s.io/client-go/kubernetes/typed/admissionregistration/v1beta1"
	kube_appsv1client "k8s.io/client-go/kubernetes/typed/apps/v1"
	appsv1beta1client "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
	appsv1beta2client "k8s.io/client-go/kubernetes/typed/apps/v1beta2"
	authenticationv1client "k8s.io/client-go/kubernetes/typed/authentication/v1"
	authenticationv1beta1client "k8s.io/client-go/kubernetes/typed/authentication/v1beta1"
	kube_authorizationv1client "k8s.io/client-go/kubernetes/typed/authorization/v1"
	authorizationv1beta1client "k8s.io/client-go/kubernetes/typed/authorization/v1beta1"
	autoscalingv1client "k8s.io/client-go/kubernetes/typed/autoscaling/v1"
	autoscalingv2beta1client "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta1"
	batchv1client "k8s.io/client-go/kubernetes/typed/batch/v1"
	batchv1beta1client "k8s.io/client-go/kubernetes/typed/batch/v1beta1"
	certificatesv1beta1client "k8s.io/client-go/kubernetes/typed/certificates/v1beta1"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	eventsv1beta1client "k8s.io/client-go/kubernetes/typed/events/v1beta1"
	extensionsv1beta1client "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
	networkingv1client "k8s.io/client-go/kubernetes/typed/networking/v1"
	policyv1beta1client "k8s.io/client-go/kubernetes/typed/policy/v1beta1"
	rbacv1client "k8s.io/client-go/kubernetes/typed/rbac/v1"
	rbacv1beta1client "k8s.io/client-go/kubernetes/typed/rbac/v1beta1"
	storagev1client "k8s.io/client-go/kubernetes/typed/storage/v1"
	storagev1beta1client "k8s.io/client-go/kubernetes/typed/storage/v1beta1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// APIs

func GetConfigs() (*rest.Config, string, error) {
	// Instantiate loader for kubeconfig file.
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{},
	)

	namespace, _, err := kubeconfig.Namespace()
	if err != nil {
		return nil, "", err
	}
	restConfig, err := kubeconfig.ClientConfig()
	if err != nil {
		return nil, "", err
	}
	return restConfig, namespace, nil
}

func GetKubeClientSet() (*kubernetes.Clientset, string, error) {
	restConfig, ns, err := GetConfigs()
	if err != nil {
		return nil, ns, err
	}
	clientset, err := kubernetes.NewForConfig(restConfig)
	return clientset, ns, nil
}

// CORE APIs

func GetCoreV1Client() (*corev1client.CoreV1Client, string, error) {
	restConfig, ns, err := GetConfigs()
	if err != nil {
		return nil, ns, err
	}
	client, err := corev1client.NewForConfig(restConfig)
	if err != nil {
		return nil, ns, err
	}
	return client, ns, nil
}

func GetNamespace(fromLocal string, fromUser string) string {
	if fromUser == "" {
		return fromLocal
	}
	return fromUser
}

func GetPodApi(namespace string) (corev1client.PodInterface, error) {
	client, ns, err := GetCoreV1Client()
	if err != nil {
		return nil, err
	}
	return client.Pods(GetNamespace(ns, namespace)), nil
}

func GetServiceApi(namespace string) (corev1client.ServiceInterface, error) {
	client, ns, err := GetCoreV1Client()
	if err != nil {
		return nil, err
	}
	return client.Services(GetNamespace(ns, namespace)), nil
}

func GetSecretApi(namespace string) (corev1client.SecretInterface, error) {
	client, ns, err := GetCoreV1Client()
	if err != nil {
		return nil, err
	}
	return client.Secrets(GetNamespace(ns, namespace)), nil
}

func GetConfigMapApi(namespace string) (corev1client.ConfigMapInterface, error) {
	client, ns, err := GetCoreV1Client()
	if err != nil {
		return nil, err
	}
	return client.ConfigMaps(GetNamespace(ns, namespace)), nil
}

func GetEventApi(namespace string) (corev1client.EventInterface, error) {
	client, ns, err := GetCoreV1Client()
	if err != nil {
		return nil, err
	}
	return client.Events(GetNamespace(ns, namespace)), nil
}

func GetPersistentVolumeClaimsApi(namespace string) (corev1client.PersistentVolumeClaimInterface, error) {
	client, ns, err := GetCoreV1Client()
	if err != nil {
		return nil, err
	}
	return client.PersistentVolumeClaims(GetNamespace(ns, namespace)), nil
}

func GetReplicationControllerApi(namespace string) (corev1client.ReplicationControllerInterface, error) {
	client, ns, err := GetCoreV1Client()
	if err != nil {
		return nil, err
	}
	return client.ReplicationControllers(GetNamespace(ns, namespace)), nil
}

func GetNamespaceApi(namespace string) (corev1client.NamespaceInterface, error) {
	client, _, err := GetCoreV1Client()
	if err != nil {
		return nil, err
	}
	return client.Namespaces(), nil
}

func GetPersistentVolumeApi(namespace string) (corev1client.PersistentVolumeInterface, error) {
	client, _, err := GetCoreV1Client()
	if err != nil {
		return nil, err
	}
	return client.PersistentVolumes(), nil
}

func GetServiceAccountApi(namespace string) (corev1client.ServiceAccountInterface, error) {
	client, ns, err := GetCoreV1Client()
	if err != nil {
		return nil, err
	}
	return client.ServiceAccounts(GetNamespace(ns, namespace)), nil
}

// AppsV1 APIs

func GetAppsV1Client() (*appsv1client.AppsV1Client, string, error) {
	restConfig, ns, err := GetConfigs()
	if err != nil {
		return nil, ns, err
	}
	client, err := appsv1client.NewForConfig(restConfig)
	if err != nil {
		return nil, ns, err
	}
	return client, ns, nil
}

func GetDeploymentConfigApi(namespace string) (appsv1client.DeploymentConfigInterface, error) {
	client, ns, err := GetAppsV1Client()
	if err != nil {
		return nil, err
	}
	return client.DeploymentConfigs(GetNamespace(ns, namespace)), nil
}

// Route APIs

func GetRouteV1Client() (*routev1client.RouteV1Client, string, error) {
	restConfig, ns, err := GetConfigs()
	if err != nil {
		return nil, ns, err
	}
	client, err := routev1client.NewForConfig(restConfig)
	if err != nil {
		return nil, ns, err
	}
	return client, ns, nil
}

func GetRouteApi(namespace string) (routev1client.RouteInterface, error) {
	client, ns, err := GetRouteV1Client()
	if err != nil {
		return nil, err
	}
	return client.Routes(GetNamespace(ns, namespace)), nil
}

// Project APIs

func GetProjectV1Client() (*projectv1client.ProjectV1Client, string, error) {
	restConfig, ns, err := GetConfigs()
	if err != nil {
		return nil, ns, err
	}
	client, err := projectv1client.NewForConfig(restConfig)
	if err != nil {
		return nil, ns, err
	}
	return client, ns, nil
}

func GetProjectApi(namespace string) (projectv1client.ProjectInterface, error) {
	client, _, err := GetProjectV1Client()
	if err != nil {
		return nil, err
	}
	return client.Projects(), nil
}

// AppsV1 Beta1 APIs

func GetAppsV1Beta1Client() (*appsv1beta1client.AppsV1beta1Client, string, error) {
	restConfig, ns, err := GetConfigs()
	if err != nil {
		return nil, ns, err
	}
	client, err := appsv1beta1client.NewForConfig(restConfig)
	if err != nil {
		return nil, ns, err
	}
	return client, ns, nil
}

func GetStatefulSetApi(namespace string) (appsv1beta1client.StatefulSetInterface, error) {
	client, ns, err := GetAppsV1Beta1Client()
	if err != nil {
		return nil, err
	}
	return client.StatefulSets(GetNamespace(ns, namespace)), nil
}

func GetKubeStatefulSetApi(namespace string) (kube_appsv1client.StatefulSetInterface, error) {
	client, ns, err := GetKubeAppsV1Client()
	if err != nil {
		return nil, err
	}
	return client.StatefulSets(GetNamespace(ns, namespace)), nil
}

func GetDeploymentsApi(namespace string) (appsv1beta1client.DeploymentInterface, error) {
	client, ns, err := GetAppsV1Beta1Client()
	if err != nil {
		return nil, err
	}
	return client.Deployments(GetNamespace(ns, namespace)), nil
}

// Role Based Authentication APIs
func GetRbacV1Client() (*rbacv1client.RbacV1Client, string, error) {
	restConfig, ns, err := GetConfigs()
	if err != nil {
		return nil, ns, err
	}
	client, err := rbacv1client.NewForConfig(restConfig)
	if err != nil {
		return nil, ns, err
	}
	return client, ns, nil
}

func GetRoleApi(namespace string) (rbacv1client.RoleInterface, error) {
	client, ns, err := GetRbacV1Client()
	if err != nil {
		return nil, err
	}
	return client.Roles(GetNamespace(ns, namespace)), nil
}

func GetRoleBindingApi(namespace string) (rbacv1client.RoleBindingInterface, error) {
	client, ns, err := GetRbacV1Client()
	if err != nil {
		return nil, err
	}
	return client.RoleBindings(GetNamespace(ns, namespace)), nil
}

// Kube Apis

func GetKubeAppsV1Client() (*kube_appsv1client.AppsV1Client, string, error) {
	restConfig, ns, err := GetConfigs()
	if err != nil {
		return nil, ns, err
	}
	client, err := kube_appsv1client.NewForConfig(restConfig)
	if err != nil {
		return nil, ns, err
	}
	return client, ns, nil
}

// Other
func GetStorageV1Client() (*storagev1client.StorageV1Client, string, error) {
	restConfig, ns, err := GetConfigs()
	if err != nil {
		return nil, ns, err
	}
	client, err := storagev1client.NewForConfig(restConfig)
	if err != nil {
		return nil, ns, err
	}
	return client, ns, nil
}

func GetAppsV1Beta2Client() (*appsv1beta2client.AppsV1beta2Client, string, error) {
	restConfig, ns, err := GetConfigs()
	if err != nil {
		return nil, ns, err
	}
	client, err := appsv1beta2client.NewForConfig(restConfig)
	if err != nil {
		return nil, ns, err
	}
	return client, ns, nil
}

func GetEventsV1Beta1Client() (*eventsv1beta1client.EventsV1beta1Client, string, error) {
	restConfig, ns, err := GetConfigs()
	if err != nil {
		return nil, ns, err
	}
	client, err := eventsv1beta1client.NewForConfig(restConfig)
	if err != nil {
		return nil, ns, err
	}
	return client, ns, nil
}

func GetExtensionsV1Beta1Client() (*extensionsv1beta1client.ExtensionsV1beta1Client, string, error) {
	restConfig, ns, err := GetConfigs()
	if err != nil {
		return nil, ns, err
	}
	client, err := extensionsv1beta1client.NewForConfig(restConfig)
	if err != nil {
		return nil, ns, err
	}
	return client, ns, nil
}

func GetStorageV1Beta1Client() (*storagev1beta1client.StorageV1beta1Client, string, error) {
	restConfig, ns, err := GetConfigs()
	if err != nil {
		return nil, ns, err
	}
	client, err := storagev1beta1client.NewForConfig(restConfig)
	if err != nil {
		return nil, ns, err
	}
	return client, ns, nil
}

// Other Openshift

func GetAuthorizationV1Client() (*authorizationv1client.AuthorizationV1Client, string, error) {
	restConfig, ns, err := GetConfigs()
	if err != nil {
		return nil, ns, err
	}
	client, err := authorizationv1client.NewForConfig(restConfig)
	if err != nil {
		return nil, ns, err
	}
	return client, ns, nil
}

func GetBuildV1Client() (*buildv1client.BuildV1Client, string, error) {
	restConfig, ns, err := GetConfigs()
	if err != nil {
		return nil, ns, err
	}
	client, err := buildv1client.NewForConfig(restConfig)
	if err != nil {
		return nil, ns, err
	}
	return client, ns, nil
}

func GetImageV1Client() (*imagev1client.ImageV1Client, string, error) {
	restConfig, ns, err := GetConfigs()
	if err != nil {
		return nil, ns, err
	}
	client, err := imagev1client.NewForConfig(restConfig)
	if err != nil {
		return nil, ns, err
	}
	return client, ns, nil
}

func GetNetworkV1Client() (*networkv1client.NetworkV1Client, string, error) {
	restConfig, ns, err := GetConfigs()
	if err != nil {
		return nil, ns, err
	}
	client, err := networkv1client.NewForConfig(restConfig)
	if err != nil {
		return nil, ns, err
	}
	return client, ns, nil
}

func GetOauthV1Client() (*oauthv1client.OauthV1Client, string, error) {
	restConfig, ns, err := GetConfigs()
	if err != nil {
		return nil, ns, err
	}
	client, err := oauthv1client.NewForConfig(restConfig)
	if err != nil {
		return nil, ns, err
	}
	return client, ns, nil
}

func GetQuotaV1Client() (*quotav1client.QuotaV1Client, string, error) {
	restConfig, ns, err := GetConfigs()
	if err != nil {
		return nil, ns, err
	}
	client, err := quotav1client.NewForConfig(restConfig)
	if err != nil {
		return nil, ns, err
	}
	return client, ns, nil
}

func GetSecurityV1Client() (*securityv1client.SecurityV1Client, string, error) {
	restConfig, ns, err := GetConfigs()
	if err != nil {
		return nil, ns, err
	}
	client, err := securityv1client.NewForConfig(restConfig)
	if err != nil {
		return nil, ns, err
	}
	return client, ns, nil
}

func GetTemplateV1Client() (*templatev1client.TemplateV1Client, string, error) {
	restConfig, ns, err := GetConfigs()
	if err != nil {
		return nil, ns, err
	}
	client, err := templatev1client.NewForConfig(restConfig)
	if err != nil {
		return nil, ns, err
	}
	return client, ns, nil
}

func GetUserV1Client() (*userv1client.UserV1Client, string, error) {
	restConfig, ns, err := GetConfigs()
	if err != nil {
		return nil, ns, err
	}
	client, err := userv1client.NewForConfig(restConfig)
	if err != nil {
		return nil, ns, err
	}
	return client, ns, nil
}

// Other Kubes

func GetAdmissionregistrationV1beta1Client() (*admissionregistrationv1beta1client.AdmissionregistrationV1beta1Client, string, error) {
	restConfig, ns, err := GetConfigs()
	if err != nil {
		return nil, ns, err
	}
	client, err := admissionregistrationv1beta1client.NewForConfig(restConfig)
	if err != nil {
		return nil, ns, err
	}
	return client, ns, nil
}

func GetAuthenticationV1beta1Client() (*authenticationv1beta1client.AuthenticationV1beta1Client, string, error) {
	restConfig, ns, err := GetConfigs()
	if err != nil {
		return nil, ns, err
	}
	client, err := authenticationv1beta1client.NewForConfig(restConfig)
	if err != nil {
		return nil, ns, err
	}
	return client, ns, nil
}

func GetAuthenticationV1Client() (*authenticationv1client.AuthenticationV1Client, string, error) {
	restConfig, ns, err := GetConfigs()
	if err != nil {
		return nil, ns, err
	}
	client, err := authenticationv1client.NewForConfig(restConfig)
	if err != nil {
		return nil, ns, err
	}
	return client, ns, nil
}

func GetAuthorizationV1beta1Client() (*authorizationv1beta1client.AuthorizationV1beta1Client, string, error) {
	restConfig, ns, err := GetConfigs()
	if err != nil {
		return nil, ns, err
	}
	client, err := authorizationv1beta1client.NewForConfig(restConfig)
	if err != nil {
		return nil, ns, err
	}
	return client, ns, nil
}

func GetKubeAuthorizationV1Client() (*kube_authorizationv1client.AuthorizationV1Client, string, error) {
	restConfig, ns, err := GetConfigs()
	if err != nil {
		return nil, ns, err
	}
	client, err := kube_authorizationv1client.NewForConfig(restConfig)
	if err != nil {
		return nil, ns, err
	}
	return client, ns, nil
}

func GetAutoscalingV1Client() (*autoscalingv1client.AutoscalingV1Client, string, error) {
	restConfig, ns, err := GetConfigs()
	if err != nil {
		return nil, ns, err
	}
	client, err := autoscalingv1client.NewForConfig(restConfig)
	if err != nil {
		return nil, ns, err
	}
	return client, ns, nil
}

func GetAutoscalingV2beta1Client() (*autoscalingv2beta1client.AutoscalingV2beta1Client, string, error) {
	restConfig, ns, err := GetConfigs()
	if err != nil {
		return nil, ns, err
	}
	client, err := autoscalingv2beta1client.NewForConfig(restConfig)
	if err != nil {
		return nil, ns, err
	}
	return client, ns, nil
}

func GetBatchV1beta1Client() (*batchv1beta1client.BatchV1beta1Client, string, error) {
	restConfig, ns, err := GetConfigs()
	if err != nil {
		return nil, ns, err
	}
	client, err := batchv1beta1client.NewForConfig(restConfig)
	if err != nil {
		return nil, ns, err
	}
	return client, ns, nil
}

func GetBatchV1Client() (*batchv1client.BatchV1Client, string, error) {
	restConfig, ns, err := GetConfigs()
	if err != nil {
		return nil, ns, err
	}
	client, err := batchv1client.NewForConfig(restConfig)
	if err != nil {
		return nil, ns, err
	}
	return client, ns, nil
}

func GetCertificatesV1beta1Client() (*certificatesv1beta1client.CertificatesV1beta1Client, string, error) {
	restConfig, ns, err := GetConfigs()
	if err != nil {
		return nil, ns, err
	}
	client, err := certificatesv1beta1client.NewForConfig(restConfig)
	if err != nil {
		return nil, ns, err
	}
	return client, ns, nil
}

func GetNetworkingV1Client() (*networkingv1client.NetworkingV1Client, string, error) {
	restConfig, ns, err := GetConfigs()
	if err != nil {
		return nil, ns, err
	}
	client, err := networkingv1client.NewForConfig(restConfig)
	if err != nil {
		return nil, ns, err
	}
	return client, ns, nil
}

func GetPolicyV1beta1Client() (*policyv1beta1client.PolicyV1beta1Client, string, error) {
	restConfig, ns, err := GetConfigs()
	if err != nil {
		return nil, ns, err
	}
	client, err := policyv1beta1client.NewForConfig(restConfig)
	if err != nil {
		return nil, ns, err
	}
	return client, ns, nil
}

func GetRbacV1beta1Client() (*rbacv1beta1client.RbacV1beta1Client, string, error) {
	restConfig, ns, err := GetConfigs()
	if err != nil {
		return nil, ns, err
	}
	client, err := rbacv1beta1client.NewForConfig(restConfig)
	if err != nil {
		return nil, ns, err
	}
	return client, ns, nil
}
