package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string   `gorm:"uniqueIndex; not null" json:"name"`
	Email        string   `gorm:"uniqueIndex; not null" json:"email"`
	PasswordHash string   `gorm:"not null" json:"password_hash"`
	IsActive     bool     `gorm:"default:true" json:"is_active"`
	Tickets      []Ticket `json:"tickets"`
}
