package table

import "time"

type PostDonate struct {
	UserId            *uint64    `gorm:"primaryKey" json:"userId"` // userid of donater มันเก็ยไงละ
	OwnerDonatePost   *Posts     `gorm:"foreignKey:OwnerDonatePostId"`
	OwnerDonatePostId *uint64    `gorm:"not null"` //postของเจ้าของที่โดนdonate
	Owner             *Users     `gorm:"foreignKey:OwnerId"`
	OwnerId           *uint64    `gorm:"not null"` // userid ของpostที่โดนdonate
	DonateAmount      *uint64    `gorm:"not null"`
	CreatedAt         *time.Time `gorm:"not null"` // Embedded field

}
