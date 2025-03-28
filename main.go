package main

import (
	"html/template"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kiritocyanpine/go-tiny-url/handler"
	"github.com/kiritocyanpine/go-tiny-url/logic"
	"github.com/kiritocyanpine/go-tiny-url/middlewares"

	"github.com/kiritocyanpine/go-tiny-url/persistant/inmemory"
)

func initGinDependencies(r *gin.Engine) {
	r.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper,
	})

	// initialize gin's static files server to use in HTML templates
	r.Static("/assets", "./templates/assets")

	r.LoadHTMLGlob("templates/*.html")

	// middleware list
	r.Use(middlewares.AllowCrossOriginRequests())
}

func main() {
	r := gin.Default()

	p := inmemory.CreateDB()
	logic.CreateTinyUrl(p)

	// initialize all the dependancies that are required for the application to work
	initGinDependencies(r)

	// temporary route
	r.GET("/", handler.WelcomePageRequestHander)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
