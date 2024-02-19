package handler

import (
	"jobfair2024/middleware"
	"jobfair2024/pkg"
	"jobfair2024/usecase"

	"github.com/gin-gonic/gin"
)

// Get All Booth
// Endpoint: /api/booth/get-all-booth [GET]
func (h *JobFairHandler) GetAllBooths(c *gin.Context) {
	userInfo := middleware.GetUserInfoFromContext(c)
	if userInfo == nil {
		return
	}

	boothList, err := h.boothUsecase.GetAllBooths(c, userInfo)
	if err != nil {
		responseServerError(c, pkg.ParseError(err))
		return
	}

	responseSuccess(c, boothList)
}

type UpdateBoothReq struct {
	ID          int64  `json:"booth_id"`
	CompanyID   int64  `json:"company_id"`
	CompanyName string `json:"company_name"`
	Level       int64  `json:"level"`
	Price       int64  `json:"price"`
}

// Update list of booth
// Endpoint: /api/booth/ [PUT]
func (h *JobFairHandler) UpdateBooth(c *gin.Context) {
	userInfo := middleware.GetUserInfoFromContext(c)
	if userInfo == nil {
		return
	}

	var req UpdateBoothReq
	if err := c.ShouldBind(&req); err != nil {
		responseBadRequestError(c, pkg.BindingFailure)
		return
	}

	err := h.boothUsecase.UpdateBooth(c, userInfo, usecase.BoothInfo{
		ID:      req.ID,
		Company: usecase.BoothCompany{ID: req.CompanyID, Name: req.CompanyName},
		Level:   int(req.Level),
		Price:   int(req.Price),
	})

	if err != nil {
		responseServerError(c, pkg.ParseError(err))
		return
	}

	responseSuccess(c, "ok")
}

// Get all company for booth check handle
// endpoint: /api/booth/company
func (h *JobFairHandler) GetAllBoothCompany(c *gin.Context) {
	userInfo := middleware.GetUserInfoFromContext(c)
	if userInfo == nil {
		return
	}

	companyBoothList, err := h.boothUsecase.GetAllBoothCompany(c, userInfo)
	if err != nil {
		responseServerError(c, pkg.ParseError(err))
		return
	}

	responseSuccess(c, companyBoothList)
}
