package model

type Booth struct {
	ID        int64 `gorm:"primaryKey"`
	CompanyID int64          // Company own
	Requests  []BoothRequest `gorm:"many2many:request_booths;"`
	ChangeRequest []BoothRequest `gorm:"many2many:des_request_booths;"`
	// Size
	Length int
	Width  int
}
