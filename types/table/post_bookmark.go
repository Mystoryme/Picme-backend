package table

import "time"

type PostBookMark struct {
	UserId              *uint64    `gorm:"primaryKey" json:"userId"` // user ที่เป็นคนกด bookmark
	OwnerBookMarkPost   *Posts     `gorm:"foreignKey:BookMarkPostId"`
	OwnerBookMarkPostId *uint64    `gorm:"not null"` //id ของpost ที่ถูก bookmark
	Owner               *Users     `gorm:"foreignKey:OwnerId"`
	OwnerId             *uint64    `gorm:"not null"` // userid ของpostที่โดนbookmark
	CreatedAt           *time.Time `gorm:"not null"` // Embedded field

}
