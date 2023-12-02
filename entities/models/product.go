package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	ID         uint            `gorm:"primaryKey;autoIncrement"`
	Name       string          `gorm:"not null"`
	Price      decimal.Decimal `gorm:"not null;type:decimal(10, 2)"`
	Stock      uint            `gorm:"not null"`
	CategoryID uint            `gorm:"not null"`
	Category   Category        `gorm:"foreignKey:CategoryID;references:ID"`
	CreatedAt  time.Time       `gorm:"not null;default:current_timestamp"`
	UpdatedAt  time.Time       `gorm:"not null;default:current_timestamp"`
	DeletedAt  gorm.DeletedAt
}
