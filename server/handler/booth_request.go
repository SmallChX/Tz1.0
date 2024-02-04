package handler

import (
	"jobfair2024/pkg"
	"jobfair2024/usecase"

	"github.com/gin-gonic/gin"
)

type GetRequestReq struct {
	RequestID int64 `json:"request_id" binding:"required"`
}

// Get request
// Router: /api/request/ [GET]
func (h *JobFairHandler) GetRequest(c *gin.Context) {
	var req GetRequestReq
	if err := c.ShouldBind(&req); err != nil {
		responseBadRequestError(c, pkg.BindingFailure)
		return
	}
	
	userInfo := h.getUserInfoFromContext(c)
	if userInfo == nil {
		return
	}

	request, err := h.boothRequestUsecase.GetRequest(c, userInfo, req.RequestID)
	if err != nil {
		responseServerError(c, pkg.ParseError(err))
	}

	responseSuccess(c, request)
}

// Get all request
// Router: /api/request/get-all-request [GET]
func (h *JobFairHandler) GetAllRequests(c *gin.Context) {
	userInfo := h.getUserInfoFromContext(c)
	if userInfo == nil {
		return
	}

	requestList, err := h.boothRequestUsecase.GetAllRequest(c, userInfo)
	if err != nil {
		responseServerError(c, pkg.ParseError(err))
		return
	}

	responseSuccess(c, requestList)
}

// Create request = Send request from company
// Router: /api/request/ [POST]
type CreateRequestReq struct {
	BoothIDList            []int64 `json:"booth_id" binding:"required"`
	Type                   string  `json:"type" binding:"required"`
	Reason                 string  `json:"reason"` // for delete request
	DestinationBoothIDList []int64 `json:"des_booth_id"`
}

func (h *JobFairHandler) CreateRequest(c *gin.Context) {
	var req CreateRequestReq
	if err := c.ShouldBind(&req); err != nil {
		responseBadRequestError(c, pkg.BindingFailure)
		return
	}

	userInfo := h.getUserInfoFromContext(c)
	if userInfo == nil {
		return
	}

	err := h.boothRequestUsecase.CreateRequest(c, userInfo, &usecase.BoothRequestInfo{
		BoothIDList:            req.BoothIDList,
		Type:                   req.Type,
		Reason:                 req.Reason,
		DestinationBoothIDList: req.DestinationBoothIDList,
	})
	if err != nil {
		responseServerError(c, pkg.ParseError(err))
		return
	}

	responseSuccess(c, "ok")
}

type UpdateRequestReq struct {
	RequestID int64 `json:"request_id" binding:"required"`
}

// Accept request
// Router: /api/request/accept/ [PUT]
func (h *JobFairHandler) AcceptRequest(c *gin.Context) {
	var req UpdateRequestReq

	if err := c.ShouldBind(&req); err != nil {
		responseServerError(c, pkg.BindingFailure)
		return
	}

	userInfo := h.getUserInfoFromContext(c)
	if userInfo == nil {
		return
	}

	err := h.boothRequestUsecase.AcceptRequest(c, userInfo, req.RequestID)
	if err != nil {
		responseServerError(c, pkg.ParseError(err))
		return
	}

	responseSuccess(c, "ok")
}

// Reject request
// Router: /api/request/reject/ [PUT]
func (h *JobFairHandler) RejectRequest(c *gin.Context) {
	var req UpdateRequestReq

	if err := c.ShouldBind(&req); err != nil {
		responseBadRequestError(c, pkg.BindingFailure)
		return
	}

	userInfo := h.getUserInfoFromContext(c)
	if userInfo == nil {
		return
	}

	err := h.boothRequestUsecase.RejectRequest(c, userInfo, req.RequestID)
	if err != nil {
		responseBadRequestError(c, pkg.ParseError(err))
		return
	}

	responseSuccess(c, "ok")
}

// Remove request
// Router: //api/request/ [DELETE]
type DeleteRequestReq struct {
	RequestID int64 `json:"request_id" binding:"required"`
}

func (h *JobFairHandler) RemoveRequest(c *gin.Context) {
	var req DeleteRequestReq

	if err := c.ShouldBind(&req); err != nil {
		responseBadRequestError(c, pkg.BindingFailure)
		return
	}

	userInfo := h.getUserInfoFromContext(c)
	if userInfo == nil {
		return
	}

	err := h.boothRequestUsecase.DeleteRequest(c, userInfo, req.RequestID)
	if err != nil {
		responseServerError(c, pkg.ParseError(err))
		return
	}

	responseSuccess(c, "ok")
}
