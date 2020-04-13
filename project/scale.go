package project

func (op *OpenshiftProject) ScaleItems(replicas int) {
	for _, item := range op.Items {
		item.Scale(replicas)
	}
}
