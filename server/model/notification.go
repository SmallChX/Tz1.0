package model

import (
	"time"
)

type NotificationType string

const (
	JobNotification      NotificationType = "job_notification"
	SystemNotification   NotificationType = "system_notification"
	PersonalNotification NotificationType = "personal_notification"
)

type Notification struct {
	ID         int64            `gorm:"primaryKey"`
	Type       NotificationType `gorm:"type:varchar(100)"`
	Title string `gorm:"type:text"`
	Content    string           `gorm:"type:text"`
	CreatedAt  time.Time        `gorm:"default:CURRENT_TIMESTAMP"`
	ReceiverID int64            `gorm:"index"`
	SenderID   *int64           `gorm:"index"`
	IsRead     bool             `gorm:"default:false"`
}
