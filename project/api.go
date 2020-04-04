package project

// TODO try this
//type OpenshiftInterface interface {
//	Deploy(items []OpenshiftItemInterface)
//}
//
//type OpDeploymentConfigApi struct {
//	api v1.DeploymentConfigInterface
//}
//
//func NewOpenshiftInterfaceByKind(kind string, namespace string, restConf *rest.Config) (*OpenshiftInterface, error) {
//	if kind == "DeploymentConfig" {
//		dcI, err := v3.GetDeploymentConfigsInterface(namespace, restConf)
//		if err != nil {
//			return nil, err
//		}
//		return NewOpenshiftApi(dcI), nil
//	}
//}
//
//func NewOpenshiftApi(dcApi v1.DeploymentConfigInterface) *OpDeploymentConfigApi {
//	return &OpDeploymentConfigApi{api:dcApi}
//}
//
//func (api OpDeploymentConfigApi) Deploy(items []OpenshiftItemInterface){
//
//	api.api.Create(items[0].Create())
//}
