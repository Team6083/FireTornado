package model

import "gopkg.in/mgo.v2/bson"

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
