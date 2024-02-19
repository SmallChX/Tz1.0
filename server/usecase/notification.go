package usecase

import (
	"context"
	"encoding/base64"
	"fmt"
	"jobfair2024/model"
	"jobfair2024/repository"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

type NotificationUsecase interface {
	GetUserNotification(c *gin.Context, userInfo UserInfo) ([]NotificationInfo, error)
	CreateNotification(c *gin.Context, userInfo UserInfo, notificationInfo NotificationInfo) error
}

type notificationUsecaseImpl struct {
	notificationRepository repository.NotificationRepository
	accountUserRepository  repository.UserAccountRepository
}

func NewNotificationUsecase(
	notificationRepository repository.NotificationRepository,
	accountUserRepository repository.UserAccountRepository,
) NotificationUsecase {
	return &notificationUsecaseImpl{
		notificationRepository: notificationRepository,
		accountUserRepository:  accountUserRepository,
	}
}

type NotificationInfo struct {
	ID         int64
	Title      string
	Content    string    `gorm:"type:text"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	ReceiverID int64     `gorm:"index"`
	SenderID   *int64    `gorm:"index"`
	IsRead     bool      `gorm:"default:false"`
}

func (u *notificationUsecaseImpl) GetUserNotification(c *gin.Context, userInfo UserInfo) ([]NotificationInfo, error) {
	value, err := u.notificationRepository.FindByUserID(userInfo.ID)
	if err != nil {
		return nil, err
	}

	notificationList := make([]NotificationInfo, 0)
	for _, notification := range value {
		notificationList = append(notificationList, NotificationInfo{
			ID:         notification.ID,
			Title:      notification.Title,
			Content:    notification.Content,
			CreatedAt:  notification.CreatedAt,
			ReceiverID: notification.ReceiverID,
			SenderID:   &notification.ID,
			IsRead:     notification.IsRead,
		})
	}

	return notificationList, nil
}

func (u *notificationUsecaseImpl) CreateNotification(c *gin.Context, userInfo UserInfo, notificationInfo NotificationInfo) error {
	notification := model.Notification{
		Title:      notificationInfo.Title,
		Content:    notificationInfo.Content,
		ReceiverID: notificationInfo.ReceiverID,
		SenderID:   notificationInfo.SenderID,
	}

	switch userInfo.Role {
	case (model.Admin):
		notification.Type = model.SystemNotification
		break
	default:
		notification.Type = model.PersonalNotification
	}

	if err := u.notificationRepository.Create(notification); err != nil {
		return err
	}

	// receiver, err := u.accountUserRepository.FindByID(notification.ReceiverID)
	// if err != nil {
	// 	return err
	// }

	if err := u.sendEmail(
		"minhthuan55588@gamil.com",
		notification.Title,
		notification.Content,
	); err != nil {
		return err
	}

	return nil
}

func getClient() *http.Client {
	config := &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Endpoint:     google.Endpoint,
		Scopes:       []string{gmail.MailGoogleComScope},
	}

	// Tạo một token mới từ refresh token
	token := &oauth2.Token{RefreshToken: "1//04AR8aSFAsdLMCgYIARAAGAQSNwF-L9IrlV5Fp0iglq8iFeRfRg79N7og0KZxQbhtIF8f2SP4zYlC7IWMUt9ZmOEoYK22HSWePUw	"}

	return config.Client(context.Background(), token)
}

func (u *notificationUsecaseImpl) sendEmail(to, subject, body string) error {
	// ctx := context.Background()
	client := getClient()

	srv, err := gmail.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail Client %v", err)
	}

	var message gmail.Message

	// Lưu ý: Địa chỉ email người gửi và người nhận cần được đóng gói đúng cách.
	emailTo := "To: " + to
	emailFrom := "From: " + "your_email@gmail.com" // Đổi "your_email@gmail.com" thành địa chỉ Gmail của bạn
	emailSubject := "Subject: " + subject
	emailBody := "\r\n" + body
	msg := []byte(emailTo + "\r\n" + emailSubject + "\r\n" + emailBody)

	// Encode the email message to base64url
	message.Raw = base64.URLEncoding.EncodeToString(msg)

	// Send the email
	_, err = srv.Users.Messages.Send(emailFrom, &message).Do()
	if err != nil {
		log.Fatalf("Unable to send email: %v", err)
		return err
	}

	fmt.Println("Email sent successfully")
	return nil
}
