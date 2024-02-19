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
}

type authenticationImpl struct {
	userAccountRepository repository.UserAccountRepository
	companyInfoRepository repository.CompanyInformationRepository
	adminRepository       repository.AdminRepository
	studentRepository     repository.StudentRepository
}

func NewAuthenticationUsecase(
	userAccountRepository repository.UserAccountRepository,
) AuthenticationUsecase {
	return &authenticationImpl{
		userAccountRepository: userAccountRepository,
	}
}

type UserInfo struct {
	ID         int64          `json:"user_id"`
	Role       model.UserRole `json:"role"`
	Email      *string        `json:"mail"`
	UserName   string         `json:"username"`
	Name       string         `json:"name"`
	FirstLogin bool           `json:"first_login"`
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
		FirstLogin: userAccount.FirstLogin,
	}

	return userInfo, nil
}
