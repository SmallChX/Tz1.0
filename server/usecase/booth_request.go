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
	GetCompanyBoothRequest(c *gin.Context, userInfo *UserInfo) ([]*BoothRequestInfo, error)
	CreateRequest(c *gin.Context, userInfo *UserInfo, booths *BoothRequestInfo) error
	// Một Request có thể có nhiều hơn 1 Booth.
	// Khi đăng ký, khởi tạo Request với status là Pending.
	// Trong quá trình Pending, Admin sẽ xử lý, dựa vào policy và thanh toán.
	AcceptRequest(c *gin.Context, userInfo *UserInfo, requestID int64) error
	RejectRequest(c *gin.Context, userInfo *UserInfo, requestID int64) error //  Quyền xử lý: admin. Chuyển status của Request sang Rejected.
	DeleteRequest(c *gin.Context, userInfo *UserInfo, requestID int64) error // Quyền xử lý: company. Chuyển status của Request sang Deleted.
	FinishRequest(c *gin.Context, userInfo *UserInfo, requestID int64) error
	GetRequestPaymentInfo(c *gin.Context, userInfo *UserInfo, requestID int64) (*PaymentInfo, error)
	UpdateRequestList(c *gin.Context, userInfo *UserInfo, requestList []RequestUpdateInfo) error
	// Đối với Reject và Delete, không xóa mà chỉ chuyển status => Xử lý và đối chứng sau này.
}

type boothRequestUsecaseImpl struct {
	boothRepository        repository.BoothRepository
	boothRequestRepository repository.BoothRequestRepository
	companyInfoRepository  repository.CompanyInformationRepository
}

func NewBoothRequestUsecase(
	boothRepository repository.BoothRepository,
	boothRequestRepository repository.BoothRequestRepository,
	companyInfoRepository repository.CompanyInformationRepository,
) BoothRequestUsecase {
	return &boothRequestUsecaseImpl{
		boothRepository:        boothRepository,
		boothRequestRepository: boothRequestRepository,
		companyInfoRepository:  companyInfoRepository,
	}
}

type BoothRequestInfo struct {
	RequestID   int64               `json:"id"`
	BoothIDList []int64             `json:"booth_id"`
	CompanyID   string              `json:"company_name"`
	Status      model.StatusRequest `json:"status"`
	Type        string              `json:"type"`
	CreateAt    time.Time           `json:"create_at"`
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
	company, err := b.companyInfoRepository.FindByID(request.CompanyID)
	if err != nil {
		return nil, err
	}

	return &BoothRequestInfo{
		RequestID:              requestID,
		BoothIDList:            getBoothIDs(request.Booths),
		CompanyID:              company.CompanyName,
		Type:                   string(request.Type),
		Status:                 request.Status,
		CreateAt:               request.CreateAt,
		Reason:                 request.Reason,
		DestinationBoothIDList: getBoothIDs(request.DestinationBooths),
	}, nil
}

func (b *boothRequestUsecaseImpl) GetCompanyBoothRequest(c *gin.Context, userInfo *UserInfo) ([]*BoothRequestInfo, error) {
	if err := validateCompanyRole(userInfo); err != nil {
		return nil, err
	}

	company, err := b.companyInfoRepository.FindByUserID(userInfo.ID)
	if err != nil {
		return nil, err
	}

	result, err := b.boothRequestRepository.GetCompanyBoothRequests(company.ID)
	if err != nil {
		return nil, err
	}

	requestList := make([]*BoothRequestInfo, 0)
	for _, request := range result {
		requestList = append(requestList, &BoothRequestInfo{
			RequestID:              request.RequestID,
			BoothIDList:            getBoothIDs(request.Booths),
			CompanyID:              company.CompanyName,
			Status:                 request.Status,
			Type:                   string(request.Type),
			CreateAt:               request.CreateAt,
			Reason:                 request.Reason,
			DestinationBoothIDList: getBoothIDs(request.DestinationBooths),
		})
	}

	return requestList, nil
}

