package project

type OpenshiftProject struct {
	Name         string
	Label        string
	Environments []OpenshiftProjectEnv
	Items        []OpenshiftItemInterface
}

func NewOpenshiftProject(name string) *OpenshiftProject {
	return &OpenshiftProject{
		Name:  name,
		Label: "project=" + name,
	}
}

func (op OpenshiftProject) GetItem(kind string, name string) OpenshiftItemInterface {
	for _, item := range op.Items {
		if kind == item.GetKind() && name == item.GetName() {
			return item
		}
	}
	return nil
}
