package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nathanoop/admin-dash/dbhelpers"
	"github.com/nathanoop/admin-dash/utils"
)

// Login
func Login(c *gin.Context) {
	msgObj := c.Request.URL.Query()
	viewObj := utils.Notificationobj(msgObj)
	viewModal := utils.Admintoken{}
	c.HTML(http.StatusOK, "admsess/login", gin.H{
		"title":    "Login",
		"message":  viewObj,
		"tokenobj": viewModal})
}

func Authenticate(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username != "" && password != "" {
		responseStr, adminId := dbhelpers.Authenticateadmin(c, username, password)
		if responseStr == "" {
			ua := c.Request.Header.Get("User-Agent")
			accessToken := dbhelpers.Getadmintoken(c, adminId, ua)
			if accessToken != "" {
				c.Redirect(http.StatusMovedPermanently, "/dashboard/"+accessToken)
			} else {
				c.Redirect(http.StatusMovedPermanently, "/?msg="+utils.ERR_LOGIN_INV_TOKEN)
			}
		} else {
			c.Redirect(http.StatusMovedPermanently, "/?msg="+responseStr)
		}
	} else {
		c.Redirect(http.StatusMovedPermanently, "/?msg="+utils.ERR_LOGIN_EMPTY_FIELDS)
	}
}

func Dashboard(c *gin.Context) {
	token := c.Param("token")
	if token != "" {
		isValidToken, adminObj := dbhelpers.Validateadmintoken(c, token)
		if isValidToken == true {
			msgObj := c.Request.URL.Query()
			viewObj := utils.Notificationobj(msgObj)
			viewModal := utils.Admintoken{token, adminObj}
			c.HTML(http.StatusOK, "admsess/dashboard", gin.H{
				"title":    "Admin Dashboard",
				"message":  viewObj,
				"tokenobj": viewModal})
		} else {
			c.Redirect(http.StatusMovedPermanently, "/?msg="+utils.ERR_LOGIN_INV_TOKEN)
		}
	} else {
		c.Redirect(http.StatusMovedPermanently, "/?msg="+utils.ERR_LOGIN_INV_TOKEN)
	}
}

func Logout(c *gin.Context) {
	token := c.Param("token")
	if token != "" {
		isValidToken := dbhelpers.Logoutadmin(c, token)
		if isValidToken == true {
			c.Redirect(http.StatusMovedPermanently, "/?msg="+utils.ERR_LOGGED_OUT)
		} else {
			c.Redirect(http.StatusMovedPermanently, "/?msg="+utils.ERR_LOGIN_INV_TOKEN)
		}
	} else {
		c.Redirect(http.StatusMovedPermanently, "/?msg="+utils.ERR_LOGIN_INV_TOKEN)
	}
}
