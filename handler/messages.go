package handler

import "net/http"

type responseMessage struct {
	message string
	code    int
}

var (
	keyNotExistOrExpired = responseMessage{
		message: "The requested url link does not exist or has expired",
		code:    http.StatusNotFound,
	}
)
