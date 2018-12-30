package store

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nathanoop/admin-dash/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func Createadminsession(c *gin.Context, adminId string, uaStr string) string {
	ds := c.MustGet("db").(*mgo.Database)
	n := ds.C(models.CollectionAdminSession)
	obj_id := bson.NewObjectId()
	var admSess models.AdminSession
	admSess.Id = obj_id
	admSess.CreatedOn = time.Now()
	admSess.Expired = false
	browserStr := uaStr
	admSess.AccessHeader = browserStr
	//TODO- INVESTIGATION
	admSess.AccessIP = ""
	admSess.AdminId = adminId
	err := n.Insert(&admSess)
	if err != nil {
		return ""
	} else {
		log.Println("returning access token", obj_id.Hex())
		return obj_id.Hex()
	}
}

func Getadminsessionbytoken(c *gin.Context, token string) (isValidToken bool, adminId string) {
	ds := c.MustGet("db").(*mgo.Database)
	n := ds.C(models.CollectionAdminSession)
	var adminSessObj models.AdminSession
	var id = bson.ObjectIdHex(token)
	var tokenObj = bson.M{"_id": id, "expired": false}
	e := n.Find(tokenObj).One(&adminSessObj)
	if e != nil {
		isValidToken = false
		adminId = ""
		return
	} else {
		isValidToken = true
		adminId = adminSessObj.AdminId
		return
	}
}

func Removeadminsession(c *gin.Context, token string) bool {
	ds := c.MustGet("db").(*mgo.Database)
	n := ds.C(models.CollectionAdminSession)
	var adminSessObj models.AdminSession
	var id = bson.ObjectIdHex(token)
	var tokenObj = bson.M{"_id": id}
	e := n.Find(tokenObj).One(&adminSessObj)
	if e != nil {
		return false
	} else {
		log.Println("validated session token")
		e = n.Update(tokenObj, bson.M{"$set": bson.M{"expired": true, "expiredon": time.Now()}})
		if e != nil {
			log.Println("error updating session token", e)
			return false
		} else {
			log.Println("updated session token")
			return true
		}
	}
}
