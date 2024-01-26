package usecase

import (
	"errors"
	"jobfair2024/model"
	"jobfair2024/repository"
	"time"
)

// Đây là Request của Company đặt Booth.
// Quyền xử lý: Admin. Quyền thêm, sửa: Company.
type BoothRequestUsecase interface {
	GetRequest()    // Lấy một Request
	GetAllRequest() // Lấy tất cả Request
	CreateRequest() // Thêm Request (Company Regist Booths), một Company có thể có nhiều Request.
	// Một Request có thể có nhiều hơn 1 Booth.
	// Khi đăng ký, khởi tạo Request với status là Pending.
	// Trong quá trình Pending, Admin sẽ xử lý, dựa vào policy và thanh toán.
	AcceptRequest() // Quyền xử lý: admin. Chuyển status của Request sang Accepted.
	RejectRequest() //  Quyền xử lý: admin. Chuyển status của Request sang Rejected.
	DeleteRequest() // Quyền xử lý: company. Chuyển status của Request sang Deleted.
	// Đối với Reject và Delete, không xóa mà chỉ chuyển status => Xử lý và đối chứng sau này.
}

type boothRequestUsecaseImpl struct {
	boothRepository        repository.BoothRepository
	boothRequestRepository repository.BoothRequestRepository
}

func NewBoothRequestUsecase(
	boothRepository repository.BoothRepository,
	boothRequestRepository repository.BoothRequestRepository,
) BoothRequestUsecase {
	return &boothRequestUsecaseImpl{
		boothRepository:        boothRepository,
		boothRequestRepository: boothRequestRepository,
	}
}

type BoothRequestInfo struct {
	BoothIDList []int64   `json:"booth_id"`
	CompanyID   int64     `json:"company_id"`
	Type        string    `json:"type"`
	CreateAt    time.Time `json:"create_at"`
	// Some for another request
	Reason                 string  `json:"reason"` // for delete request
	DestinationBoothIDList []int64 `json:"des_booth_id"`
}

func (b *boothRequestUsecaseImpl) GetRequest(requestID int64) (*model.BoothRequest, error) {
	request, err := b.boothRequestRepository.FindByID(requestID)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func (b *boothRequestUsecaseImpl) GetAllRequest() ([]model.BoothRequest, error) {
	requests, err := b.boothRequestRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return requests, nil
}

func (b *boothRequestUsecaseImpl) CreateRequest(request *BoothRequestInfo) error {
	// Check role: company
	switch request.Type {
	case string(model.RegistTypeRequest):
		if !isAvailableBooth(b.boothRepository, request.BoothIDList) {
			return errors.New("booths not available")
		}
		if !isContiniousBooths(request.BoothIDList) {
			return errors.New("booth not continious")
		}
	case string(model.ChangeTypeRequest):
		if len(request.BoothIDList) != len(request.DestinationBoothIDList) {
			return errors.New("not match booths number")
		}
		if !isAvailableBooth(b.boothRepository, request.DestinationBoothIDList) {
			return errors.New("booths not available")
		}
		if !isContiniousBooths(request.DestinationBoothIDList) {
			return errors.New("booth not continious")
		}
		if !isCompanyBoothOwner(b.boothRepository, request.CompanyID, request.BoothIDList) {
			return errors.New("not match company own booths")
		}
	case string(model.RemoveTypeRequest):
		if !isCompanyBoothOwner(b.boothRepository, request.CompanyID, request.BoothIDList) {
			return errors.New("not match company own booths")
		}

	default:
		return errors.New("no match request type found!")
	}

	return nil
}

func isAvailableBooth(boothRepository repository.BoothRepository, boothIDList []int64) bool {
	for _, boothID := range boothIDList {
		booth, err := boothRepository.FindByID(boothID)
		if err != nil {
			return false
		}

		if booth.CompanyID != 0 {
			return false
		}
	}

	return true
}

func isContiniousBooths(boothIDList []int64) bool {
	temp := 0
	for _, boothID := range boothIDList {
		if temp != 0 && int(boothID)-temp > 1 {
			return false
		}
		temp = int(boothID)
	}
	return true
}

func isCompanyBoothOwner(boothRepository repository.BoothRepository, companyID int64, boothIDList []int64) bool {
	for _, boothID := range boothIDList {
		booth, _ := boothRepository.FindByID(boothID)
		if booth.CompanyID != companyID {
			return false
		}
	}
	return true
}
