package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/nathanoop/admin-dash/dbhelpers"

	"github.com/gin-gonic/gin"
	"github.com/nathanoop/admin-dash/utils"
)

func Createadminuser(c *gin.Context) {
	token := c.Param("token")
	if token != "" {
		isValidToken, adminObj := dbhelpers.Validateadmintoken(c, token)
		if isValidToken == true {
			msgObj := c.Request.URL.Query()
			viewObj := utils.Notificationobj(msgObj)
			viewModal := utils.Admintoken{token, adminObj}
			c.HTML(http.StatusOK, "admin/adminform", gin.H{
				"title":    "Create New Admin",
				"message":  viewObj,
				"tokenobj": viewModal})
		} else {
			c.Redirect(http.StatusMovedPermanently, utils.SITE_URL_ADMIN_LOGIN+"?msg="+utils.ERR_LOGIN_INV_TOKEN)
		}
	} else {
		c.Redirect(http.StatusMovedPermanently, utils.SITE_URL_ADMIN_LOGIN+"?msg="+utils.ERR_LOGIN_INV_TOKEN)
	}
}

func Saveadminuser(c *gin.Context) {
	token := c.Param("token")
	if token != "" {
		isValidToken, adminObj := dbhelpers.Validateadmintoken(c, token)
		if isValidToken == true {
			firstname := c.PostForm("firstname")
			lastname := c.PostForm("lastname")
			username := c.PostForm("username")
			password := c.PostForm("password")
			confirmpassword := c.PostForm("confirmpassword")
			useremail := c.PostForm("useremail")
			createdby := adminObj.Id.Hex()
			mobile := c.PostForm("mobile")
			if firstname != "" && lastname != "" && username != "" && password != "" && confirmpassword != "" && useremail != "" && mobile != "" {
				if password == confirmpassword {
					msg := dbhelpers.Createadminuserutils(c, firstname, lastname, username, password, useremail, mobile, createdby)
					if msg != "" {
						c.Redirect(http.StatusMovedPermanently, utils.SITE_URL_ADMIN_CREATE+"/"+token+"?msg="+msg)
					} else {
						c.Redirect(http.StatusMovedPermanently, utils.SITE_URL_ADMIN_LIST+"/"+token)
					}
				} else {
					c.Redirect(http.StatusMovedPermanently, utils.SITE_URL_ADMIN_CREATE+"/"+token+"?msg="+utils.ERR_ADM_CREATE_PWD_ERR)
				}
			} else {
				c.Redirect(http.StatusMovedPermanently, utils.SITE_URL_ADMIN_CREATE+"/"+token+"?msg="+utils.ERR_ADM_CREATE_REQ_FIELDS)
			}

		} else {
			c.Redirect(http.StatusMovedPermanently, utils.SITE_URL_ADMIN_LOGIN+"?msg="+utils.ERR_LOGIN_INV_TOKEN)
		}
	} else {
		c.Redirect(http.StatusMovedPermanently, utils.SITE_URL_ADMIN_LOGIN+"?msg="+utils.ERR_LOGIN_INV_TOKEN)
	}
}

func Editadminuser(c *gin.Context) {
	token := c.Param("token")
	if token != "" {
		isValidToken, adminObj := dbhelpers.Validateadmintoken(c, token)
		if isValidToken == true {
			msgObj := c.Request.URL.Query()
			viewObj := utils.Notificationobj(msgObj)
			viewModal := utils.Admintoken{token, adminObj}
			adminId := c.Param("adminId")
			if adminId != adminObj.Id.Hex() {
				admObj, err := dbhelpers.Getadminuserbyidutils(c, adminId)
				if err != nil {
					c.Redirect(http.StatusMovedPermanently, utils.SITE_URL_ADMIN_DASHBOARD+"/token?msg="+utils.ERR_DASHBRD_INV_ADMINID)
				} else {
					log.Println("editing admin obj", admObj.Id.Hex())
				}
				c.HTML(http.StatusOK, "admin/admineditform", gin.H{
					"title":     "Edit Admin " + admObj.FirstName + " " + admObj.LastName,
					"message":   viewObj,
					"tokenobj":  viewModal,
					"editadmin": admObj})
			} else {
				admObj := adminObj
				c.HTML(http.StatusOK, "admin/admineditform", gin.H{
					"title":     "Edit Admin " + admObj.FirstName + " " + admObj.LastName,
					"message":   viewObj,
					"tokenobj":  viewModal,
					"editadmin": admObj})
			}
		} else {
			c.Redirect(http.StatusMovedPermanently, utils.SITE_URL_ADMIN_LOGIN+"?msg="+utils.ERR_LOGIN_INV_TOKEN)
		}
	} else {
		c.Redirect(http.StatusMovedPermanently, utils.SITE_URL_ADMIN_LOGIN+"?msg="+utils.ERR_LOGIN_INV_TOKEN)
	}
}

