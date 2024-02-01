package model

import "time"

type Booth struct {
	ID            int64          `gorm:"primaryKey"`
	CompanyID     int64          
	Requests      []BoothRequest `gorm:"many2many:request_booths;"`
	ChangeRequest []BoothRequest `gorm:"many2many:des_request_booths;"`
	CreateAt      time.Time
	// Size
	Length int
	Width  int
}
