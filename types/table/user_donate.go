package table

import "time"

type UserDonate struct {
	Donor         *User      `gorm:"foreignKey:DonorId"`
	DonorId       *uint64    `gorm:"not null"` //คนที่donate
	Paid          *bool      `gorm:"not null"`
	User          *User      `gorm:"foreignKey:UserId"`
	UserId        *uint64    `gorm:"not null"` //postที่โดนdonate
	DonateAmount  *uint64    `gorm:"not null"`
	TransactionId *string    `gorm:"not null"`
	CreatedAt     *time.Time `gorm:"not null"` // Embedded field

}
