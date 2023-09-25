package model

import (
	"picme-backend/types/enum"
	"time"
)

type Posts struct {
	Id        *uint64        `gorm:"primaryKey"`
	ImageUrl  *string        `gorm:"type:TEXT; not null"`
	Owner     *User          `gorm:"foreignKey:OwnerId"`
	OwnerId   *uint64        `gorm:"not null"`
	Caption   *string        `gorm:"type:TEXT; not null"`
	Category  *enum.Category `gorm:"type:ENUM('painting', 'drawing'); not null"`
	CreatedAt *time.Time     `gorm:"not null"` // Embedded field
	UpdatedAt *time.Time     `gorm:"not null"` // Embedded field
}
