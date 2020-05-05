package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Item struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Genre       int    `json:"genre"`
	Quantity    int    `json:"quantity"`
	Unit        string `json:"unit"`
	StoragePos  string `json:"storage_pos"`
	Status      int    `json:"status"`
}

func (database *Database) GetItemByName(name string) (*Item, error) {
	item := Item{}
	err := database.DB.C("items").Find(bson.M{"name": name}).One(&item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (database *Database) GetAllItem() ([]Item, error) {
	var items []Item
	err := database.DB.C("items").Find(bson.M{}).All(&items)
	if err != nil {
		return nil, err
	}
	return items, err
}

func (database *Database) SaveItem(item Item) (*mgo.ChangeInfo, error) {
	info, err := database.DB.C("items").Upsert(bson.M{"name": item.Name}, item)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (database *Database) DeleteItem(item Item) error {
	err := database.DB.C("items").Remove(bson.M{"name": item.Name})
	if err != nil {
		return err
	}
	return nil
}
