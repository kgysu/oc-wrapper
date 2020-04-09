package items

import (
	"fmt"
	"github.com/kgysu/oc-wrapper/converter"
	"github.com/kgysu/oc-wrapper/files"
	v3 "github.com/kgysu/oc-wrapper/v3"
	v1 "github.com/openshift/api/route/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"strings"
)

var OpRouteTypeMeta = v12.TypeMeta{
	Kind:       "Route",
	APIVersion: "route.openshift.io/v1",
}

type OpRoute struct {
	Route *v1.Route
}

func NewOpRoute(Route v1.Route) *OpRoute {
	Route.TypeMeta = OpRouteTypeMeta
	return &OpRoute{
		Route: &Route,
	}
}

// Methods

func (oRoute *OpRoute) GetFileName() string {
	return fmt.Sprintf("%s-%s.yaml", oRoute.GetName(), oRoute.GetKind())
}

func (oRoute *OpRoute) WriteToFile(file string) error {
	yamlContent, err := oRoute.ToYaml()
	if err != nil {
		return err
	}
	return files.CreateFile(file, []byte(yamlContent))
}

func (oRoute *OpRoute) LoadFromFile(file string, envs map[string]string) error {
	tempData, err := files.ReadFile(file)
	if err != nil {
		return err
	}
	data := files.ReplaceEnvs(string(tempData), envs)
	files.CheckContent(data, file)
	err = oRoute.FromData([]byte(data))
	if err != nil {
		return err
	}
	return nil
}

func (oRoute *OpRoute) Get(namespace string, restConf *rest.Config, name string) error {
	RouteInterface, err := v3.GetRoutesInterface(namespace, restConf)
	if err != nil {
		return err
	}
	Route, err := RouteInterface.Get(name, v12.GetOptions{})
	if err != nil {
		return err
	}
	oRoute.Route = Route
	return nil
}

func (oRoute *OpRoute) Create(namespace string, restConf *rest.Config) error {
	RouteInterface, err := v3.GetRoutesInterface(namespace, restConf)
	if err != nil {
		return err
	}
	_, err = RouteInterface.Create(oRoute.Route)
	if err != nil {
		return err
	}
	return nil
}

func (oRoute *OpRoute) Delete(namespace string, restConf *rest.Config, options *v12.DeleteOptions) error {
	RouteInterface, err := v3.GetRoutesInterface(namespace, restConf)
	if err != nil {
		return err
	}
	err = RouteInterface.Delete(oRoute.Route.Name, options)
	if err != nil {
		return err
	}
	return nil
}

func (oRoute *OpRoute) String() string {
	return fmt.Sprintf("%s %s \n", oRoute.Info(), oRoute.Status())
}

func (oRoute *OpRoute) Info() string {
	return fmt.Sprintf("[%s] %s ",
		oRoute.GetKind(),
		oRoute.GetName())
}

func (oRoute *OpRoute) Status() string {
	return fmt.Sprintf("%s %s (%v) [%s]",
		oRoute.Route.Spec.Host,
		oRoute.Route.Spec.To.Name,
		oRoute.Route.Spec.Port,
		oRoute.Route.Spec.Path)
}

func (oRoute *OpRoute) GetName() string {
	return oRoute.Route.Name
}

func (oRoute *OpRoute) GetKind() string {
	return oRoute.Route.Kind
}

func (oRoute *OpRoute) ToYaml() (string, error) {
	var contentBuilder strings.Builder
	err := converter.ObjToYaml(oRoute.Route, &contentBuilder, true, true)
	if err != nil {
		return "", err
	}
	return contentBuilder.String(), nil
}

func (oRoute *OpRoute) FromData(data []byte) error {
	_, _, err := converter.YamlToObject(data, false, oRoute.Route)
	if err != nil {
		return err
	}
	return nil
}
