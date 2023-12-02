package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID              uint      `gorm:"primaryKey;autoIncrement"`
	UserID          uint      `gorm:"not null"`
	TransactionDate time.Time `gorm:"not null;default:current_timestamp"`
	CreatedAt       time.Time `gorm:"not null;default:current_timestamp"`
	UpdatedAt       time.Time `gorm:"not null;default:current_timestamp"`
	DeletedAt       gorm.DeletedAt
}
