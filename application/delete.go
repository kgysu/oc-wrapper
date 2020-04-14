package application

import (
	"fmt"
	"io"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

func (app *Application) DeleteItems(w io.Writer, namespace string, restConf *rest.Config, options *v12.DeleteOptions) {
	for _, item := range app.Items {
		err := item.Delete(namespace, restConf, options)
		if err != nil {
			// only Print on Error
			fmt.Fprintln(w, err.Error())
		} else {
			fmt.Fprintf(w, "Deleted %s \n", item.Info())
		}
	}
}
