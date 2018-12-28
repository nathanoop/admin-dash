package main

import (
	"net/http"
	"os"

	"github.com/nathanoop/admin-dash/db"
	GinHTMLRender "github.com/nathanoop/admin-dash/gin_html_render"
	admsess "github.com/nathanoop/admin-dash/handlers/admsess"
	"github.com/nathanoop/admin-dash/middlewares"

	"github.com/gin-gonic/gin"
)

const (
	// Port at which the server starts listening
	Port = "8080"
)

func init() {
	db.Connect()
}

func main() {
	// Configure
	router := gin.Default()

	// Set html render options
	htmlRender := GinHTMLRender.New()
	htmlRender.Debug = gin.IsDebugging()
	htmlRender.Layout = "layouts/default"
	router.HTMLRender = htmlRender.Create()

	router.RedirectTrailingSlash = true
	router.RedirectFixedPath = true
	// Middlewares
	router.Use(middlewares.Connect)
	router.Use(middlewares.ErrorHandler)

	// Statics
	router.Static("/public", "./public")

	// Routes

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/login")
	})

	//Auth
	router.GET("/login", admsess.showLogin)

	// Start listening
	port := Port
	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}
	router.Run(":" + port)
}
