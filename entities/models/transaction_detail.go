package models

import (
	"time"

	"gorm.io/gorm"
)

type TransactionDetail struct {
	ID            uint      `gorm:"primaryKey;autoIncrement"`
	TransactionID uint      `gorm:"not null"`
	ProductID     uint      `gorm:"not null"`
	Quantity      uint      `gorm:"not null"`
	CreatedAt     time.Time `gorm:"not null;default:current_timestamp"`
	UpdatedAt     time.Time `gorm:"not null;default:current_timestamp"`
	DeletedAt     gorm.DeletedAt
}
