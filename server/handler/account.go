package handler

import (
	"jobfair2024/middleware"
	"jobfair2024/model"
	"jobfair2024/pkg"
	"jobfair2024/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create account
// Endpoint: /api/auth/create-account [POST]
type CreateAccountReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name"`
	Role     string `json:"role" binding:"required"`
	// Role     model.UserRole `json:"role" binding:"required"`
}

func (h *JobFairHandler) CreateAccount(c *gin.Context) {
	userInfo := middleware.GetUserInfoFromContext(c)
	if userInfo == nil {
		return
	}

	var req CreateAccountReq
	if err := c.ShouldBind(&req); err != nil {
		responseServerError(c, pkg.ParseError(err))
		return
	}

	err := h.userAccountUsecase.CreateAccount(c, req.Username, req.Password, model.UserRole(req.Role), req.Name, userInfo)
	if err != nil {
		responseServerError(c, pkg.ParseError(err))
		return
	}

	responseSuccess(c, "ok")
}

// Reset password account
// Routes: /api/admin/account/reset-password
type ResetPasswordReq struct {
	UserID int64 `json:"user_id" binding:"required"`
}

func (h *JobFairHandler) ResetPassword(c *gin.Context) {
	userInfo := middleware.GetUserInfoFromContext(c)
	if userInfo == nil {
		return
	}

	var req ResetPasswordReq
	if err := c.ShouldBind(&req); err != nil {
		responseServerError(c, pkg.ParseError(err))
		return
	}

	if err := h.userAccountUsecase.ResetPassword(c, req.UserID, userInfo); err != nil {
		responseServerError(c, pkg.ParseError(err))
		return
	}

	responseSuccess(c, "200 ok")
}

// Delete account
// Routes: /api/admin/account [DELETE]

func (h *JobFairHandler) DeleteAccount(c *gin.Context) {
	value := c.Param("account_id")
	if value == "" {
		responseBadRequestError(c, pkg.BindingFailure)
	}

	userInfo := middleware.GetUserInfoFromContext(c)
	if userInfo == nil {
		return
	}
	accountID, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		responseServerError(c, pkg.ParseError(err))
		return
	}

	if err := h.userAccountUsecase.DeleteAccount(c, accountID, userInfo); err != nil {
		responseServerError(c, pkg.ParseError(err))
		return
	}

	responseSuccess(c, "200 ok")
}

// Get all account
// Routes: /api/admin/account/get-all-info

func (h *JobFairHandler) GetAllUserInfo(c *gin.Context) {
	userInfo := middleware.GetUserInfoFromContext(c)
	if userInfo == nil {
		return
	}

	userList, err := h.userAccountUsecase.GetAllUserInfo(c, userInfo)
	if err != nil {
		responseServerError(c, pkg.ParseError(err))
		return
	}

	responseSuccess(c, userList)
}

// Update useraccount Info
// Routes: /api/profile/company [PUT]
type CompanyInfoReq struct {
	RepresentName        string `json:"represent_name"`
	RepresentPhoneNumber string `json:"represent_phone_number"`
	RepresentMail        string `json:"represent_mail"`
}

func (h *JobFairHandler) UpdateCompanyAccountInfo(c *gin.Context) {

	userInfo := middleware.GetUserInfoFromContext(c)
	if userInfo == nil {
		return
	}

	var req CompanyInfoReq
	if err := c.ShouldBind(&req); err != nil {
		responseBadRequestError(c, pkg.BindingFailure)
		return
	}

	err := h.userAccountUsecase.UpdateUserAccountInfo(c, userInfo, usecase.CompanyUpdateInfo{
		RepresentName:        req.RepresentName,
		RepresentPhoneNumber: req.RepresentPhoneNumber,
		RepresentMail:        req.RepresentMail,
	})

	if err != nil {
		responseServerError(c, pkg.ParseError(err))
		return
	}

	responseSuccess(c, "ok")

}
