package models

import (
	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	Name        string
	Description string
	IsOpen      bool
}
