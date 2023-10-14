package subscriber

import (
	"translation-app/application/model"
)

type Subscriber interface {
	Handle(text model.TranslationText)
	Print()
}
