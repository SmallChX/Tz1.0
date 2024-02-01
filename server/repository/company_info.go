package repository

import (
	"jobfair2024/model"

	"gorm.io/gorm"
)

type CompanyInformationRepository interface {
	Create(companyInfo *model.CompanyInformation) error 
	Update(companyInfo *model.CompanyInformation) error
	Delete(id int64) error
	FindByID(id int64) (*model.CompanyInformation, error)
	FindAll() ([]model.CompanyInformation, error)
}

type companyInformationRepositoryImpl struct {
	db *gorm.DB
}

func NewCompanyInformationRepository(db *gorm.DB) CompanyInformationRepository{
	return &companyInformationRepositoryImpl{
		db: db,
	}
}

func (repo *companyInformationRepositoryImpl) Create(companyInfo *model.CompanyInformation) error {
	return repo.db.Create(&companyInfo).Error
}

func (repo *companyInformationRepositoryImpl) Update(companyInfo *model.CompanyInformation) error {
	return repo.db.Save(&companyInfo).Error
}

func (repo *companyInformationRepositoryImpl) Delete(id int64) error {
	return repo.db.Delete(&model.CompanyInformation{}, id).Error 
} 

func (repo *companyInformationRepositoryImpl) FindByID(id int64) (*model.CompanyInformation ,error) {
	var companyInfo model.CompanyInformation
	err := repo.db.First(&companyInfo, id).Error
	return &companyInfo, err
}

func (repo *companyInformationRepositoryImpl) FindAll() ([]model.CompanyInformation, error) {
	var companyInfos []model.CompanyInformation
	err := repo.db.Find(&companyInfos).Error
	return companyInfos, err
}