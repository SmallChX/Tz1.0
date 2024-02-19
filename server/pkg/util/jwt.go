package util

import (
	"errors"
	"jobfair2024/model"
	"jobfair2024/usecase"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	UserID   int64          `json:"user_id"`
	UserRole model.UserRole `json:"user_role"`
	Email    *string
	jwt.StandardClaims
}

func GenerateToken(c *gin.Context, userInfo *usecase.UserInfo) (string, time.Time, error) {
	tokenLifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))
	if err != nil {
		return "", time.Time{}, err
	}

	expiredTime := time.Now().Add(time.Hour * time.Duration(tokenLifespan))
	claims := &Claims{
		UserID:   userInfo.ID,
		UserRole: userInfo.Role,
		Email:    userInfo.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("API_SECRET")))
	if err != nil {
		return "", time.Time{}, err
	}

	return signedToken, expiredTime, nil
}

func ValidateToken(c *gin.Context, tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("API_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
