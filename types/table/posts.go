package table

import (
	"picme-backend/types/enum"
	"time"
)

type Posts struct {
	Id          *uint64           `gorm:"primaryKey"`
	ImageUrl    *string           `gorm:"type:TEXT; not null"`
	Owner       *Users            `gorm:"foreignKey:OwnerId"`
	OwnerId     *uint64           `gorm:"not null"`
	Caption     *string           `gorm:"type:TEXT; not null"`
	Category    *enum.Category    `gorm:"type:ENUM('painting', 'drawing_&_illustration','mixed_media_&_collage','digital_/_graphic','other'); not null"`
	Application *enum.Application `gorm:"type:ENUM('procreate', 'ibis_paintX','clip_studio_paint','blender','photoshop','other'); not null"`
	CreatedAt   *time.Time        `gorm:"not null"` // Embedded field
	UpdatedAt   *time.Time        `gorm:"not null"` // Embedded field
}
