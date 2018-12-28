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
	msg = c.Param("msg")
	if msg != "" {
		msg = "ERR"
		msgStr = utils.Errormessage(msg)
		log.Println(msg, msgStr)
	}
	msgObj := utils.Message{msg, msgStr}
	c.HTML(http.StatusOK, "admsess/login", gin.H{
		"title":   "Login",
		"Message": msgObj})
}

func Authenticate(c *gin.Context) {

}
