package repository

import (
	"jobfair2024/model"

	"gorm.io/gorm"
)

type StudentRepository interface {
	FindByID(id int64) (model.StudentInformation, error)
}

type studentRepositoryImpl struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepositoryImpl{
		db: db,
	}
}

func (repo *studentRepositoryImpl) FindByID(id int64) (model.StudentInformation, error) {
	var studentInfo model.StudentInformation
	err := repo.db.First(&studentInfo, id).Error
	return studentInfo, err
}

