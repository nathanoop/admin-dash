package dbhelpers

import (
	"admin-dash/models"
	"log"

	"github.com/nathanoop/admin-dash/utils"

	"github.com/gin-gonic/gin"
	"github.com/nathanoop/admin-dash/store"
	"golang.org/x/crypto/bcrypt"
)

func Authenticateadmin(c *gin.Context, userName string, password string) (msgCode string, adminId string) {
	msgCode, adminId = "", ""
	if userName != "" && password != "" {
		result, err := store.Getuserbyname(c, userName)
		if err != nil {
			return utils.ERR_LOGIN_INV_USR, adminId
		} else {
			if result.AccountStatus == false {
				return utils.ERR_LOGIN_DELETED_ACCNT, adminId
			} else {
				// Validate password
				err = bcrypt.CompareHashAndPassword(result.HashPassword, []byte(password))
				if err != nil {
					return utils.ERR_LOGIN_INC_PSSWD, adminId
				} else {
					adminId = result.Id.Hex()
					return msgCode, adminId
				}
			}
		}
	} else {
		return utils.ERR_LOGIN_INV_USR_PSS, adminId
	}
}

func Getadmintoken(c *gin.Context, adminId string, uaStr string) string {
	if adminId != "" {
		accessToken := store.Createadminsession(c, adminId, uaStr)
		return accessToken
	} else {
		return ""
	}
}

func Validateadmintoken(c *gin.Context, token string) (isValidToken bool, adminObj models.Admin) {
	adminId := ""
	if token != "" {
		isValidToken, adminId = store.Getadminsessionbytoken(c, token)
		if adminId != "" {
			adminObj, err := store.Getadminuserbyid(c, adminId)
			if err != nil {
				log.Println("unable to validate token ", token)
			} else {
				log.Println("validate token in admin", adminObj.UserName)
				return isValidToken, adminObj
			}
		}
	}
	return
}
func Logoutadmin(c *gin.Context, token string) bool {
	if token != "" {
		isValidToken := store.Removeadminsession(c, token)
		return isValidToken
	} else {
		return false
	}
}
