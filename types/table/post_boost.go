package table

import "time"

type PostBoost struct {
	BoostId     *uint64    `gorm:"primaryKey" json:"boostId"`
	BoostPost   *Posts     `gorm:"foreignKey:BoostPostId"`
	BoostPostId *uint64    `gorm:"not null"` //id ของpost ที่ถูก boost
	BoostEnd    *time.Time `gorm:"not null"` // Embedded field

}
