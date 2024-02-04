package repository

import (
	"jobfair2024/model"

	"gorm.io/gorm"
)

type AdminRepository interface {
	FindByID(id int64) (model.AdminInformation, error)
}

type adminRepositoryImpl struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepositoryImpl{
		db: db,
	}
}

func (repo *adminRepositoryImpl) FindByID(id int64) (model.AdminInformation, error) {
	var adminInfo model.AdminInformation
	err := repo.db.First(&adminInfo, id).Error
	return adminInfo, err
}

