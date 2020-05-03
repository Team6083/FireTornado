package model

type Item struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Genre       int    `json:"genre"`
	Quantity    int    `json:"quantity"`
	Unit        int    `json:"unit"`
	StoragePos  string `json:"storage_pos"`
	Status      int    `json:"status"`
}
