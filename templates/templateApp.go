package templates

import (
	"github.com/kgysu/oc-wrapper/appitem"
	"github.com/kgysu/oc-wrapper/application"
	"github.com/kgysu/oc-wrapper/items"
)

func GetSampleApp() application.Application {
	return GetTemplateApp("sample")
}

func GetTemplateApp(name string) application.Application {
	return application.Application{
		Name:  name,
		Label: "app=" + name,
		Items: []appitem.AppItem{
			items.NewOpDeploymentConfig(GetTemplateDeploymentConfig(name)),
			items.NewOpService(GetTemplateService(name)),
			items.NewOpRoute(GetTemplateRoute(name)),
			items.NewOpStatefulSet(GetTemplateStatefulSet(name)),
			items.NewOpServiceAccount(GetTemplateServiceAccount(name)),
			items.NewOpRole(GetTemplateRole(name)),
			items.NewOpRoleBinding(GetTemplateRoleBinding(name)),
			items.NewOpConfigMap(GetTemplateConfigMap(name)),
		},
	}
}
