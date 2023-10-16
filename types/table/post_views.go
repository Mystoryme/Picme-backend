package table

import "time"

type PostViews struct {
	Id          *uint64    `gorm:"primaryKey"`
	OwnerPost   *Posts     `gorm:"foreignKey:OwnerPostId"`
	OwnerPostId *uint64    `gorm:"not null"` // postที่โดนดู
	Owner       *Users     `gorm:"foreignKey:OwnerId"`
	OwnerId     *uint64    `gorm:"not null"` // owner ของpost ที่โดนดู
	CreatedAt   *time.Time `gorm:"not null"` // Embedded field
}
