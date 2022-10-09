package models

import (
	"time"
)

type Order struct {
	OrderID      uint   `json:"order_id" gorm:"primaryKey"`
	CustomerName string `json:"customer_name"`
	OrderAt      time.Time   `json:"order_at" gorm:"autoCreateTime"`
	Items []Item `json:"items" gorm:"foreignKey:OrderID"`
}