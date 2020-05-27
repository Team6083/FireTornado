package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Name            string        `json:"name"`
	Email           string        `json:"email"`
	Password        string        `json:"password"`
	Id              bson.ObjectId `bson:"_id,omitempty" json:"id"`
	PermissionLevel int           `json:"permission_level"`
}

func (database *Database) GetUserById(id string) (*User, error) {
	user := User{}
	err := database.DB.C("users").FindId(bson.ObjectIdHex(id)).One(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (database *Database) GetUserByName(name string) (*User, error) {
	user := User{}
	err := database.DB.C("users").Find(bson.M{"name": name}).One(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (database *Database) GetAllUsers() ([]User, error) {
	var users []User
	err := database.DB.C("users").Find(bson.M{}).All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (database *Database) SaveUser(user User) (*mgo.ChangeInfo, error) {
	info, err := database.DB.C("users").UpsertId(user.Id, user)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (database *Database) DeleteUser(user User) error {
	err := database.DB.C("users").RemoveId(user.Id)
	if err != nil {
		return err
	}
	return nil
}
