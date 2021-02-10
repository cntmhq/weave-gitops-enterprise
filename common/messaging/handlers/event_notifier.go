package handlers

import (
	"context"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
)

type EventNotifier struct {
	source string
	client cloudevents.Client
}

func NewEventNotifier(source string, client cloudevents.Client) *EventNotifier {
	return &EventNotifier{
		source: source,
		client: client,
	}
}

func (en *EventNotifier) Notify(eventType string, obj interface{}) error {
	event, ok := obj.(*v1.Event)
	if !ok {
		// Log an error but carry on
		log.Errorf("Expected object of type %T but received object of type %T instead.", &v1.Event{}, obj)
		return nil
	}

	e := cloudevents.NewEvent()
	e.SetID(uuid.New().String())
	e.SetType("Event")
	e.SetTime(time.Now())
	e.SetSource(en.source)
	if err := e.SetData(cloudevents.ApplicationJSON, event); err != nil {
		log.Errorf("Unable to set event as data: %v.", err)
		return err
	}
	if result := en.client.Send(context.Background(), e); cloudevents.IsUndelivered(result) {
		log.Errorf("Unable to send data: %v.", result)
		return result
	} else {
		log.Debugf("Successfully sent %s(%s), recipient acknowledged: %t.", e.Type(), e.ID(), cloudevents.IsACK(result))
		log.Tracef("Event payload: \n%v", e)
		return nil
	}
}