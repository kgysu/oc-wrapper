package application

import (
	"github.com/kgysu/oc-wrapper/appitem"
	"strings"
)

type Application struct {
	Name  string
	Label string
	Items []appitem.AppItem
}

func NewApp(name string) *Application {
	return &Application{
		Name:  name,
		Label: "app=" + name,
	}
}

func NewAppWithLabel(name, label string) *Application {
	return &Application{
		Name:  name,
		Label: label,
	}
}

func (app Application) GetItem(kind string, name string) appitem.AppItem {
	for _, item := range app.Items {
		if kind == item.GetKind() && name == item.GetName() {
			return item
		}
	}
	return nil
}

func (app Application) GetItemsByKinds(kinds string) []appitem.AppItem {
	var result []appitem.AppItem
	for _, item := range app.Items {
		if strings.Contains(kinds, item.GetKind()) || kinds == "" {
			result = append(result, item)
		}
	}
	return result
}

func (app Application) GetScalableItems() []appitem.AppItem {
	var result []appitem.AppItem
	for _, item := range app.Items {
		if item.IsScalable() {
			result = append(result, item)
		}
	}
	return result
}
