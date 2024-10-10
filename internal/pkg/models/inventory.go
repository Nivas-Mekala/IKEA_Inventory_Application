package models

type Inventory struct {
	Article_id string `json:"art_id" gorm:"primaryKey"`
	Name       string `json:"name" gorm:"name"`
	Stock      string `json:"stock" gorm:"stock"`
}

type InventoryRequest struct {
	Inventory []struct {
		Article_id string `json:"art_id" `
		Name       string `json:"name" `
		Stock      string `json:"stock"`
	} `json:"inventory"`
}
