package table

import "time"

type PostBoost struct {
	Id            *uint64    `gorm:"primaryKey" `
	Paid          *bool      `gorm:"not null;default:false"`
	Amount        *uint64    `gorm:"not null;default:0"`
	Post          *Post      `gorm:"foreignKey:PostId"`
	PostId        *uint64    `gorm:"not null"` //id ของpost ที่ถูก boost
	BoostEnd      *time.Time `gorm:"not null"` // Embedded field
	TransactionId *string    `gorm:"not null"`
}
