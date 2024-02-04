package usecase

import (
	"jobfair2024/model"
	"jobfair2024/repository"

	"github.com/gin-gonic/gin"
)

// Lấy từ model Booth. Sau khi Request được accept, cập nhật Company của Booth.
// Mỗi Booth chỉ được có một Company.
// Quyền xử lý: admin, company.
type BoothUsecase interface {
	GetBooth(c *gin.Context, userInfo *UserInfo, boothID int64) (*model.Booth, error) // Lấy thông tin của Booth. Chưa biết dùng ở đâu
	GetAllBooths(c *gin.Context, userInfo *UserInfo) ([]model.Booth, error)           // Lấy tất cả danh sách của Booth

}

type boothImpl struct {
	boothRepository repository.BoothRepository
}

func NewBoothUsecase(
	boothRepository repository.BoothRepository,
) BoothUsecase {
	return &boothImpl{
		boothRepository: boothRepository,
	}
}
func (b *boothImpl) GetBooth(c *gin.Context, userInfo *UserInfo, boothID int64) (*model.Booth, error) {
	// Xác thực role từ jwt: company và admin.
	// Lấy thông tin từ database và return.
	// Đối với company: Hiển thị và có quyền xử lý Booth.

	booth, err := b.boothRepository.FindByID(boothID)
	if err != nil {
		return nil, err
	}

	return booth, nil
}

func (b *boothImpl) GetAllBooths(c *gin.Context, userInfo *UserInfo) ([]model.Booth, error) {
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

	return booths, nil
}
