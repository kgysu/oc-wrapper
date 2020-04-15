package application

import (
	"fmt"
	"io"
	"k8s.io/client-go/rest"
)

func (app *Application) ScaleItems(replicas int32, w io.Writer, namespace string, restConf *rest.Config) {
	for _, item := range app.Items {
		if item.IsScalable() {
			err := item.UpdateScale(replicas, namespace, restConf)
			if err != nil {
				// only Print on Error
				fmt.Fprint(w, err.Error())
			} else {
				fmt.Fprintf(w, "Scaled %s to %d \n", item.Info(), replicas)
			}
		}
	}
}

func (app *Application) ScaleItemsDefault(w io.Writer, namespace string, restConf *rest.Config) {
	for _, item := range app.Items {
		if item.IsScalable() {
			err := item.UpdateScale(item.GetScale(), namespace, restConf)
			if err != nil {
				// only Print on Error
				fmt.Fprint(w, err.Error())
			} else {
				fmt.Fprintf(w, "Scaled %s to %d \n", item.Info(), item.GetScale())
			}
		}
	}
}
