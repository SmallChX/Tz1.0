package main

import (
	"jobfair2024/handler"
	"jobfair2024/repository"
	"jobfair2024/setting"
	"jobfair2024/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)

	setting.InitDB()
	setting.MigrateDB()

	boothRequestRepository := repository.NewBoothRequestRepository(setting.GetDB())
	boothRepository := repository.NewBoothRepository(setting.GetDB())

	boothRequestUsecase := usecase.NewBoothRequestUsecase(boothRepository, boothRequestRepository)
	boothUsecase := usecase.NewBoothUsecase(boothRepository)

	jobFairHandler := handler.NewHandler(boothRequestUsecase, boothUsecase)

	router := gin.Default()
	api := router.Group("/api")
	{
		// auth := api.Group("/auth")
		// auth.GET("/login/google/",)
		// auth.GET("/login/google/callback/",)
		// auth.GET("/login/account/",)
		// auth.GET("/logout/",)

		booth := api.Group("/booth")
		booth.GET("/get-all-booth", jobFairHandler.GetAllBooths)

		request := api.Group("/request")
		request.GET("/", jobFairHandler.GetRequest)
		request.GET("/get-all-request", jobFairHandler.GetAllRequests)
		request.POST("/", jobFairHandler.CreateRequest)
		request.PUT("/accept", jobFairHandler.AcceptRequest)
		request.PUT("/reject", jobFairHandler.RejectRequest)
		request.POST("/remove", jobFairHandler.RemoveRequest)
	}

	router.Run(":8080")
}
