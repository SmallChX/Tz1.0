package handler

import (
	"jobfair2024/pkg"
	"jobfair2024/pkg/util"

	"github.com/gin-gonic/gin"
)

// Login with account
// Endpoint: /api/auth/login [POST]
type LoginInfoReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *JobFairHandler) LoginWithAccount(c *gin.Context) {
	var req LoginInfoReq
	if err := c.ShouldBind(&req); err != nil {
		responseBadRequestError(c, pkg.BindingFailure)
		return
	}

	userInfo, err := h.authenticationUsecase.Login(c, req.Username, req.Password)
	if err != nil {
		responseServerError(c, pkg.ParseError(err))
		return
	}

	err = util.GenerateToken(c, userInfo)
	if err != nil {
		responseServerError(c, pkg.ParseError(err))
		return
	}

	responseSuccess(c, "ok")
}

// Logout
// Endpoint: /api/auth/logout [POST]
func (h *JobFairHandler) Logout(c *gin.Context) {
	c.SetCookie(
		"authToken",
		"",
		0,
		"/",
		"",
		false,
		true,
	)

	responseSuccess(c, "ok")
}

// Create account
// Endpoint: /api/auth/create-account [POST]
type CreateAccountReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *JobFairHandler) CreateAccount(c *gin.Context) {
	var req CreateAccountReq
	if err := c.ShouldBind(&req); err != nil {
		responseServerError(c, pkg.ParseError(err))
		return
	}

	err := h.authenticationUsecase.CreateAccount(c, req.Username, req.Password)
	if err != nil {
		responseServerError(c, pkg.ParseError(err))
	}

	responseSuccess(c, "ok")
}
