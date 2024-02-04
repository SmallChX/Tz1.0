package usecase

import (
	"errors"
	"jobfair2024/model"
	"jobfair2024/repository"
	"time"

	"github.com/gin-gonic/gin"
)

// Đây là Request của Company đặt Booth.
// Quyền xử lý: Admin. Quyền thêm, sửa: Company.
type BoothRequestUsecase interface {
	GetRequest(c *gin.Context, userInfo *UserInfo, request int64) (*BoothRequestInfo, error)
	GetAllRequest(c *gin.Context, userInfo *UserInfo) ([]*BoothRequestInfo, error)
	CreateRequest(c *gin.Context, userInfo *UserInfo, booths *BoothRequestInfo) error
	// Một Request có thể có nhiều hơn 1 Booth.
	// Khi đăng ký, khởi tạo Request với status là Pending.
	// Trong quá trình Pending, Admin sẽ xử lý, dựa vào policy và thanh toán.
	AcceptRequest(c *gin.Context, userInfo *UserInfo, requestID int64) error
	RejectRequest(c *gin.Context, userInfo *UserInfo, requestID int64) error //  Quyền xử lý: admin. Chuyển status của Request sang Rejected.
	DeleteRequest(c *gin.Context, userInfo *UserInfo, requestID int64) error // Quyền xử lý: company. Chuyển status của Request sang Deleted.
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

func getBoothIDs(booths []model.Booth) []int64 {
	var ids []int64
	for _, booth := range booths {
		ids = append(ids, booth.ID)
	}

	return ids
}

func (b *boothRequestUsecaseImpl) GetRequest(c *gin.Context, userInfo *UserInfo, requestID int64) (*BoothRequestInfo, error) {
	request, err := b.boothRequestRepository.FindByID(requestID)
	if err != nil {
		return nil, err
	}

	return &BoothRequestInfo{
		BoothIDList:            getBoothIDs(request.Booths),
		CompanyID:              request.CompanyID,
		Type:                   string(request.Type),
		CreateAt:               request.CreateAt,
		Reason:                 request.Reason,
		DestinationBoothIDList: getBoothIDs(request.DestinationBooths),
	}, nil
}

func (b *boothRequestUsecaseImpl) GetAllRequest(c *gin.Context, userInfo *UserInfo) ([]*BoothRequestInfo, error) {
	result, err := b.boothRequestRepository.FindAll()
	if err != nil {
		return nil, err
	}

	requestList := make([]*BoothRequestInfo, len(result))
	for _, request := range result {
		requestList = append(requestList, &BoothRequestInfo{
			BoothIDList:            getBoothIDs(request.Booths),
			CompanyID:              request.CompanyID,
			Type:                   string(request.Type),
			CreateAt:               request.CreateAt,
			Reason:                 request.Reason,
			DestinationBoothIDList: getBoothIDs(request.DestinationBooths),
		})
	}
	return requestList, nil
}

func (b *boothRequestUsecaseImpl) CreateRequest(c *gin.Context, userInfo *UserInfo, requestInfo *BoothRequestInfo) error {
	if err := validateCompanyRole(userInfo); err != nil {
		return err
	}

	booths, err := b.boothRepository.FindByIds(requestInfo.BoothIDList)
	if err != nil {
		return err
	}

	var desBooths []model.Booth
	if len(requestInfo.DestinationBoothIDList) > 0 {
		desBooths, err = b.boothRepository.FindByIds(requestInfo.DestinationBoothIDList)
		if err != nil {
			return err
		}
	}

	switch requestInfo.Type {
	case string(model.RegistTypeRequest):
		if !isAvailableBooth(booths) {
			return errors.New("booths not available")
		}
		if !isContiniousBooths(booths) {
			return errors.New("booth not continious")
		}
		break
	case string(model.ChangeTypeRequest):
		if len(requestInfo.BoothIDList) != len(requestInfo.DestinationBoothIDList) {
			return errors.New("not match booths number")
		}
		if !isAvailableBooth(desBooths) {
			return errors.New("booths not available")
		}
		if !isContiniousBooths(booths) {
			return errors.New("booth not continious")
		}
		if !isCompanyBoothOwner(userInfo.ID, booths) {
			return errors.New("not match company own booths")
		}
		break
	case string(model.RemoveTypeRequest):
		if !isCompanyBoothOwner(userInfo.ID, booths) {
			return errors.New("not match company own booths")
		}
		break
	default:
		return errors.New("no match request type found!")
	}

	err = b.boothRequestRepository.Create(&model.BoothRequest{
		Booths:            booths,
		CompanyID:         userInfo.ID,
		Status:            model.PedingRequest,
		Type:              model.TypeRequest(requestInfo.Type),
		Reason:            requestInfo.Reason,
		DestinationBooths: desBooths,
		CreateAt:          time.Now(),
	})

	if err != nil {
		return err
	}

	return nil
}

func isAvailableBooth(booths []model.Booth) bool {
	for _, booth := range booths {
		if booth.CompanyID != 0 {
			return false
		}
	}

	return true
}

func isContiniousBooths(booths []model.Booth) bool {
	temp := 0
	for _, booth := range booths {
		if temp != 0 && int(booth.ID)-temp > 1 {
			return false
		}
		temp = int(booth.ID)
	}
	return true
}

func isCompanyBoothOwner(companyID int64, booths []model.Booth) bool {
	for _, booth := range booths {
		if booth.CompanyID != companyID {
			return false
		}
	}
	return true
}

func (b *boothRequestUsecaseImpl) AcceptRequest(c *gin.Context, userInfo *UserInfo, requestID int64) error {
	if err := validateAdminRole(userInfo); err != nil {
		return err
	}

	request, err := b.boothRequestRepository.FindByID(requestID)
	if err != nil {
		return err
	}

	if request.Status != model.PedingRequest {
		return errors.New("request is invalid")
	}

	switch request.Type {
	case model.RegistTypeRequest:
		if !isAvailableBooth(request.Booths) {
			return errors.New("booths not available")
		}
		for _, booth := range request.Booths {
			// Cập nhật CompanyID cho booth
			booth.CompanyID = request.CompanyID
			err = b.boothRepository.Update(&booth)
			if err != nil {
				return err
			}
		}
	case model.ChangeTypeRequest:
		desBooths := request.DestinationBooths

		if !isAvailableBooth(desBooths) {
			return errors.New("booths not available")
		}
		if !isCompanyBoothOwner(request.CompanyID, request.Booths) {
			return errors.New("not match company own booths")
		}

		for _, booth := range request.Booths {
			booth.CompanyID = 0
			err = b.boothRepository.Update(&booth)
			if err != nil {
				return err
			}
		}
		for _, booth := range desBooths {
			booth.CompanyID = request.CompanyID
			err = b.boothRepository.Update(&booth)
			if err != nil {
				return err
			}
		}
	case model.RemoveTypeRequest:
		if !isCompanyBoothOwner(requestID, request.Booths) {
			return errors.New("not match company own booths")
		}
		for _, booth := range request.Booths {
			booth.CompanyID = 0
			err = b.boothRepository.Update(&booth)
			if err != nil {
				return err
			}
		}
	default:
		return errors.New("no match request type found!")
	}

	request.Status = model.AcceptedRequest
	err = b.boothRequestRepository.Update(request)
	if err != nil {
		return err
	}

	return nil
}

func (b *boothRequestUsecaseImpl) RejectRequest(c *gin.Context, userInfo *UserInfo, requestID int64) error {
	if err := validateAdminRole(userInfo); err != nil {
		return err
	}

	request, err := b.boothRequestRepository.FindByID(requestID)
	if err != nil {
		return err
	}

	request.Status = model.RejectedRequest
	err = b.boothRequestRepository.Update(request)

	if err != nil {
		return err
	}

	return nil
}

func (b *boothRequestUsecaseImpl) DeleteRequest(c *gin.Context, userInfo *UserInfo, requestID int64) error {
	if err := validateCompanyRole(userInfo); err != nil {
		return err
	}
	request, err := b.boothRepository.FindByID(requestID)
	if err != nil {
		return err
	}
	if request.CompanyID != userInfo.ID {
		return errors.New("not have right")
	}

	err = b.boothRequestRepository.Delete(requestID)
	if err != nil {
		return err
	}

	return nil
}

func (b *boothRequestUsecaseImpl) FinishRequest(c *gin.Context, userInfo *UserInfo, requestID int64) error {
	if err := validateAdminRole(userInfo); err != nil {
		return err
	}
	
	request, err := b.boothRequestRepository.FindByID(requestID)
	if err != nil {
		return err
	}

	request.Status = model.FinishedRequest

	err = b.boothRequestRepository.Update(request)
	if err != nil {
		return err
	}

	return nil
}
