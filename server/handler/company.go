package handler

import (
	"jobfair2024/middleware"
	"jobfair2024/pkg"

	"github.com/gin-gonic/gin"
)

// Get company BoothID
// Endpoint: /api/booth/company-booth
func (h *JobFairHandler) GetCompanyOwnedBoothIDs(c *gin.Context) {
	userInfo := middleware.GetUserInfoFromContext(c)
	if userInfo == nil {
		return
	}

	boothIDList, err := h.companyInfoUsecase.GetCompanyBoothIDs(c, userInfo)
	if err != nil {
		responseServerError(c, pkg.ParseError(err))
		return
	}

	responseSuccess(c, boothIDList)
}

// get company profile info 
// Endpoint: /api/profile/company [GET]
func (h *JobFairHandler) GetCompanyInfo(c *gin.Context) {
	userInfo := middleware.GetUserInfoFromContext(c)
	if userInfo == nil {
		return
	}

	companyInfo, err := h.companyInfoUsecase.GetCompanyInfo(c, userInfo)
	if err != nil {
		responseServerError(c, pkg.ParseError(err))
		return
	}

	responseSuccess(c, companyInfo)
}