func (b *boothRequestUsecaseImpl) GetAllRequest(c *gin.Context, userInfo *UserInfo) ([]*BoothRequestInfo, error) {
	result, err := b.boothRequestRepository.FindAll()
	if err != nil {
		return nil, err
	}

	requestList := make([]*BoothRequestInfo, 0)

	for _, request := range result {
		company, err := b.companyInfoRepository.FindByID(request.CompanyID)
		if err != nil {
			return nil, err
		}
		requestList = append(requestList, &BoothRequestInfo{
			RequestID:              request.RequestID,
			BoothIDList:            getBoothIDs(request.Booths),
			CompanyID:              company.CompanyName,
			Type:                   string(request.Type),
			Status:                 request.Status,
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

	company, err := b.companyInfoRepository.FindByUserID(userInfo.ID)
	if err != nil {
		return err
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
		if !isCompanyBoothOwner(company.ID, booths) {
			return errors.New("not match company own booths")
		}
		break
	case string(model.RemoveTypeRequest):
		if !isCompanyBoothOwner(company.ID, booths) {
			return errors.New("not match company own booths")
		}
		break
	default:
		return errors.New("no match request type found!")
	}

	err = b.boothRequestRepository.Create(&model.BoothRequest{
		Booths:            booths,
		CompanyID:         company.ID,
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

		company, err := b.companyInfoRepository.FindByID(request.CompanyID)
		if err != nil {
			return err
		}

		company.Booths = request.Booths
		err = b.companyInfoRepository.Update(company)
		if err != nil {
			return err
		}

		request.Status = model.AcceptedRequest
		err = b.boothRequestRepository.Update(request)
		if err != nil {
			return err
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
		request.Status = model.FinishedRequest
		err = b.boothRequestRepository.Update(request)
		if err != nil {
			return err
		}
		break
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
		request.Status = model.FinishedRequest
		err = b.boothRequestRepository.Update(request)
		if err != nil {
			return err
		}
	default:
		return errors.New("no match request type found")
	}

	// Hủy các request còn lại
	remainCompanyRequest, err := b.boothRequestRepository.GetCompanyBoothRequests(request.CompanyID)
	if err != nil {
		return err
	}
	for _, remainRequest := range remainCompanyRequest {
		if remainRequest.RequestID != requestID && remainRequest.Status == model.PedingRequest {
			remainRequest.Status = model.RejectedRequest
			if err := b.boothRequestRepository.Update(&remainRequest); err != nil {
				return err
			}
		}
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

	request, err := b.boothRequestRepository.FindByID(requestID)
	if err != nil {
		return err
	}
	company, err := b.companyInfoRepository.FindByUserID(userInfo.ID)

	if request.CompanyID != company.ID {
		return errors.New("not have right")
	}

	if request.Status == model.AcceptedRequest {
		return errors.New("invalid request to delete")
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

type PaymentInfo struct {
	Amount      string  `json:"amount"`
	BoothIDs    []int64 `json:"booths_id"`
	CompanyName string  `json:"company_name"`
}

func (b *boothRequestUsecaseImpl) GetRequestPaymentInfo(c *gin.Context, userInfo *UserInfo, requestID int64) (*PaymentInfo, error) {
	err1 := validateAdminRole(userInfo)
	err2 := validateCompanyRole(userInfo)

	if err1 != nil && err2 != nil {
		return nil, err1
	}

	request, err := b.boothRequestRepository.FindByID(requestID)
	if err != nil {
		return nil, err
	}

	company, err := b.companyInfoRepository.FindByID(request.CompanyID)
	if err != nil {
		return nil, err
	}

	countBooths := len(request.Booths)
	var amount string
	if countBooths == 2 {
		amount = "24.000.000"
	} else {
		amount = "12.000.000"
	}
	return &PaymentInfo{
		CompanyName: company.CompanyName,
		BoothIDs:    getBoothIDs(request.Booths),
		Amount:      amount,
	}, nil
}

type RequestUpdateInfo struct {
	RequestID int64  `json:"id"`
	Action    string `json:"action"`
}

func (b *boothRequestUsecaseImpl) UpdateRequestList(c *gin.Context, userInfo *UserInfo, requestList []RequestUpdateInfo) error {
	if err := validateAdminRole(userInfo); err != nil {
		return err
	}

	for _, request := range requestList {
		switch request.Action {
		case "accept":
			err := b.AcceptRequest(c, userInfo, request.RequestID)
			if err != nil {
				return err
			}
		case "reject":
			err := b.RejectRequest(c, userInfo, request.RequestID)
			if err != nil {
				return err
			}
		case "finish":
			err := b.FinishRequest(c, userInfo, request.RequestID)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
