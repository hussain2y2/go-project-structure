package models

import (
	"time"
)

type Book struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	CreatedAt time.Time `gorm:"default:null"`
	UpdatedAt time.Time `gorm:"default:null"`
}
