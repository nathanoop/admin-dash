package admSess

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// showLogin
func showLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "admsess/login", gin.H{
		"title": "Login"})
}
