package handler

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func WelcomePageRequestHander(c *gin.Context) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	somedata := gin.H{
		"message": "Welcome to the page",
		"actions": "go to home page to get started",
	}
	tmpl.Execute(c.Writer, somedata)
}

func ShortenNewURLRequestHander(c *gin.Context) {

	data, err := c.GetRawData()
	if err != nil {
		c.AbortWithStatus(http.StatusUnprocessableEntity)
	}

	var shortingData ShortenUrlRequest
	if err := json.Unmarshal(data, &shortingData); err != nil {
		c.AbortWithError()
	}

}
