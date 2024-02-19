package usecase

import (
	"jobfair2024/repository"

	"github.com/gin-gonic/gin"
)

// Quyền: admin, company.
// Đối với Admin: quản lý tất cả company đăng ký tham gia JobFair và tạo account của Company
// Policy: Gửi Proposal kèm thông tin tài khoản và phương thức đăng ký cho Company.
// Đối với các công ty không nhận được proposal, sẽ có phương thức liên lạc với admin để nhận account.
type CompanyUsecase interface {
	GetCompanyBoothIDs(c *gin.Context, userInfo *UserInfo) ([]int64, error)
	GetCompanyInfo(c *gin.Context, userInfo *UserInfo) (*CompanyInfo, error)
}

type companyUsecaseImpl struct {
	companyInfoRepository repository.CompanyInformationRepository
}

func NewCompanyUsecase(
	companyInfoRepository repository.CompanyInformationRepository,
) CompanyUsecase {
	return &companyUsecaseImpl{
		companyInfoRepository: companyInfoRepository,
	}
}

func (u *companyUsecaseImpl) GetCompanyBoothIDs(c *gin.Context, userInfo *UserInfo) ([]int64, error) {
	if err := validateCompanyRole(userInfo); err != nil {
		return nil, err
	}
	companyInfo, err := u.companyInfoRepository.FindByUserID(userInfo.ID)
	if err != nil {
		return nil, err
	}

	boothIDs := getBoothIDs(companyInfo.Booths)
	return boothIDs, nil
}

type CompanyInfo struct {
	ID int64
	CompanyName string
	BoothIDs    []int64
	// Contact Represent = HR
	RepresentName        string `json:"represent_name"`
	RepresentPhoneNumber string `json:"represent_phone_number"`
	RepresentMail        string `json:"represent_mail"`
}

func (u *companyUsecaseImpl) GetCompanyInfo(c *gin.Context, userInfo *UserInfo) (*CompanyInfo, error) {
	if err := validateCompanyRole(userInfo); err != nil {
		return nil, err
	}

	companyInfoDB, err := u.companyInfoRepository.FindByUserID(userInfo.ID)
	if err != nil {
		return nil, err
	}

	return &CompanyInfo{
		CompanyName:          companyInfoDB.CompanyName,
		RepresentName:        companyInfoDB.RepresentName,
		RepresentMail:        companyInfoDB.RepresentMail,
		RepresentPhoneNumber: companyInfoDB.RepresentPhoneNumber,
		BoothIDs:             getBoothIDs(companyInfoDB.Booths),
	}, nil

}
