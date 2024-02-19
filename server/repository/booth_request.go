package repository

import (
	"jobfair2024/model"

	"gorm.io/gorm"
)

type BoothRequestRepository interface {
	Create(boothRequest *model.BoothRequest) error
	Update(boothRequest *model.BoothRequest) error
	Delete(id int64) error
	FindByID(id int64) (*model.BoothRequest, error)
	FindAll() ([]model.BoothRequest, error)
	GetCompanyBoothRequests(companyID int64) ([]model.BoothRequest, error)
}

type boothRequestRepositoryImpl struct {
	db *gorm.DB
}

// NewBoothRepository creates a new instance of BoothRepository
func NewBoothRequestRepository(db *gorm.DB) BoothRequestRepository {
	return &boothRequestRepositoryImpl{
		db: db,
	}
}

// Create adds a new Booth to the database
func (repo *boothRequestRepositoryImpl) Create(boothRequest *model.BoothRequest) error {
	return repo.db.Create(boothRequest).Error
}

// Update modifies an existing Booth in the database
func (repo *boothRequestRepositoryImpl) Update(boothRequest *model.BoothRequest) error {
	return repo.db.Save(boothRequest).Error
}

// Delete removes a Booth from the database
func (repo *boothRequestRepositoryImpl) Delete(id int64) error {
	request, _ := repo.FindByID(id)
	repo.db.Model(&request).Select("Booths", "DestinationBooths").Delete(request)
	return repo.db.Delete(&model.BoothRequest{}, id).Error
}

// FindByID finds a Booth by its ID
func (repo *boothRequestRepositoryImpl) FindByID(id int64) (*model.BoothRequest, error) {
	var request model.BoothRequest
	err := repo.db.Preload("Booths").Preload("DestinationBooths").First(&request, id).Error

	if err != nil {
		return nil, err
	}

	return &request, err
}

// FindAll returns all Booths in the database
func (repo *boothRequestRepositoryImpl) FindAll() ([]model.BoothRequest, error) {
	var requests []model.BoothRequest
	err := repo.db.Preload("Booths").Preload("DestinationBooths").Find(&requests).Error

	if err != nil {
		return nil, err
	}

	return requests, err
}

func (repo *boothRequestRepositoryImpl) GetCompanyBoothRequests(companyID int64) ([]model.BoothRequest, error) {
    var boothRequests []model.BoothRequest
    err := repo.db.Preload("Booths").Preload("DestinationBooths").Where("company_id = ?", companyID).Find(&boothRequests).Error
    if err != nil {
        return nil, err
    }
    return boothRequests, nil
}

