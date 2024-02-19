package handler

import (
	"jobfair2024/pkg"
	"jobfair2024/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Login with account
// Endpoint: /api/auth/login [POST]
type LoginInfoReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponseData struct {
	UserRole   string `json:"user_role"`
	FirstLogin bool   `json:"first_login"`
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
	}

	http.SetCookie(c.Writer, cookie)

	responseSuccess(c, &LoginResponseData{
		UserRole:   string(userInfo.Role),
		FirstLogin: userInfo.FirstLogin,
	})
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

// Google auth

var (
	config *oauth2.Config
)

func init() {
	config = &oauth2.Config{
		ClientID:     "975360526012-4hg45uff2f33576dbnf4oinqmuf31jof.apps.googleusercontent.com",
		ClientSecret: "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",
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:8080/google/callback",
		Scopes:       []string{"https://mail.google.com"},
	}
}

func (h *JobFairHandler) GoogleAuthorize(c *gin.Context) {
	// Tạo URL cho người dùng ủy quyền truy cập
	url := config.AuthCodeURL("state", oauth2.AccessTypeOffline)

	// Chuyển hướng người dùng đến trang xác thực của Google
	c.Redirect(http.StatusFound, url)
}

func (h *JobFairHandler) GoogleCallback(c *gin.Context) {
	// Lấy mã xác thực từ Google
	code := c.Query("code")

	// Sử dụng mã xác thực để trao đổi lấy token
	token, err := config.Exchange(c, code)
	if err != nil {
		responseServerError(c, pkg.ParseError(err))
		return
	}

	accessCookie := &http.Cookie{
		Name:     "googleAccessToken",
		Value:    token.AccessToken,
		Path:     "/",
		Expires:  token.Expiry,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(c.Writer, accessCookie)

	refreshCokie := &http.Cookie{
		Name:     "googleRefreshToken",
		Value:    token.RefreshToken,
		Path:     "/",
		Expires:  token.Expiry,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(c.Writer, refreshCokie)

	responseSuccess(c, "ok")
}
