package middleware

import (
	"jobfair2024/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EnsureProfileUpdated() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Lấy thông tin người dùng từ context, ví dụ qua token
		user := GetUserInfoFromContext(c)

		// Kiểm tra nếu là đăng nhập lần đầu và thông tin chưa được cập nhật
		if user.FirstLogin {
			// Trả về lỗi hoặc thông báo yêu cầu cập nhật thông tin cá nhân
			c.JSON(http.StatusBadRequest, pkg.InvalidProfile)
			c.Abort() // Ngăn không cho xử lý tiếp các handler khác
			return
		}

		c.Next() // Tiếp tục xử lý request nếu đã cập nhật thông tin
	}
}
