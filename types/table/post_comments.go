package table

import "time"

type PostComment struct {
	User      *User      `gorm:"foreignKey:UserId"`
	UserId    *uint64    `gorm:"not null"` //คนที่like
	Post      *Post      `gorm:"foreignKey:PostId"`
	PostId    *uint64    `gorm:"not null"` //postของเจ้าของทีโดน comment
	Message   *string    `gorm:"TEXT; not null"`
	CreatedAt *time.Time `gorm:"not null"` // Embedded field
	UpdatedAt *time.Time `gorm:"not null"` // Embedded field
}
