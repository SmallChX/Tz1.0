package usecase

import (
	"jobfair2024/repository"

	"github.com/gin-gonic/gin"
)

// Lấy từ model Booth. Sau khi Request được accept, cập nhật Company của Booth.
// Mỗi Booth chỉ được có một Company.
// Quyền xử lý: admin, company.
type BoothUsecase interface {
	GetBooth(c *gin.Context, userInfo *UserInfo, boothID int64) (*BoothInfo, error)
	GetAllBooths(c *gin.Context, userInfo *UserInfo) ([]*BoothInfo, error)
	UpdateBooths(c *gin.Context, userInfo *UserInfo, boothInfoList []BoothInfo) error
	UpdateBooth(c *gin.Context, userInfo *UserInfo, boothInfo BoothInfo) error 
	GetAllBoothCompany(c *gin.Context, userInfo *UserInfo) ([]*BoothCompany, error)
}

type boothImpl struct {
	boothRepository   repository.BoothRepository
	companyRepository repository.CompanyInformationRepository
}

func NewBoothUsecase(
	boothRepository repository.BoothRepository,
	companyRepository repository.CompanyInformationRepository,
) BoothUsecase {
	return &boothImpl{
		boothRepository:   boothRepository,
		companyRepository: companyRepository,
	}
}

type BoothInfo struct {
	ID      int64        `json:"ID"`
	Company BoothCompany `json:"company_info"`
	Level   int          `json:"level"` // Level 1 dành cho doanh nghiệp thường, 2 ản cho nhà tài trợ
	Price   int          `json:"price"`
}

type BoothCompany struct {
	ID   int64  `json:"company_id"`
	Name string `json:"name"`
}

func (b *boothImpl) GetAllBoothCompany(c *gin.Context, userInfo *UserInfo) ([]*BoothCompany, error) {
	// if err := validateAdminRole(userInfo); err != nil {
	// 	return nil, err
	// }

	companyList, err := b.companyRepository.FindAll()
	if err != nil {
		return nil, err
	}

	boothCompanyList := make([]*BoothCompany, 0)
	for _, company := range companyList {
		boothCompanyList = append(boothCompanyList, &BoothCompany{
			ID: company.ID,
			Name: company.CompanyName,
		})
	}

	return boothCompanyList, nil
}

func (b *boothImpl) GetBooth(c *gin.Context, userInfo *UserInfo, boothID int64) (*BoothInfo, error) {
	// Xác thực role từ jwt: company và admin.
	// Lấy thông tin từ database và return.
	// Đối với company: Hiển thị và có quyền xử lý Booth.

	booth, err := b.boothRepository.FindByID(boothID)
	if err != nil {
		return nil, err
	}
	company, err := b.companyRepository.FindByID(booth.CompanyID)
	if err != nil {
		return nil, err
	}

	companyInfo := &BoothCompany{
		ID:   company.ID,
		Name: company.CompanyName,
	}
	return &BoothInfo{
		ID:      boothID,
		Company: *companyInfo,
		Level:   booth.Level,
		Price:   booth.Price,
	}, nil
}

func (b *boothImpl) GetAllBooths(c *gin.Context, userInfo *UserInfo) ([]*BoothInfo, error) {
	// Xác thực role từ jwt: company và admin.
	// Đối với company, kiểm tra thêm phân hạng, dựa vào phân hạng mà trả về danh sách tương ứng.
	// Lấy thông tin từ database và return.
	err1 := validateAdminRole(userInfo)
	err2 := validateCompanyRole(userInfo)
	if err1 != nil && err2 != nil {
		return nil, err1
	}

	booths, err := b.boothRepository.FindAll()
	if err != nil {
		return nil, err
	}

	boothInfoList := make([]*BoothInfo, 0)
	for _, booth := range booths {
		var boothCompany BoothCompany
		if booth.CompanyID != 0 {
			company, err := b.companyRepository.FindByID(booth.CompanyID)
			if err != nil {
				return nil, err
			}
			boothCompany = BoothCompany{
				ID:   company.ID,
				Name: company.CompanyName,
			}
		}
		boothInfoList = append(boothInfoList, &BoothInfo{
			ID:      booth.ID,
			Company: boothCompany,
			Price:   booth.Price,
			Level:   booth.Level,
		})
	}

	return boothInfoList, nil
}

func (b *boothImpl) UpdateBooths(c *gin.Context, userInfo *UserInfo, boothInfoList []BoothInfo) error {
	// if err := validateAdminRole(userInfo); err != nil {
	// 	return err
	// }
	for _, boothInfo := range boothInfoList {
		booth, err := b.boothRepository.FindByID(boothInfo.ID)
		if err != nil {
			return err
		}

		booth.CompanyID = boothInfo.Company.ID
		booth.Level = boothInfo.Level
		booth.Price = boothInfo.Price

		err = b.boothRepository.Update(booth)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *boothImpl) UpdateBooth(c *gin.Context, userInfo *UserInfo, boothInfo BoothInfo) error {
	// if err := validateAdminRole(userInfo); err != nil {
	// 	return err
	// }
	booth, err := b.boothRepository.FindByID(boothInfo.ID)
	if err != nil {
		return err
	}

	booth.CompanyID = boothInfo.Company.ID
	booth.Level = boothInfo.Level
	booth.Price = boothInfo.Price

	err = b.boothRepository.Update(booth)
	if err != nil {
		return err
	}
	
	return nil
}