package models

type Item struct {
	ItemID      int    `json:"item_id" gorm:"primaryKey"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     int    `json:"order_id"`
}