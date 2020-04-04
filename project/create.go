package project

import (
	"fmt"
	"k8s.io/client-go/rest"
)

func (op *OpenshiftProject) CreateAll(namespace string, restConf *rest.Config) {
	for _, item := range op.Items {
		err := item.Create(namespace, restConf)
		if err != nil {
			// only Print on Error
			fmt.Println(err.Error())
		} else {
			fmt.Printf("Created %s \n", item.Info())
		}
	}
}
