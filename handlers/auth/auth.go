package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// showLogin
func showLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "auth/login", gin.H{
		"title": "Login"})
}
