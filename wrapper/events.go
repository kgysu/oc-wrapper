package wrapper

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
)

type EventList []v1.Event

func ListEvents(ns string, options metav1.ListOptions) (EventList, error) {
	eventsApi, err := GetEventApi(ns)
	if err != nil {
		return nil, err
	}
	events, err := eventsApi.List(options)
	if err != nil {
		return nil, err
	}
	return events.Items, nil
}

func WatchEvents(ns string, options metav1.ListOptions) (watch.Interface, error) {
	eventsApi, err := GetEventApi(ns)
	if err != nil {
		return nil, err
	}
	return eventsApi.Watch(options)
}
