package v2

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type OpenshiftProject struct {
	name      string
	namespace string
	items     []OpenshiftItem
}

func NewFromRemote(name string, namespace string, itemTypes string, options v1.ListOptions) (*OpenshiftProject, error) {
	remoteItems, err := ListAllRemoteItems(namespace, itemTypes, options)
	if err != nil {
		return nil, err
	}
	return &OpenshiftProject{
		name:      name,
		namespace: namespace,
		items:     remoteItems,
	}, nil
}

func NewFromLocal(namespace string, folderPath string, debug bool) (*OpenshiftProject, error) {
	localProject, err := parseProjectFile(folderPath)
	if err != nil {
		return nil, err
	}

	remoteItems, err := ListAllLocalItems(namespace, folderPath, debug)
	if err != nil {
		return nil, err
	}
	return &OpenshiftProject{
		name:      localProject.Name,
		namespace: namespace,
		items:     remoteItems,
	}, nil
}

// Methods

func (op OpenshiftProject) GetName() string {
	return op.name
}

func (op OpenshiftProject) GetNamespace() string {
	return op.namespace
}

func (op OpenshiftProject) GetItems() []OpenshiftItem {
	return op.items
}

func (op OpenshiftProject) Create() error {
	createdItems, err := CreateAllItemsRemote(op.namespace, op.items)
	if err != nil {
		return err
	}
	op.items = createdItems
	return nil
}

func (op OpenshiftProject) Update() error {
	createdItems, err := UpdateAllItemsRemote(op.namespace, op.items)
	if err != nil {
		return err
	}
	op.items = createdItems
	return nil
}

func (op OpenshiftProject) CreateOrUpdate() error {
	createdItems, err := CreateOrUpdateAllRemoteItems(op.namespace, op.items)
	if err != nil {
		return err
	}
	op.items = createdItems
	return nil
}

func (op OpenshiftProject) Delete(options v1.DeleteOptions) error {
	return DeleteAllRemoteItems(op.namespace, op.items, options)
}

func (op OpenshiftProject) Scale(replicas int) error {
	scaledItems, err := ScaleDeploymentConfigsAndStatefulSets(op.namespace, op.items, replicas)
	if err != nil {
		return err
	}
	op.items = scaledItems
	return nil
}
