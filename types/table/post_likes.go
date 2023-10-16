package table

import "time"

type PostLikes struct {
	UserId          *uint64    `gorm:"primaryKey" json:"UserId"` //คนที่like
	OwnerPostLike   *Posts     `gorm:"foreignKey:OwnerPostLikeId"`
	OwnerPostLikeId *uint64    `gorm:"not null"` //postของเจ้าของที่โดนlike
	Owner           *Users     `gorm:"foreignKey:OwnerId"`
	OwnerId         *uint64    `gorm:"not null"` // userid ของpostที่โดนlike
	CreatedAt       *time.Time `gorm:"not null"` // Embedded field
}
