package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	// CollectionAdmin holds the name of the Admin user collection
	CollectionAdmin = "Admin"
)

// Admin model
type Admin struct {
	Id            bson.ObjectId `bson:"_id,omitempty"`
	FirstName     string
	LastName      string
	UserName      string
	Password      string `json:"password,omitempty"`
	HashPassword  []byte `json:"hashpassword,omitempty"`
	UserEmail     string
	Mobile        string
	AccountStatus bool
	CreatedBy     string
	CreatedOn     time.Time
	ModifedBy     string
	ModifiedOn    time.Time
}