func Updateadminuser(c *gin.Context) {
	token := c.Param("token")
	if token != "" {
		isValidToken, adminObj := dbhelpers.Validateadmintoken(c, token)
		if isValidToken == true {
			adminId := c.Param("adminId")
			firstname := c.PostForm("firstname")
			lastname := c.PostForm("lastname")
			useremail := c.PostForm("useremail")
			modifiedby := adminObj.Id.Hex()
			mobile := c.PostForm("mobile")
			if firstname != "" && lastname != "" && useremail != "" && mobile != "" {
				msg := dbhelpers.Updateadminuserutils(c, firstname, lastname, useremail, mobile, modifiedby, adminId)
				if msg != "" {
					c.Redirect(http.StatusMovedPermanently, utils.SITE_URL_ADMIN_DASHBOARD+"/"+token+"?msg="+msg)
				} else {
					c.Redirect(http.StatusMovedPermanently, utils.SITE_URL_ADMIN_EDIT+"/"+token+"/"+adminId)
				}
			} else {
				c.Redirect(http.StatusMovedPermanently, utils.SITE_URL_ADMIN_EDIT+"/"+token+"/"+adminId+"?msg="+utils.ERR_ADM_CREATE_REQ_FIELDS)
			}
		} else {
			c.Redirect(http.StatusMovedPermanently, utils.SITE_URL_ADMIN_LOGIN+"?msg="+utils.ERR_LOGIN_INV_TOKEN)
		}
	} else {
		c.Redirect(http.StatusMovedPermanently, utils.SITE_URL_ADMIN_LOGIN+"?msg="+utils.ERR_LOGIN_INV_TOKEN)
	}
}

func Editsettingsadminuser(c *gin.Context) {
	token := c.Param("token")
	if token != "" {
		isValidToken, adminObj := dbhelpers.Validateadmintoken(c, token)
		if isValidToken == true {
			msgObj := c.Request.URL.Query()
			viewObj := utils.Notificationobj(msgObj)
			viewModal := utils.Admintoken{token, adminObj}
			adminId := c.Param("adminId")
			if adminId != adminObj.Id.Hex() {
				admObj, err := dbhelpers.Getadminuserbyidutils(c, adminId)
				if err != nil {
					c.Redirect(http.StatusMovedPermanently, utils.SITE_URL_ADMIN_DASHBOARD+"/token?msg="+utils.ERR_DASHBRD_INV_ADMINID)
				} else {
					log.Println("edit setting admin obj", admObj.Id.Hex())
				}
				c.HTML(http.StatusOK, "admin/adminsettingsform", gin.H{
					"title":     "Edit Settings " + admObj.FirstName + " " + admObj.LastName,
					"message":   viewObj,
					"tokenobj":  viewModal,
					"editadmin": admObj})
			} else {
				admObj := adminObj
				c.HTML(http.StatusOK, "admin/adminsettingsform", gin.H{
					"title":     "Edit Settings " + admObj.FirstName + " " + admObj.LastName,
					"message":   viewObj,
					"tokenobj":  viewModal,
					"editadmin": admObj})
			}
		} else {
			c.Redirect(http.StatusMovedPermanently, utils.SITE_URL_ADMIN_LOGIN+"?msg="+utils.ERR_LOGIN_INV_TOKEN)
		}
	} else {
		c.Redirect(http.StatusMovedPermanently, utils.SITE_URL_ADMIN_LOGIN+"?msg="+utils.ERR_LOGIN_INV_TOKEN)
	}
}

func Changeadminusername(c *gin.Context) {
	token := c.Param("token")
	if token != "" {
		isValidToken, adminObj := dbhelpers.Validateadmintoken(c, token)
		if isValidToken == true {
			adminId := c.Param("adminId")
			username := c.PostForm("username")
			if username != "" {
				modifiedby := adminObj.Id.Hex()
				msg := dbhelpers.Updateadminusernameutils(c, username, modifiedby, adminId)
				if msg != "" {
					c.Redirect(http.StatusMovedPermanently, utils.SITE_URL_ADMIN_DASHBOARD+"/"+token+"?msg="+msg)
				} else {
					c.Redirect(http.StatusMovedPermanently, utils.SITE_URL_SETTING+"/"+token+"/"+adminId)
				}
			} else {
				c.Redirect(http.StatusMovedPermanently, utils.SITE_URL_SETTING+"/"+token+"/"+adminId+"?msg="+utils.ERR_ADM_CREATE_REQ_FIELDS)
			}
		} else {
			c.Redirect(http.StatusMovedPermanently, utils.SITE_URL_ADMIN_LOGIN+"?msg="+utils.ERR_LOGIN_INV_TOKEN)
		}
	} else {
		c.Redirect(http.StatusMovedPermanently, utils.SITE_URL_ADMIN_LOGIN+"?msg="+utils.ERR_LOGIN_INV_TOKEN)
	}
}

