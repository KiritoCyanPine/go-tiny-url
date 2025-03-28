package logic

import (
	"errors"
	"time"

	"github.com/kiritocyanpine/go-tiny-url/persistant"
)

func (app *TinyUrl) AddNewUrlQuery(url string) (string, error) {
	urlQuery, err := formatUrlHash(getHasedValue(url))
	if err != nil {
		return "", err
	}

	if err := app.db.Set(urlQuery, url); err != nil {
		if !errors.Is(err, persistant.ErrKeyCollision) {
			return "", err
		}

		return app.AddNewUrlQuery(url + time.Now().String())
	}

	return urlQuery, nil
}

func (app *TinyUrl) GetOriginalUrl(query string) (string, error) {
	urlObject, err := app.db.Get(query)
	if err != nil {
		return "", err
	}

	url, ok := urlObject.(string)
	if !ok {
		return "", ErrAssertionFailed
	}

	return url, nil
}
