package project

import (
	"fmt"
	"io"
	"k8s.io/client-go/rest"
)

func (op *OpenshiftProject) CreateItems(w io.Writer, namespace string, restConf *rest.Config) {
	for _, item := range op.Items {
		err := item.Create(namespace, restConf)
		if err != nil {
			// only Print on Error
			fmt.Fprintln(w, err.Error())
		} else {
			fmt.Fprintf(w, "Created %s \n", item.Info())
		}
	}
}
