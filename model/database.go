package model

import "gopkg.in/mgo.v2"

type Database struct {
	Session  *mgo.Session
	DialInfo *mgo.DialInfo
	DB       *mgo.Database
}
