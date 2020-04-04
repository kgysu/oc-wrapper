package project

import (
	"fmt"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

func (op *OpenshiftProject) DeleteAll(namespace string, restConf *rest.Config, options *v12.DeleteOptions) error {
	for _, item := range op.Items {
		err := item.Delete(namespace, restConf, options)
		if err != nil {
			return err
		}
		fmt.Printf("Deleted %s \n", item.Info())
	}
	return nil
}
