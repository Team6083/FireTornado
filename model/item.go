package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Item struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Category    string        `json:"category"`
	Genre       int           `json:"genre"`
	Quantity    int           `json:"quantity"`
	Unit        string        `json:"unit"`
	StoragePos  string        `json:"storage_pos"`
	Status      int           `json:"status"`
	Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
}

func (database *Database) GetItemById(id string) (*Item, error) {
	item := Item{}
	err := database.DB.C("items").FindId(bson.ObjectIdHex(id)).One(&item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (database *Database) GetItemsByCategory(categoryName string) ([]Item, error) {
	var items []Item
	err := database.DB.C("items").Find(bson.M{"category": categoryName}).All(&items)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (database *Database) GetAllItem() ([]Item, error) {
	var items []Item
	err := database.DB.C("items").Find(bson.M{}).All(&items)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (database *Database) SaveItem(item *Item) (*mgo.ChangeInfo, error) {
	info, err := database.DB.C("items").UpsertId(item.Id, item)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (database *Database) DeleteItem(id string) error {
	err := database.DB.C("items").RemoveId(bson.ObjectIdHex(id))
	if err != nil {
		return err
	}
	return nil
}
