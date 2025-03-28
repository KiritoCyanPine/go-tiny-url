package logic

import "github.com/kiritocyanpine/go-tiny-url/persistant"

var TinyUrlInstance *TinyUrl

func CreateTinyUrl(persistance persistant.Persistant) *TinyUrl {
	if TinyUrlInstance != nil {
		return TinyUrlInstance
	}

	TinyUrlInstance = &TinyUrl{
		db: persistance,
	}

	return TinyUrlInstance
}

type TinyUrl struct {
	db persistant.Persistant
}
