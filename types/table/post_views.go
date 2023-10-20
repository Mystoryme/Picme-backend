package table

import "time"

type PostViews struct {
	User      *User      `gorm:"foreignKey:UserId"`
	UserId    *uint64    `gorm:"not null"` //คนที่ดู
	Post      *Post      `gorm:"foreignKey:PostId"`
	PostId    *uint64    `gorm:"not null"` // postที่โดนดู
	CreatedAt *time.Time `gorm:"not null"` // Embedded field
}
