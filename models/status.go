package models

import "gorm.io/gorm"

type Status struct {
	gorm.Model
	Value string
}
