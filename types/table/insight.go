package table

import (
	"picme-backend/types/enum"
	"time"
)

type Insight struct {
	Id          *uint64 `gorm:"primaryKey"`
	Trigger     *User   `gorm:"foreignKey:TriggerId"`
	TriggerId   *uint64
	Triggee     *User `gorm:"foreignKey:TriggeeId"`
	TriggeeId   *uint64
	InsightType *enum.Insight `gorm:"type:ENUM('view','like','search'); not null"`
	CreatedAt   *time.Time    `gorm:"not null"` // Embedded field
	UpdatedAt   *time.Time    `gorm:"not null"` // Embedded field
}
