package model

type UserRole string

const (
	Company UserRole = "company"
	Admin   UserRole = "admin"
	Student UserRole = "student"
)

type UserAccount struct {
	ID       int64 `gorm:"primaryKey"`
	Username  string
	Password string
	Role     UserRole
	Email    string
}
