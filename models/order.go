package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	FullName string
	StatusID uint
	Status   Status `gorm:"foreignKey:StatusID"`
}
