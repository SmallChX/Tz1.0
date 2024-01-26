package model

type Booth struct {
	ID        int64
	CompanyID int64          // Company own
	Requests  []BoothRequest `gorm:"many2many:request_booths;"`
	// Size
	Length int
	Width  int
}
