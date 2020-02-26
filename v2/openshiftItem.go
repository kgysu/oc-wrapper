package v2

import "encoding/json"

type OpenshiftItem struct {
	name     string
	kind     string
	itemJson string
}

func NewOpenshiftItem(name string, kind string, data string) OpenshiftItem {
	return OpenshiftItem{
		name:     name,
		kind:     kind,
		itemJson: data,
	}
}

func NewOpenshiftItemR(name string, kind string, data string) *OpenshiftItem {
	return &OpenshiftItem{
		name:     name,
		kind:     kind,
		itemJson: data,
	}
}

// Methods

func (oi OpenshiftItem) GetName() string {
	return oi.name
}

func (oi OpenshiftItem) GetKind() string {
	return oi.kind
}

func (oi OpenshiftItem) GetJson() string {
	return oi.itemJson
}

func (oi OpenshiftItem) ParseTo(v interface{}) error {
	return json.Unmarshal([]byte(oi.itemJson), v)
}

func (oi OpenshiftItem) SetJsonDataFrom(v interface{}) error {
	marshalled, err := json.Marshal(v)
	if err != nil {
		return err
	}
	oi.itemJson = string(marshalled)
	return nil
}

// Sorting by Kind
type ByKind []OpenshiftItem

func (a ByKind) Len() int {
	return len(a)
}

func (a ByKind) Less(i, j int) bool {
	if a[i].kind == ServiceAccountKey {
		return true
	}
	if a[i].kind == RoleKey && a[j].kind != ServiceAccountKey {
		return true
	}
	if a[i].kind == RoleBindingKey && a[j].kind != ServiceAccountKey && a[j].kind != RoleKey {
		return true
	}
	if a[i].kind == ConfigMapKey && a[j].kind != ServiceAccountKey && a[j].kind != RoleKey && a[j].kind != RoleBindingKey {
		return true
	}
	if a[i].kind == DeploymentConfigKey && a[j].kind != ConfigMapKey && a[j].kind != ServiceAccountKey &&
		a[j].kind != RoleKey && a[j].kind != RoleBindingKey {
		return true
	}
	if a[i].kind == StatefulSetKey && a[j].kind != DeploymentConfigKey && a[j].kind != ConfigMapKey &&
		a[j].kind != ServiceAccountKey && a[j].kind != RoleKey && a[j].kind != RoleBindingKey {
		return true
	}
	if a[i].kind == ServiceKey && a[j].kind != StatefulSetKey && a[j].kind != DeploymentConfigKey &&
		a[j].kind != ConfigMapKey && a[j].kind != ServiceAccountKey && a[j].kind != RoleKey && a[j].kind != RoleBindingKey {
		return true
	}
	return false
}

func (a ByKind) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
