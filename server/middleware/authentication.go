package middleware

import (
	"jobfair2024/pkg"
	"jobfair2024/pkg/util"
	"jobfair2024/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtToken, err := c.Cookie("authToken")
		if err == http.ErrNoCookie {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}

		claims, err := util.ValidateToken(c, jwtToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		c.Set("userInfo", &usecase.UserInfo{
			ID:    claims.UserID,
			Role:  claims.UserRole,
			Email: claims.Email,
		})

		c.Next()
	}
}

func GetUserInfoFromContext(c *gin.Context) *usecase.UserInfo {
	value, exists := c.Get("userInfo")
	if !exists {
		c.JSON(http.StatusForbidden, pkg.NotExist)
		return nil
	}

	userInfo, ok := value.(*usecase.UserInfo)
	if !ok {
		c.JSON(http.StatusForbidden, pkg.NotHaveRight)
		return nil
	}
	return userInfo
}


