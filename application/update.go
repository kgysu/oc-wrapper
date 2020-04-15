package application

import (
	"fmt"
	"io"
	"k8s.io/client-go/rest"
)

func (app *Application) UpdateItems(w io.Writer, namespace string, restConf *rest.Config) {
	for _, item := range app.Items {
		err := item.Update(namespace, restConf)
		if err != nil {
			// only Print on Error
			fmt.Fprintln(w, err.Error())
		} else {
			fmt.Fprintf(w, "Updated %s \n", item.Info())
		}
	}
}
