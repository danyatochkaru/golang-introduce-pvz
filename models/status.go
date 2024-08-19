package models

import "gorm.io/gorm"

type StatusValue string

const (
	StatusPending    StatusValue = "pending"
	StatusInProgress StatusValue = "in_progress"
	StatusCompleted  StatusValue = "completed"
)

type Status struct {
	gorm.Model
	Value StatusValue
}
