package model

import "time"

type PostComments struct {
	UserId    *uint64    `gorm:"primaryKey" json:"userId"`
	PostId    *uint64    `gorm:"primaryKey" json:"postId"`
	Message   *string    `gorm:"TEXT; not null"`
	CreatedAt *time.Time `gorm:"not null"` // Embedded field
	UpdatedAt *time.Time `gorm:"not null"` // Embedded field
}
