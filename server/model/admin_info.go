package model

type AdminInformation struct {
	AdminId int64 `gorm:"primaryKey"`
	Name string	
	UserAccountID int64 `gorm:"primaryKey"`
	UserAccount UserAccount `gorm:"foreignKey:UserAccountID"`
}
