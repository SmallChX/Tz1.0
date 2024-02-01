package usecase

import (
	"errors"
	"jobfair2024/model"
	"jobfair2024/repository"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Xử lý về phân quyền và xác thực người dùng
type AuthenticationUsecase interface {
	Login(c *gin.Context, username string, password string) (*UserInfo, error)
	CreateAccount(c *gin.Context, username string, password string) error
	GetUserInfo(c *gin.Context, id int64) (*UserInfo, error)
}

type authenticationImpl struct {
	userAccountRepository repository.UserAccountRepository
}

func NewAuthenticationUsecase(
	userAccountRepository repository.UserAccountRepository,
) AuthenticationUsecase {
	return &authenticationImpl{
		userAccountRepository: userAccountRepository,
	}
}

type UserInfo struct {
	ID    int64
	Role  model.UserRole
	Email string
}

func (a *authenticationImpl) Login(c *gin.Context, username string, password string) (*UserInfo, error) {
	userAccount, err := a.userAccountRepository.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	if userAccount == nil {
		return nil, errors.New("no user found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userAccount.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	userInfo := &UserInfo{
		ID:    userAccount.ID,
		Role:  userAccount.Role,
		Email: userAccount.Email,
	}

	return userInfo, nil
}

func (a *authenticationImpl) CreateAccount(c *gin.Context, username string, password string) error {
	// userAccount, _ := a.userAccountRepository.FindByUsername(username)
	// if userAccount != nil {
	// 	return errors.New("already have user")
	// }

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	userAccount := &model.UserAccount{
		Username: username,
		Password: string(hashPassword),
		Role: model.Company,
	}

	err = a.userAccountRepository.Create(userAccount)
	if err != nil {
		return err
	}

	return nil
}

func (a *authenticationImpl) GetUserInfo(c *gin.Context, id int64) (*UserInfo, error) {
	userInfo, err := a.userAccountRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return &UserInfo{
		ID: userInfo.ID,
		Role: userInfo.Role,
		Email: userInfo.Email,
	}, nil
}