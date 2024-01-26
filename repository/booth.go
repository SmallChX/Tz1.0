package repository

import (
	"jobfair2024/model"

	"gorm.io/gorm"
)

type BoothRepository interface {
	Create(booth *model.Booth) error
	Update(booth *model.Booth) error
	Delete(id int64) error
	FindByID(id int64) (*model.Booth, error)
	FindAll() ([]model.Booth, error)
	FindByIds(ids []int64) ([]model.Booth, error) 
}

type boothRepositoryImpl struct {
	db *gorm.DB
}

// NewBoothRepository creates a new instance of BoothRepository
func NewBoothRepository(db *gorm.DB) BoothRepository {
	return &boothRepositoryImpl{
		db: db,
	}
}

// Create adds a new Booth to the database
func (repo *boothRepositoryImpl) Create(booth *model.Booth) error {
	return repo.db.Create(booth).Error
}

// Update modifies an existing Booth in the database
func (repo *boothRepositoryImpl) Update(booth *model.Booth) error {
	return repo.db.Save(booth).Error
}

// Delete removes a Booth from the database
func (repo *boothRepositoryImpl) Delete(id int64) error {
	return repo.db.Delete(&model.Booth{}, id).Error
}

// FindByID finds a Booth by its ID
func (repo *boothRepositoryImpl) FindByID(id int64) (*model.Booth, error) {
	var booth model.Booth
	err := repo.db.First(&booth, id).Error
	return &booth, err
}

// FindAll returns all Booths in the database
func (repo *boothRepositoryImpl) FindAll() ([]model.Booth, error) {
	var booths []model.Booth
	err := repo.db.Find(&booths).Error
	return booths, err
}

// FindBoothsByIds nhận vào một slice của int64 và trả về một slice của Booth và error
func (repo *boothRepositoryImpl) FindByIds(ids []int64) ([]model.Booth, error) {
    var booths []model.Booth

    // Truy vấn cơ sở dữ liệu sử dụng GORM
    result := repo.db.Where("id IN ?", ids).Find(&booths)
    if result.Error != nil {
        return nil, result.Error
    }

    return booths, nil
}