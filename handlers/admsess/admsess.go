package admsess

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nathanoop/admin-dash/utils"
)

// Login
func Login(c *gin.Context) {
	msg, msgStr := "", ""
	msgObj := c.Request.URL.Query()
	msg = msgObj["msg"][0]
	if msg != "" {
		msgStr = utils.Errormessage(msg)
		msg = "ERR"
		log.Println(msg, msgStr)
	}
	viewObj := utils.Message{msg, msgStr}
	c.HTML(http.StatusOK, "admsess/login", gin.H{
		"title":   "Login",
		"message": viewObj})
}

func Authenticate(c *gin.Context) {

}
