package model

import "time"

type StudentInformation struct {
	UserID      int64
	StudentID   int64
	FirstName   string
	LastName    string
	DateofBirth time.Time
	PhoneNumber string
	Department  string
	Major       string

	UserAccountID int64 `gorm:"primaryKey"`
	UserAccount `gorm:"foreignKey:UserAccountID"`
}
