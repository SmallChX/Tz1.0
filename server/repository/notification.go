package repository

import (
	"jobfair2024/model"

	"gorm.io/gorm"
)

type NotificationRepository interface {
	FindByID(id int64) (*model.Notification, error)
	FindByUserID(userID int64) ([]model.Notification, error)
	FindAll() ([]model.Notification, error)
	Create(notification model.Notification) error
	Delete(id int64) error
}

type notificationRepositoryImpl struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &notificationRepositoryImpl{
		db: db,
	}
}

// FindByID finds a notification by its ID
func (repo *notificationRepositoryImpl) FindByID(id int64) (*model.Notification, error) {
	var notification model.Notification
	err := repo.db.First(&notification, id).Error
	if err != nil {
		return nil, err
	}
	return &notification, nil
}

// FindByUserID finds all notifications for a specific user
func (repo *notificationRepositoryImpl) FindByUserID(userID int64) ([]model.Notification, error) {
	var notifications []model.Notification
	err := repo.db.Where("receiver_id = ?", userID).Find(&notifications).Error
	if err != nil {
		return nil, err
	}
	return notifications, nil
}

// FindAll retrieves all notifications from the database
func (repo *notificationRepositoryImpl) FindAll() ([]model.Notification, error) {
	var notifications []model.Notification
	err := repo.db.Find(&notifications).Error
	if err != nil {
		return nil, err
	}
	return notifications, nil
}

// Create adds a new notification to the database
func (repo *notificationRepositoryImpl) Create(notification model.Notification) error {
	return repo.db.Create(&notification).Error
}

// Delete removes a notification from the database by its ID
func (repo *notificationRepositoryImpl) Delete(id int64) error {
	return repo.db.Delete(&model.Notification{}, id).Error
}
