package dbhelpers

import (
	"admin-dash/models"
	"log"

	"github.com/nathanoop/admin-dash/utils"

	"github.com/gin-gonic/gin"
	"github.com/nathanoop/admin-dash/store"
	"golang.org/x/crypto/bcrypt"
)

func Createadminuserutils(c *gin.Context, firstname string, lastname string, username string, password string, useremail string, mobile string, createdby string) string {
	msg := ""
	result, err := store.Getuserbyname(c, username)
	if err != nil {
		hashpassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
		if err != nil {
			msg = utils.ERR_ADM_CREATE_PWD_HSH_ERR
		} else {
			msg = store.Saveadmin(c, firstname, lastname, username, password, hashpassword, useremail, mobile, createdby)
		}
	} else {
		log.Println("duplicate user ...", result.Id.Hex())
		msg = utils.ERR_ADM_CREATE_DUP_USER_ERR
	}
	return msg
}

func Getadminuserbyidutils(c *gin.Context, adminId string) (adm models.Admin, err error) {
	admobj, er := store.Getadminuserbyid(c, adminId)
	log.Println("adminbyuserid", admobj.Id.Hex(), er)
	return
}

func Updateadminuserutils(c *gin.Context, firstname string, lastname string, useremail string, mobile string, modifiedby string, adminId string) string {
	msg := store.Updateadmin(c, firstname, lastname, useremail, mobile, modifiedby, adminId)
	return msg
}

func Updateadminusernameutils(c *gin.Context, username string, modifiedby string, adminId string) string {
	msg := store.Updateadminusername(c, username, modifiedby, adminId)
	return msg
}

func Updateadminuserpasswordutils(c *gin.Context, password string, modifiedby string, adminId string) string {
	msg := ""
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		msg = utils.ERR_ADM_CREATE_PWD_HSH_ERR
	} else {
		msg = store.Updateadminuserpassword(c, password, hashpassword, modifiedby, adminId)
	}
	return msg
}

func Listadminusersutils(c *gin.Context, page int) (adms []models.Admin, err error) {
	limit := utils.DEF_PAGE_COUNT
	skip := 0
	if page > 0 {
		skip = limit * page
	}
	lstadms, er := store.Listadminusers(c, limit, skip)
	log.Println("adminbyuserid", len(lstadms), er)
	return
}

func Deleteadminuserutils(c *gin.Context, adminId string, modifiedby string) string {
	msg := store.Deleteadminuser(c, adminId, modifiedby)
	return msg
}
