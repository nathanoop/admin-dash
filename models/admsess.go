package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	// CollectionAdminSession holds the name of the Admin user login session collection
	CollectionAdminSession = "AdminSession"
)

// AdminSession model
type AdminSession struct {
	Id           bson.ObjectId `bson:"_id,omitempty"`
	AdminId      string
	AccessHeader string
	AccessIP     string
	CreatedOn    time.Time
	ExpiredOn    time.Time
	Expired      bool
}
