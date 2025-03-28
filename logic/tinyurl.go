package logic

import "github.com/kiritocyanpine/go-tiny-url/persistant"

var tinyUrlInstance *TinyUrl

func CreateTinyUrl(persistance persistant.Persistant) *TinyUrl {
	if tinyUrlInstance != nil {
		return tinyUrlInstance
	}

	tinyUrlInstance = &TinyUrl{
		db: persistance,
	}

	return tinyUrlInstance
}

func GetTinyUrlInstance() *TinyUrl {
	return tinyUrlInstance
}

type TinyUrl struct {
	db persistant.Persistant
}
