package model

import "gorm.io/gorm"

type Booth struct {
	ID        int64
	CompanyID int64          // Company own
	Requests  []BoothRequest `gorm:"many2many:request_booths;"`
	// Size
	Length int
	Width  int
}

type BoothRepository interface {
	Create(booth *Booth) error
	Update(booth *Booth) error
	Delete(id int64) error
	FindByID(id int64) (*Booth, error)
	FindAll() ([]Booth, error)
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
func (repo *boothRepositoryImpl) Create(booth *Booth) error {
	return repo.db.Create(booth).Error
}

// Update modifies an existing Booth in the database
func (repo *boothRepositoryImpl) Update(booth *Booth) error {
	return repo.db.Save(booth).Error
}

// Delete removes a Booth from the database
func (repo *boothRepositoryImpl) Delete(id int64) error {
	return repo.db.Delete(&Booth{}, id).Error
}

// FindByID finds a Booth by its ID
func (repo *boothRepositoryImpl) FindByID(id int64) (*Booth, error) {
	var booth Booth
	err := repo.db.First(&booth, id).Error
	return &booth, err
}

// FindAll returns all Booths in the database
func (repo *boothRepositoryImpl) FindAll() ([]Booth, error) {
	var booths []Booth
	err := repo.db.Find(&booths).Error
	return booths, err
}
