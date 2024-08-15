package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID       uint `gorm:"primaryKey"`
	FullName string
}
