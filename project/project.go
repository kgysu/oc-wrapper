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
