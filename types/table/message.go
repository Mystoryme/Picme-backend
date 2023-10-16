package table

import "time"

type Message struct {
	MessageId  *uint64    `gorm:"primaryKey" json:"userId"`
	Receiver   *Users     `gorm:"foreignKey:ReceiverId"`
	ReceiverId *uint64    `gorm:"not null"`
	Message    *string    `gorm:"TEXT; not null"`
	CreatedAt  *time.Time `gorm:"not null"`
	// Embedded field

}
