package table

import (
	"picme-backend/types/enum"
	"time"
)

type Notification struct {
	Id               *uint64 `gorm:"primaryKey"`
	Trigger          *User   `gorm:"foreignKey:TriggerId"`
	TriggerId        *uint64
	Triggee          *User `gorm:"foreignKey:TriggeeId"`
	TriggeeId        *uint64
	Post             *Post `gorm:"foreignKey:PostId"`
	PostId           *uint64
	NotificationType *enum.Notification `gorm:"type:ENUM('comment','like','user_donate','post_donate'); not null"`
	CreatedAt        *time.Time         `gorm:"not null"` // Embedded field
	UpdatedAt        *time.Time         `gorm:"not null"` // Embedded field
}
