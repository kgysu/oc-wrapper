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
