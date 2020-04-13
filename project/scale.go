package project

import (
	"fmt"
	"io"
	"k8s.io/client-go/rest"
)

func (op *OpenshiftProject) ScaleItems(replicas int, w io.Writer, namespace string, restConf *rest.Config) {
	for _, item := range op.Items {
		err := item.UpdateScale(replicas, namespace, restConf)
		if err != nil {
			// only Print on Error
			fmt.Fprint(w, err.Error())
		} else {
			fmt.Fprintf(w, "Scaled %s to %d \n", item.Info(), replicas)
		}
	}
}
