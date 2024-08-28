package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Title  string
	Price  uint
	Orders []*Order `gorm:"many2many:order_products;"`
}
