package auth

import (
	"net/http"

	"github.com/nathanoop/admin-dash/models"

	"github.com/gin-gonic/gin"
)

// showLogin
func showLogin(c *gin.Context) {
	article := models.Article{}

	c.HTML(http.StatusOK, "auth/login", gin.H{
		"title":   "Login"
	})
}
