package store

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nathanoop/admin-dash/models"
	"github.com/nathanoop/admin-dash/utils"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func Getuserbyname(c *gin.Context, userName string) (r models.Admin, e error) {
	ds := c.MustGet("db").(*mgo.Database)
	n := ds.C(models.CollectionAdmin)
	var userNameObj = bson.M{"username": userName}
	e = n.Find(userNameObj).One(&r)
	return
}

func Getadminuserbyid(c *gin.Context, adminId string) (r models.Admin, e error) {
	ds := c.MustGet("db").(*mgo.Database)
	n := ds.C(models.CollectionAdmin)
	var userNameObj = bson.M{"_id": bson.ObjectIdHex(adminId)}
	e = n.Find(userNameObj).One(&r)
	return
}

func Saveadmin(c *gin.Context, firstname string, lastname string, username string, password string, hashpassword []byte, useremail string, mobile string, createdby string) string {
	ds := c.MustGet("db").(*mgo.Database)
	n := ds.C(models.CollectionAdmin)
	var adm models.Admin
	adm.FirstName = firstname
	adm.LastName = lastname
	adm.UserName = username
	adm.Password = password
	adm.HashPassword = hashpassword
	adm.UserEmail = useremail
	adm.Mobile = mobile
	adm.AccountStatus = true
	adm.CreatedBy = createdby
	adm.CreatedOn = time.Now()
	adm.ModifedBy = createdby
	adm.ModifiedOn = time.Now()
	err := n.Insert(&adm)
	if err != nil {
		return utils.ERR_ADM_CREATE_INS_ERR
	} else {
		return ""
	}
}

func Updateadmin(c *gin.Context, firstname string, lastname string, useremail string, mobile string, modifiedby string, adminId string) string {
	adm, err := Getadminuserbyid(c, adminId)
	if err != nil {
		return utils.ERR_DASHBRD_INV_ADMINID
	} else {
		ds := c.MustGet("db").(*mgo.Database)
		n := ds.C(models.CollectionAdmin)
		var adminIdObj = bson.M{"_id": bson.ObjectIdHex(adminId)}
		log.Println("updating admin user", adm.Id.Hex(), adminId)
		err = n.Update(adminIdObj, bson.M{"$set": bson.M{"firstname": firstname, "lastname": lastname, "useremail": useremail, "mobile": mobile, "modifedby": modifiedby, "modifiedon": time.Now()}})
		if err != nil {
			log.Println("error updating admin user", err)
			return utils.ERR_ADM_UPDATE_UP_ERR
		} else {
			return ""
		}
	}
}

func Updateadminusername(c *gin.Context, username string, modifiedby string, adminId string) string {
	adm, err := Getadminuserbyid(c, adminId)
	if err != nil {
		return utils.ERR_DASHBRD_INV_ADMINID
	} else {
		result, err := Getuserbyname(c, username)
		if err != nil {
			ds := c.MustGet("db").(*mgo.Database)
			n := ds.C(models.CollectionAdmin)
			var adminIdObj = bson.M{"_id": bson.ObjectIdHex(adminId)}
			log.Println("updating admin userName", adm.Id.Hex(), adminId)
			err = n.Update(adminIdObj, bson.M{"$set": bson.M{"username": username, "modifedby": modifiedby, "modifiedon": time.Now()}})
			if err != nil {
				log.Println("error updating admin username", err)
				return utils.ERR_ADM_UPDATE_UP_ERR
			} else {
				return ""
			}
		} else {
			log.Println("Duplicate user error", result.Id.Hex())
			return utils.ERR_ADM_CREATE_DUP_USER_ERR
		}
	}
}

func Updateadminuserpassword(c *gin.Context, password string, hashpassword []byte, modifiedby string, adminId string) string {
	adm, err := Getadminuserbyid(c, adminId)
	if err != nil {
		return utils.ERR_DASHBRD_INV_ADMINID
	} else {
		ds := c.MustGet("db").(*mgo.Database)
		n := ds.C(models.CollectionAdmin)
		var adminIdObj = bson.M{"_id": bson.ObjectIdHex(adminId)}
		log.Println("updating admin user  password", adm.Id.Hex(), adminId)
		err = n.Update(adminIdObj, bson.M{"$set": bson.M{"password": password, "hashpassword": hashpassword, "modifedby": modifiedby, "modifiedon": time.Now()}})
		if err != nil {
			log.Println("error updating admin user  password", err)
			return utils.ERR_ADM_UPDATE_UP_ERR
		} else {
			return ""
		}
	}
}

func Listadminusers(c *gin.Context, limit int, skip int) (adms []models.Admin, total int, err error) {
	ds := c.MustGet("db").(*mgo.Database)
	n := ds.C(models.CollectionAdmin)
	accountstatusObj := bson.M{"accountstatus": true}
	total, err = n.Find(accountstatusObj).Count()
	if skip > 0 {
		err := n.Find(accountstatusObj).Sort("-modifiedon").Skip(skip).Limit(limit).All(&adms)
		log.Println("skip listing err", err)
	} else {
		err := n.Find(accountstatusObj).Sort("-modifiedon").Limit(limit).All(&adms)
		log.Println("lmit listing err", err)
	}
	return
}

func Deleteadminuser(c *gin.Context, adminId string, modifiedby string) string {
	adm, err := Getadminuserbyid(c, adminId)
	if err != nil {
		return utils.ERR_DASHBRD_INV_ADMINID
	} else {
		ds := c.MustGet("db").(*mgo.Database)
		n := ds.C(models.CollectionAdmin)
		var adminIdObj = bson.M{"_id": bson.ObjectIdHex(adminId)}
		log.Println("deleting admin user ", adm.Id.Hex(), adminId)
		err = n.Update(adminIdObj, bson.M{"$set": bson.M{"accountstatus": false, "modifedby": modifiedby, "modifiedon": time.Now()}})
		if err != nil {
			log.Println("error deleting admin user ", err)
			return utils.ERR_ADM_DELETE_ERR
		} else {
			return ""
		}
	}
}
