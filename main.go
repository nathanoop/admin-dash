package main

import (
	"net/http"
	"os"

	"github.com/nathanoop/admin-dash/utils"

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

	router.GET(utils.SITE_URL_INDEX, func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, utils.SITE_URL_ADMIN_LOGIN)
	})

	//Auth
	router.GET(utils.SITE_URL_ADMIN_LOGIN, handlers.Login)
	router.POST(utils.SITE_URL_ADMIN_AUTHENTICATE, handlers.Authenticate)
	router.GET(utils.SITE_URL_ADMIN_DASHBOARD+"/:token", handlers.Dashboard)
	router.GET(utils.SITE_URL_ADMIN_LOGOUT+"/:token", handlers.Logout)

	router.GET(utils.SITE_URL_ADMIN_CREATE+"/:token", handlers.Createadminuser)
	router.POST(utils.SITE_URL_ADMIN_SAVE+"/:token", handlers.Saveadminuser)
	router.GET(utils.SITE_URL_ADMIN_LIST+"/:token", handlers.Listadminuser)
	router.GET(utils.SITE_URL_ADMIN_EDIT+"/:token/:adminId", handlers.Editadminuser)
	router.GET(utils.SITE_URL_PROFILE+"/:token/:adminId", handlers.Editadminuser)
	router.POST(utils.SITE_URL_ADMIN_UPDATE+"/:token/:adminId", handlers.Updateadminuser)
	router.GET(utils.SITE_URL_SETTING+"/:token/:adminId", handlers.Editsettingsadminuser)
	router.POST(utils.SITE_URL_ADMIN_CHANGE_USERNAME+"/:token/:adminId", handlers.Changeadminusername)
	router.POST(utils.SITE_URL_ADMIN_CHANGE_PASSWORD+"/:token/:adminId", handlers.Changeadminpassword)
	router.GET(utils.SITE_URL_ADMIN_DELETE+"/:token/:adminId", handlers.Deleteadminuser)

	// Start listening
	port := Port
	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}
	router.Run(":" + port)
}
