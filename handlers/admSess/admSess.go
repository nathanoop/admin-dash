package admSess

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// showLogin
func showLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "admSess/login", gin.H{
		"title": "Login"})
}
