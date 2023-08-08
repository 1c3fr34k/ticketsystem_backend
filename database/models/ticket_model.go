package models

import (
	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	IsOpen      bool `gorm:"default:true"`
	UserID      uint `gorm:"not null"`
}
