package v2

type OpenshiftProject struct {
	name      string
	namespace string
	items     []OpenshiftItem
}

func NewFromRemote(name string, namespace string) (*OpenshiftProject, error) {

	remoteItems, err := ListAllFromRemote(namespace)
	if err != nil {
		return nil, err
	}

	return &OpenshiftProject{
		name:      name,
		namespace: namespace,
		items:     remoteItems,
	}, nil
}

func (op OpenshiftProject) Create() {

}
