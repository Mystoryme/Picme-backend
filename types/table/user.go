package table

import "time"

type User struct {
	Id             *uint64    `gorm:"primaryKey"`
	Username       *string    `gorm:"index:username,unique; not null"`
	HashedPassword *string    `gorm:"not null"`
	CreatedAt      *time.Time `gorm:"not null"`
	UpdatedAt      *time.Time `gorm:"not null"`
	Email          *string    `gorm:"index:email,unique; not null"`
}
