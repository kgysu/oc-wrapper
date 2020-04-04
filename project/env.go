package project

type OpenshiftProjectEnv struct {
	Namespace string
	Envs      map[string]string
}

func NewOpenshiftProjectEnv(namespace string) OpenshiftProjectEnv {
	return OpenshiftProjectEnv{Namespace: namespace}
}
