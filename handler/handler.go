package handler

import (
	"encoding/json"
	"errors"
	"html/template"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kiritocyanpine/go-tiny-url/config"
	"github.com/kiritocyanpine/go-tiny-url/logic"
	"github.com/kiritocyanpine/go-tiny-url/persistant"
)

type TinyUrlHandler struct {
	logic  *logic.TinyUrl
	config *config.Configuration
}

func CreateTinyUrlHandler(tinyUrlLogic *logic.TinyUrl,
	configs *config.Configuration) TinyUrlHandler {
	return TinyUrlHandler{
		logic:  tinyUrlLogic,
		config: configs,
	}
}

func (h *TinyUrlHandler) WelcomePageRequestHander(c *gin.Context) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	somedata := gin.H{
		"message": "Welcome to the page",
		"actions": "go to home page to get started",
	}
	tmpl.Execute(c.Writer, somedata)
}

func (h *TinyUrlHandler) GenericErrorPage(c *gin.Context, message responseMessage) {

}

func (h *TinyUrlHandler) ShortenNewURLRequestHander(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		c.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	var shortingData ShortenUrlRequest
	if err := json.Unmarshal(data, &shortingData); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	url := shortingData.Url
	queryPath, err := h.logic.AddNewUrlQuery(url)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := ShortnerUrlResponse{
		Url: h.config.HostAddress + "/" + queryPath,
	}

	c.JSON(http.StatusOK, response)
}

func (h *TinyUrlHandler) RedirectShortUrl(c *gin.Context) {
	queryKey := c.Param("queryKey")

	originalUrl, err := h.logic.GetOriginalUrl(queryKey)
	if err != nil {
		if errors.Is(err, persistant.ErrKeyNotFound) {
			h.GenericErrorPage(c, keyNotExistOrExpired)
			return
		}

		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	originalUrl = sanitizeUrl(originalUrl)

	tmpl := template.Must(template.ParseFiles("templates/redirect.html"))

	urlData := gin.H{
		"url": originalUrl,
	}

	tmpl.Execute(c.Writer, urlData)
}

func sanitizeUrl(s string) string {
	if strings.HasPrefix(s, "http") {
		return s
	}

	return "https://" + s
}
