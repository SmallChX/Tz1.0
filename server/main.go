package main

import (
	"jobfair2024/setting"

	"github.com/gin-gonic/gin"
)

func main() {
	setting.InitDB()
	// setting.MigrateDB()

	router := gin.Default()
	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		auth.GET("/login/google/",)
		auth.GET("/login/google/callback/",)
		auth.GET("/login/account/",)
		auth.GET("/logout/",)

	}

}
