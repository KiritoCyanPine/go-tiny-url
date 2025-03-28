package main

import (
	"html/template"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kiritocyanpine/go-tiny-url/config"
	"github.com/kiritocyanpine/go-tiny-url/handler"
	"github.com/kiritocyanpine/go-tiny-url/logic"
	"github.com/kiritocyanpine/go-tiny-url/middlewares"
	"github.com/kiritocyanpine/go-tiny-url/persistant"

	"github.com/kiritocyanpine/go-tiny-url/persistant/inmemory"
)

type AppDependencies struct {
	handler handler.TinyUrlHandler
	logic   logic.TinyUrl
	config  config.Configuration
	db      persistant.Persistant
}

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

func initApp() *AppDependencies {
	configInstance := config.GetConfigurations()
	dbInstance := inmemory.CreateDB()
	logicInstance := logic.CreateTinyUrl(dbInstance)
	handlerInstance := handler.CreateTinyUrlHandler(logicInstance, &configInstance)

	return &AppDependencies{
		handler: handlerInstance,
		logic:   *logicInstance,
		config:  configInstance,
		db:      dbInstance,
	}
}

func main() {
	r := gin.Default()

	app := initApp()

	// initialize all the dependancies that are required for the application to work
	initGinDependencies(r)

	// temporary route
	r.GET("/", app.handler.WelcomePageRequestHander)
	r.POST("/tinify", app.handler.ShortenNewURLRequestHander)
	r.GET("/:queryKey", app.handler.RedirectShortUrl)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
