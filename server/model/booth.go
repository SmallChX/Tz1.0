package model

type Booth struct {
	ID            int64 `gorm:"primaryKey"`
	CompanyID     *int64
	Level         int            // Level 1 dành cho doanh nghiệp thường, 2 ản cho nhà tài trợ
	Requests      []BoothRequest `gorm:"many2many:request_booths;"`
	ChangeRequest []BoothRequest `gorm:"many2many:des_request_booths;"`
	Price         int
}
