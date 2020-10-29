package models

import "time"

type Service struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name default:null"`
	CreatedAt time.Time `json:"created_at" gorm:"default:null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:null"`
}
