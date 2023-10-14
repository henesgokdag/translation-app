package subscriber

import (
	"translation-app/application/model"
)

type TurkishTranslation struct {
	translatedText string
}

func (t *TurkishTranslation) Handle(translationText model.TranslationText) {
	if translationText.Text == "Hi Everyone." {
		t.translatedText = "Herkese Merhaba."
		t.Print()
		return
	}
	if translationText.Text == "Today We Will Talk About Observer Pattern." {
		t.translatedText = "Bugün Observer Pattern Hakkında Konuşacağız."
		t.Print()
	}
}

func (t *TurkishTranslation) Print() {
	println(t.translatedText)
}
