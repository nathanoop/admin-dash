package admsess

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nathanoop/admin-dash/utils"
)

// Login
func Login(c *gin.Context) {
	msg, msgStr, cls := "", "", ""
	var viewObj = utils.Message{}
	msgObj := c.Request.URL.Query()
	if msgObj != nil && msgObj["msg"] != nil {
		msg = msgObj["msg"][0]
		if msg != "" {
			msgStr = utils.Errormessage(msg)
			cls = "alert-danger"
		}
		viewObj = utils.Message{msg, msgStr, cls}
	}
	c.HTML(http.StatusOK, "admsess/login", gin.H{
		"title":   "Login",
		"message": viewObj})
}

func Authenticate(c *gin.Context) {

}
