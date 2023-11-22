package table

import "time"

type PostDonate struct {
	User         *User      `gorm:"foreignKey:UserId"`
	UserId       *uint64    `gorm:"not null"` //คนที่donate
	Post         *Post      `gorm:"foreignKey:PostId"`
	PostId       *uint64    `gorm:"not null"` //postที่โดนdonate
	DonateAmount *uint64    `gorm:"not null"`
	Paid         *bool      `gorm:"not null"`
	CreatedAt    *time.Time `gorm:"not null"` // Embedded field

}
