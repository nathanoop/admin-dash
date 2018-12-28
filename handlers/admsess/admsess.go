package admsess

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login
func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "admsess/login", gin.H{
		"title": "Login"})
}
