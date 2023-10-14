package publisher

import (
	"translation-app/application/subscriber"
)

type Publisher interface {
	AddSubscriber(subscriber subscriber.Subscriber)
	RemoveSubscriber(subscriber subscriber.Subscriber) error
	Notify()
	SetText(string)
}
