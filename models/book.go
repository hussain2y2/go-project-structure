package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	ID        uint      `json:"id", gorm:"primary_key"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at", gorm:"default:null"`
	UpdatedAt time.Time `json:"updated_at", gorm:"default:null"`
}
