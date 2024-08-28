package models

import "gorm.io/gorm"

type CartProducts struct {
	gorm.Model
	CartID    uint `gorm:"primaryKey"`
	Cart      Cart
	ProductID uint `gorm:"primaryKey"`
	Product   Product

	Amount uint
}
