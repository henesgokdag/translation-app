package subscriber

import (
	"translation-app/application/model"
)

type GermanTranslation struct {
	translatedText string
}

func (g *GermanTranslation) Handle(translationText model.TranslationText) {
	if translationText.Text == "Hi Everyone." {
		g.translatedText = "Hallo allerseits."
		g.Print()
		return
	}
	if translationText.Text == "Today We Will Talk About Observer Pattern." {
		g.translatedText = "Heute werden wir über das Beobachtermuster sprechen."
		g.Print()
	}
}

func (g *GermanTranslation) Print() {
	println(g.translatedText)
}
