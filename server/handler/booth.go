package handler

import (
	"jobfair2024/pkg"

	"github.com/gin-gonic/gin"
)

// Get All Booth
// Endpoint: /api/booth/get-all-booth [GET]
func (h *JobFairHandler) GetAllBooths(c *gin.Context) {
	boothList, err := h.boothUsecase.GetAllBooths(c)
	if err != nil {
		responseServerError(c, pkg.ParseError(err))
		return 
	}

	responseSuccess(c, boothList)
}