package repository

import (
	"jobfair2024/model"
	"jobfair2024/setting"

	"gorm.io/gorm"
)

type UserAccountRepository interface {
	Create(user *model.UserAccount) error
	Update(user *model.UserAccount) error
	Delete(id int64) error
	FindByID(id int64) (*model.UserAccount, error)
	FindAll() ([]model.UserAccount, error)
	FindByUsername(username string) (*model.UserAccount, error)
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
func (repo *userAccountRepositoryImpl) Create(user *model.UserAccount) error {
	return repo.db.Create(user).Error
}

// Update modifies an existing UserAccount in the database
func (repo *userAccountRepositoryImpl) Update(user *model.UserAccount) error {
	return repo.db.Save(user).Error
}

// Delete removes a UserAccount from the database
func (repo *userAccountRepositoryImpl) Delete(id int64) error {
	user, err := repo.FindByID(id)
	if err != nil {
		return err
	}
	return repo.db.Select("CompanyInfo", "AdminInfo", "StudentInfo").Delete(&user).Error
}

// FindByID finds a UserAccount by its ID
func (repo *userAccountRepositoryImpl) FindByID(id int64) (*model.UserAccount, error) {
	var user model.UserAccount
	err := repo.db.First(&user, id).Error
	return &user, err
}

// FindAll returns all UserAccounts in the database
func (repo *userAccountRepositoryImpl) FindAll() ([]model.UserAccount, error) {
	var users []model.UserAccount
	err := repo.db.Preload("CompanyInfo").Preload("StudentInfo").Preload("AdminInfo").Find(&users).Error
	return users, err
}

func (repo *userAccountRepositoryImpl) FindByUsername(username string) (*model.UserAccount, error) {
	var user model.UserAccount
	err := repo.db.Where("username = ?", username).First(&user).Error
	return &user, err
}
