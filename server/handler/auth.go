package handler

import (
	"jobfair2024/pkg"
	"jobfair2024/pkg/util"
	"jobfair2024/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *JobFairHandler) getUserInfoFromContext(c *gin.Context) *usecase.UserInfo {
	value, exists := c.Get("userInfo")
	if !exists {
		responseNotAuthorized(c, pkg.NotExist)
		return nil
	}

	userInfo, ok := value.(*usecase.UserInfo)
	if !ok {
		responseServerError(c, pkg.GeneralFailure)
		return nil
	}
	return userInfo
}

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

	signedToken, expiredTime, err := util.GenerateToken(c, userInfo)
	if err != nil {
		responseServerError(c, pkg.ParseError(err))
		return
	}

	cookie := &http.Cookie{
		Name:     "authToken",
		Value:    signedToken,
		Path:     "/",
		Expires:  expiredTime,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(c.Writer, cookie)

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
	userInfo := h.getUserInfoFromContext(c)
	if userInfo == nil {
		return
	}

	var req CreateAccountReq
	if err := c.ShouldBind(&req); err != nil {
		responseServerError(c, pkg.ParseError(err))
		return
	}

	err := h.authenticationUsecase.CreateAccount(c, req.Username, req.Password, userInfo)
	if err != nil {
		responseServerError(c, pkg.ParseError(err))
		return
	}

	responseSuccess(c, "ok")
}
