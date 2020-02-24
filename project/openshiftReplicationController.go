package project

import (
	"fmt"
	"github.com/kgysu/oc-wrapper/wrapper"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type OpenshiftReplicationController struct {
	name                  string
	replicationController v1.ReplicationController
}

func fromReplicationController(replicationController v1.ReplicationController) OpenshiftReplicationController {
	return OpenshiftReplicationController{
		name:                  replicationController.Name,
		replicationController: replicationController,
	}
}

func (oReplicationController OpenshiftReplicationController) setReplicationController(replicationController v1.ReplicationController) {
	oReplicationController.name = replicationController.Name
	oReplicationController.replicationController = replicationController
}

func (oReplicationController OpenshiftReplicationController) GetName() string {
	return oReplicationController.name
}

func (oReplicationController OpenshiftReplicationController) GetKind() string {
	return ReplicationControllerKey
}

func (oReplicationController OpenshiftReplicationController) GetStatus() string {
	return fmt.Sprintf("%d (%d/%d)", oReplicationController.replicationController.Status.Replicas,
		oReplicationController.replicationController.Status.ReadyReplicas,
		oReplicationController.replicationController.Status.AvailableReplicas)
}

func (oReplicationController OpenshiftReplicationController) GetReplicationController() v1.ReplicationController {
	return oReplicationController.replicationController
}

func (oReplicationController OpenshiftReplicationController) Create(namespace string) error {
	_, err := wrapper.CreateReplicationController(namespace, &oReplicationController.replicationController)
	if err != nil {
		return err
	}
	//oReplicationController.setReplicationController(createdReplicationController)
	return nil
}

func (oReplicationController OpenshiftReplicationController) Update(namespace string) error {
	_, err := wrapper.UpdateReplicationController(namespace, &oReplicationController.replicationController)
	if err != nil {
		return err
	}
	//oReplicationController.setReplicationController(updatedReplicationController)
	return nil
}

func (oReplicationController OpenshiftReplicationController) Delete(namespace string, options v12.DeleteOptions) error {
	return wrapper.DeleteReplicationController(namespace, oReplicationController.name, options)
}