func Changeadminpassword(c *gin.Context) {
	token := c.Param("token")
	if token != "" {
		isValidToken, adminObj := dbhelpers.Validateadmintoken(c, token)
		if isValidToken == true {
			password := c.PostForm("password")
			adminId := c.Param("adminId")
			confirmpassword := c.PostForm("confirmpassword")
			if password != "" && confirmpassword != "" {
				if password == confirmpassword {
					modifiedby := adminObj.Id.Hex()
					msg := dbhelpers.Updateadminuserpasswordutils(c, password, modifiedby, adminId)
					if msg != "" {
						c.Redirect(http.StatusMovedPermanently, utils.SITE_URL_ADMIN_DASHBOARD+"/"+token+"?msg="+msg)
					} else {
						c.Redirect(http.StatusMovedPermanently, utils.SITE_URL_SETTING+"/"+token+"/"+adminId)
					}
				} else {
					c.Redirect(http.StatusMovedPermanently, utils.SITE_URL_SETTING+"/"+token+"/"+adminId+"?msg="+utils.ERR_ADM_CREATE_PWD_ERR)
				}
			} else {
				c.Redirect(http.StatusMovedPermanently, utils.SITE_URL_SETTING+"/"+token+"/"+adminId+"?msg="+utils.ERR_ADM_CREATE_REQ_FIELDS)
			}
		} else {
			c.Redirect(http.StatusMovedPermanently, utils.SITE_URL_ADMIN_LOGIN+"?msg="+utils.ERR_LOGIN_INV_TOKEN)
		}
	} else {
		c.Redirect(http.StatusMovedPermanently, utils.SITE_URL_ADMIN_LOGIN+"?msg="+utils.ERR_LOGIN_INV_TOKEN)
	}
}

func Deleteadminuser(c *gin.Context) {
	token := c.Param("token")
	if token != "" {
		isValidToken, adminObj := dbhelpers.Validateadmintoken(c, token)
		if isValidToken == true {
			adminId := c.Param("adminId")
			modifiedby := adminObj.Id.Hex()
			if modifiedby != adminId {
				msg := dbhelpers.Deleteadminuserutils(c, adminId, modifiedby)
				if msg == "" {
					c.Redirect(http.StatusMovedPermanently, utils.SITE_URL_ADMIN_LIST+"/"+token)
				} else {
					c.Redirect(http.StatusMovedPermanently, utils.SITE_URL_ADMIN_LIST+"/"+token+"?msg="+msg)
				}
			} else {
				c.Redirect(http.StatusMovedPermanently, utils.SITE_URL_ADMIN_LIST+"/"+token+"?msg="+utils.ERR_ADM_DEL_SELF)
			}
		} else {
			c.Redirect(http.StatusMovedPermanently, utils.SITE_URL_ADMIN_LOGIN+"?msg="+utils.ERR_LOGIN_INV_TOKEN)
		}
	} else {
		c.Redirect(http.StatusMovedPermanently, utils.SITE_URL_ADMIN_LOGIN+"?msg="+utils.ERR_LOGIN_INV_TOKEN)
	}
}

func Listadminuser(c *gin.Context) {
	token := c.Param("token")
	if token != "" {
		isValidToken, adminObj := dbhelpers.Validateadmintoken(c, token)
		if isValidToken == true {
			msgObj := c.Request.URL.Query()
			viewObj := utils.Notificationobj(msgObj)
			viewModal := utils.Admintoken{token, adminObj}
			p := utils.Getqueryparams(msgObj, "p")
			page, err := strconv.Atoi(p)
			if err != nil {
				page = 0
			}
			if p == "" {
				page = 0
			}
			adms, err := dbhelpers.Listadminusersutils(c, page)
			if err != nil {
				msg := utils.ERR_ADMIN_LISTING
				viewObj = utils.Notificationobjfromstr(msg)
			}
			c.HTML(http.StatusOK, "admin/adminlist", gin.H{
				"title":     "List  Admins  ",
				"message":   viewObj,
				"tokenobj":  viewModal,
				"listadmin": adms})
		} else {
			c.Redirect(http.StatusMovedPermanently, utils.SITE_URL_ADMIN_LOGIN+"?msg="+utils.ERR_LOGIN_INV_TOKEN)
		}
	} else {
		c.Redirect(http.StatusMovedPermanently, utils.SITE_URL_ADMIN_LOGIN+"?msg="+utils.ERR_LOGIN_INV_TOKEN)
	}
}
