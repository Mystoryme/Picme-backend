package model

import "time"

type PostLikes struct {
	UserId    *uint64    `gorm:"primaryKey" json:"userId"`
	PostId    *uint64    `gorm:"primaryKey" json:"postId"`
	CreatedAt *time.Time `gorm:"not null"` // Embedded field
}
