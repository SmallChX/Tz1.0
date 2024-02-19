package usecase

import (
	"fmt"
	"jobfair2024/model"
	"jobfair2024/repository"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserAccountUsecase interface {
	CreateAccount(c *gin.Context, username string, password string, role model.UserRole, name string, userInfo *UserInfo) error
	GetUserInfo(c *gin.Context, id int64, role model.UserRole, userInfo *UserInfo) (interface{}, error)
	GetAllUserInfo(c *gin.Context, userInfo *UserInfo) ([]*UserInfo, error)
	ResetPassword(c *gin.Context, id int64, userInfo *UserInfo) error
	DeleteAccount(c *gin.Context, id int64, userInfo *UserInfo) error
	UpdateUserAccountInfo(c *gin.Context, userInfo *UserInfo, companyInfo CompanyUpdateInfo) error
}

type userAccountUsecaseImpl struct {
	userAccountInfoRepository repository.UserAccountRepository
	companyInfoRepository     repository.CompanyInformationRepository
	adminRepository           repository.AdminRepository
	studentRepository         repository.StudentRepository
}

func NewUserAccountUsecase(
	userAccountInfoRepository repository.UserAccountRepository,
	companyInfoRepository repository.CompanyInformationRepository,
	adminRepository repository.AdminRepository,
	studentRepository repository.StudentRepository,
) UserAccountUsecase {
	return &userAccountUsecaseImpl{
		userAccountInfoRepository: userAccountInfoRepository,
		companyInfoRepository:     companyInfoRepository,
		adminRepository:           adminRepository,
		studentRepository:         studentRepository,
	}
}

func (u *userAccountUsecaseImpl) CreateAccount(c *gin.Context, username string, password string, role model.UserRole, name string, userInfo *UserInfo) error {
	if err := validateAdminRole(userInfo); err != nil {
		return err
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	userAccount := &model.UserAccount{
		Username: username,
		Password: string(hashPassword),
		Role:     role,
		Email:    nil,
		FirstLogin: true,
		CompanyInfo: model.CompanyInformation{
			CompanyName: name,
		},
	}

	err = u.userAccountInfoRepository.Create(userAccount)
	if err != nil {
		return err
	}

	return nil
}

func (u *userAccountUsecaseImpl) GetUserInfo(c *gin.Context, id int64, role model.UserRole, userInfo *UserInfo) (interface{}, error) {
	switch role {
	case model.Admin:
		userInfo, err := u.adminRepository.FindByID(id)
		if err != nil {
			return nil, err
		}
		return userInfo, nil
	case model.Company:
		companyInfo, err := u.companyInfoRepository.FindByID(id)
		if err != nil {
			return nil, err
		}
		return companyInfo, nil
	case model.Student:
		studentInfo, err := u.studentRepository.FindByID(id)
		if err != nil {
			return nil, err
		}
		return studentInfo, nil

	}

	return nil, nil
}

func (u *userAccountUsecaseImpl) GetAllUserInfo(c *gin.Context, userInfo *UserInfo) ([]*UserInfo, error) {
	if err := validateAdminRole(userInfo); err != nil {
		return nil, err
	}

	value, err := u.userAccountInfoRepository.FindAll()
	if err != nil {
		return nil, err
	}

	userList := make([]*UserInfo, 0)
	for _, user := range value {
		var name string
		switch user.Role {
		case model.Admin:
			name = user.AdminInfo.Name
		case model.Company:
			name = user.CompanyInfo.CompanyName
		case model.Student:
			name = user.StudentInfo.FirstName + " " + user.StudentInfo.LastName
		default:
			name = ""
		}

		userList = append(userList, &UserInfo{
			ID:       user.ID,
			Role:     user.Role,
			Email:    user.Email,
			UserName: user.Username,
			Name:     name,
		})
	}

	return userList, nil
}

func (u *userAccountUsecaseImpl) ResetPassword(c *gin.Context, id int64, userInfo *UserInfo) error {
	if err := validateAdminRole(userInfo); err != nil {
		return err
	}

	userAccount, err := u.userAccountInfoRepository.FindByID(id)
	if err != nil {
		return err
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	userAccount.Password = string(hashPassword)
	if err := u.userAccountInfoRepository.Update(userAccount); err != nil {
		return err
	}

	return nil
}

func (u *userAccountUsecaseImpl) DeleteAccount(c *gin.Context, id int64, userInfo *UserInfo) error {
	if err := validateAdminRole(userInfo); err != nil {
		return err
	}

	if err := u.userAccountInfoRepository.Delete(id); err != nil {
		return err
	}

	return nil
}

func (u *userAccountUsecaseImpl) UpdateUserAccountInfo(c *gin.Context, userInfo *UserInfo, companyInfo CompanyUpdateInfo) error {
	userAccountInfo, err := u.userAccountInfoRepository.FindByID(userInfo.ID)
	if err != nil {
		return err
	}
	fmt.Print("aaaa")
	switch userInfo.Role {
	case model.Company:
		err = updateCompanyInfo(u, userInfo, companyInfo)
		if err != nil {
			return err
		}
	}
	userAccountInfo.FirstLogin = false
	fmt.Print("bbbb")
	err = u.userAccountInfoRepository.Update(userAccountInfo)
	if err != nil {
		return err
	}

	return nil
}

type CompanyUpdateInfo struct {
	RepresentName        string
	RepresentPhoneNumber string
	RepresentMail        string
}

func updateCompanyInfo(u *userAccountUsecaseImpl, userInfo *UserInfo, updateInfo CompanyUpdateInfo) error {
	company, err := u.companyInfoRepository.FindByUserID(userInfo.ID)
	if err != nil {
		return err
	}

	company.RepresentName = updateInfo.RepresentName
	company.RepresentPhoneNumber = updateInfo.RepresentPhoneNumber
	company.RepresentMail = updateInfo.RepresentMail

	err = u.companyInfoRepository.Update(company)
	if err != nil {
		return err
	}

	return nil
}
