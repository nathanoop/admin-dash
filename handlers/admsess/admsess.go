package admsess

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nathanoop/admin-dash/utils"
)

// Login
func Login(c *gin.Context) {
	msgObj := c.Request.URL.Query()
	viewObj := utils.Notificationobj(msgObj)
	c.HTML(http.StatusOK, "admsess/login", gin.H{
		"title":   "Login",
		"message": viewObj})
}

func Authenticate(c *gin.Context) {

}
