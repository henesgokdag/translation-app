package publisher

import (
	"errors"
	"translation-app/application/model"
	"translation-app/application/subscriber"
)

type EnglishText struct {
	Subs            []subscriber.Subscriber
	translationText model.TranslationText
}

func NewText() EnglishText {
	return EnglishText{
		Subs:            make([]subscriber.Subscriber, 0),
		translationText: model.TranslationText{},
	}
}

func (t *EnglishText) AddSubscriber(subscriber subscriber.Subscriber) {
	t.Subs = append(t.Subs, subscriber)

}
func (t *EnglishText) RemoveSubscriber(subscriber subscriber.Subscriber) error {
	ok := false
	for i := range t.Subs {
		if t.Subs[i] == subscriber {
			ok = true
			t.Subs = append(t.Subs[:i], t.Subs[i+1:]...)
		}
	}
	if !ok {
		return errors.New("subscriber not found in subs list")
	}
	return nil
}
func (t *EnglishText) Notify() {
	for i := range t.Subs {
		t.Subs[i].Handle(t.translationText)
	}
}

func (t *EnglishText) SetText(text string) {
	t.translationText.Text = text
}
