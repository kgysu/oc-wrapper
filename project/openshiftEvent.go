package project

import (
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type OpenshiftEvent struct {
	name  string
	event v1.Event
}

func fromEvent(event v1.Event) OpenshiftEvent {
	return OpenshiftEvent{
		name:  event.Name,
		event: event,
	}
}

func (oEvent OpenshiftEvent) setEvent(event v1.Event) {
	oEvent.name = event.Name
	oEvent.event = event
}

func (oEvent OpenshiftEvent) GetName() string {
	return oEvent.name
}

func (oEvent OpenshiftEvent) GetKind() string {
	return EventKey
}

func (oEvent OpenshiftEvent) GetStatus() string {
	return oEvent.event.Message
}

func (oEvent OpenshiftEvent) GetEvent() v1.Event {
	return oEvent.event
}

func (oEvent OpenshiftEvent) Create(namespace string) error {
	return nil
}

func (oEvent OpenshiftEvent) Update(namespace string) error {
	return nil
}

func (oEvent OpenshiftEvent) Delete(namespace string, options v12.DeleteOptions) error {
	return nil
}
