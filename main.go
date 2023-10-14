package main

import (
	"translation-app/application/publisher"
	"translation-app/application/subscriber"
)

func main() {
	text := publisher.NewText()
	turkishTranslation := &subscriber.TurkishTranslation{}
	germanTranslation := &subscriber.GermanTranslation{}

	text.AddSubscriber(turkishTranslation)
	text.AddSubscriber(germanTranslation)

	text.SetText("Hi Everyone.")
	text.Notify()
	println("")

	err := text.RemoveSubscriber(germanTranslation)
	if err != nil {
		println(err)
	}

	text.SetText("Today We Will Talk About Observer Pattern.")
	text.Notify()

	println("")
	text.AddSubscriber(germanTranslation)
	text.SetText("Today We Will Talk About Observer Pattern.")
	text.Notify()
}
