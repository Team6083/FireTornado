package model

import "gopkg.in/mgo.v2/bson"

type User struct {
	Name            string        `json:"name"`
	Email           string        `json:"email"`
	Password        string        `json:"password"`
	Id              bson.ObjectId `bson:"id, omitempty" json:"id"`
	PermissionLevel int           `json:"permission_level"`
}
