package main

import (
	"net/http"
	"os"

	"github.com/nathanoop/admin-dash/db"
	GinHTMLRender "github.com/nathanoop/admin-dash/gin_html_render"
	"github.com/nathanoop/admin-dash/handlers"
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
	router.GET("/login", handlers.Login)
	router.POST("/authenticate", handlers.Authenticate)
	router.GET("/dashboard/:token", handlers.Dashboard)
	router.GET("/logout/:token", handlers.Logout)

	router.GET("/admin/create/:token", handlers.Createadminuser)
	router.POST("/admin/save/:token", handlers.Saveadminuser)
	router.GET("/admin/list/:token", handlers.Listadminuser)
	router.GET("/admin/edit/:token/:adminId", handlers.Editadminuser)
	router.GET("/profile/:token/:adminId", handlers.Editadminuser)
	router.GET("/settings/:token/:adminId", handlers.Editsettingsadminuser)
	router.POST("/admin/update/:token/:adminId", handlers.Updateadminuser)
	router.POST("/admin/changeusername/:token/:adminId", handlers.Changeadminusername)
	router.POST("/admin/changepassword/:token/:adminId", handlers.Changeadminpassword)
	router.GET("/admin/delete/:token/:adminId", handlers.Deleteadminuser)

	// Start listening
	port := Port
	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}
	router.Run(":" + port)
}
