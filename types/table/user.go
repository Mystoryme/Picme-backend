package table

import "time"

type User struct {
	Id        *uint64    `gorm:"primaryKey"`
	Username  *string    `gorm:"index:username,unique; not null"`
	Email     *string    `gorm:"type:VARCHAR(60); not null"`
	Password  *string    `gorm:"type:TEXT; not null"`
	Bio       *string    `gorm:"type:VARCHAR(160); null"`
	Contact   *string    `gorm:"type:VARCHAR(160); null"`
	AvatarUrl *string    `gorm:"type:TEXT; null"`
	CreatedAt *time.Time `gorm:"not null"` // Embedded field
	UpdatedAt *time.Time `gorm:"not null"` // Embedded field
}
