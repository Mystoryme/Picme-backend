package table

import "time"

type PostComments struct {
	UserId             *uint64    `gorm:"primaryKey" json:"userId"` // userId ของคนpost
	OwnerCommentPost   *Posts     `gorm:"foreignKey:OwnerCommentPostId"`
	OwnerCommentPostId *uint64    `gorm:"not null"` //postของเจ้าของทีโดน comment
	Owner              *Users     `gorm:"foreignKey:OwnerId"`
	OwnerId            *uint64    `gorm:"not null"` // userid ของpostที่โดนcomment
	Message            *string    `gorm:"TEXT; not null"`
	CreatedAt          *time.Time `gorm:"not null"` // Embedded field
	UpdatedAt          *time.Time `gorm:"not null"` // Embedded field
}
