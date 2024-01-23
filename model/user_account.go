package model

import (
	"jobfair2024/setting"

	"gorm.io/gorm"
)

type UserRole string

const (
	Company UserRole = "company"
	Admin   UserRole = "admin"
	Student UserRole = "student"
)

type UserAccount struct {
	ID       int64
	Account  string
	Password string
	Role     UserRole
	Email    string
}

type UserAccountRepository interface {
	Create(user *UserAccount) error
	Update(user *UserAccount) error
	Delete(id int64) error
	FindByID(id int64) (*UserAccount, error)
	FindAll() ([]UserAccount, error)
}

type userAccountRepositoryImpl struct {
	db *gorm.DB
}

func NewUserAccountRepository(db *gorm.DB) UserAccountRepository {
	return &userAccountRepositoryImpl{
		db: db,
	}
}

func (repo *userAccountRepositoryImpl) GetDB() *gorm.DB {
	return setting.GetDB()
}

// Create adds a new UserAccount to the database
func (repo *userAccountRepositoryImpl) Create(user *UserAccount) error {
	return repo.db.Create(user).Error
}

// Update modifies an existing UserAccount in the database
func (repo *userAccountRepositoryImpl) Update(user *UserAccount) error {
	return repo.db.Save(user).Error
}

// Delete removes a UserAccount from the database
func (repo *userAccountRepositoryImpl) Delete(id int64) error {
	return repo.db.Delete(&UserAccount{}, id).Error
}

// FindByID finds a UserAccount by its ID
func (repo *userAccountRepositoryImpl) FindByID(id int64) (*UserAccount, error) {
	var user UserAccount
	err := repo.db.First(&user, id).Error
	return &user, err
}

// FindAll returns all UserAccounts in the database
func (repo *userAccountRepositoryImpl) FindAll() ([]UserAccount, error) {
	var users []UserAccount
	err := repo.db.Find(&users).Error
	return users, err
}
