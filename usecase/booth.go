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
	GetBooth(c *gin.Context, boothID int64) (*model.Booth, error)   // Lấy thông tin của Booth. Chưa biết dùng ở đâu
	GetAllBooths(c* gin.Context) ([]model.Booth, error) // Lấy tất cả danh sách của Booth
	// RegistBooth()  // Đăng ký Booth => Tạo một Request gồm Id Company (lấy từ jwt)
	// và danh sách các Booths.
	// Quyền xử lý: Admin
	// ChangeBoothCompany() // Đổi Company của Booth (đối với các Company được yêu cầu trực tiếp or muộn)
	// RemoveBoothCompany() // Xóa Company của Booth (xóa các Company khi được yêu cầu trực tiếp với admin)
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
func (b *boothImpl) GetBooth(c* gin.Context, boothID int64) (*model.Booth, error) {
	// Xác thực role từ jwt: company và admin.
	// Lấy thông tin từ database và return.
	// Đối với company: Hiển thị và có quyền xử lý Booth.

	booth, err := b.boothRepository.FindByID(boothID)
	if err != nil {
		return nil, err
	}

	return booth, nil
}

func (b *boothImpl) GetAllBooths(c* gin.Context) ([]model.Booth, error) {
	// Xác thực role từ jwt: company và admin.
	// Đối với company, kiểm tra thêm phân hạng, dựa vào phân hạng mà trả về danh sách tương ứng.
	// Lấy thông tin từ database và return.

	booths, err := b.boothRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return booths, nil
}


// Xác thực role từ jwt: admin, company
// Kiểm tra Booth đích trong database đã có Company sở hữu chưa.
// Kiểm tra phân hạng của Company, nếu không thì xem như Booth đó đã có người đăng ký???
// Create Request với Type: Change và chờ admin xử lý.
// Đổi theo số lượng tương ứng mà Company chọn??
// func (b *boothImpl) ChangeBoothCompany() {}

// Xác thực role từ jwt: admin, company.
// Lấy id từ jwt trong Company => Kiểm tra trong database xem có đúng với id Booth của Request.
// Create Request với Type: Delete, kèm lý do và chờ admin xử lý.
// func (b *boothImpl) RemoveBoothCompany() {}
