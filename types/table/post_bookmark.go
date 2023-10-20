package table

import "time"

type PostBookMark struct {
	User      *User      `gorm:"foreignKey:UserId"`
	UserId    *uint64    `gorm:"not null"` //คนที่bookmark
	Post      *Post      `gorm:"foreignKey:PostId"`
	PostId    *uint64    `gorm:"not null"` //postของเจ้าของทีโดน bookmark
	CreatedAt *time.Time `gorm:"not null"` // Embedded field

}
