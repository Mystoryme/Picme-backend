package table

import "time"

type PostLike struct {
	User      *User      `gorm:"foreignKey:UserId"`
	UserId    *uint64    `gorm:"not null"` //คนที่like
	Post      *Post      `gorm:"foreignKey:PostId"`
	PostId    *uint64    `gorm:"not null"` //post ที่ถูกไลค์
	CreatedAt *time.Time `gorm:"not null"` // Embedded field
}
