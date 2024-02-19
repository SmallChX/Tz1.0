package main

import (
	"jobfair2024/handler"
	"jobfair2024/middleware"
	"jobfair2024/repository"
	"jobfair2024/setting"
	"jobfair2024/usecase"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	setting.InitDB()
	setting.MigrateDB()

	boothRequestRepository := repository.NewBoothRequestRepository(setting.GetDB())
	boothRepository := repository.NewBoothRepository(setting.GetDB())
	userAccountRepository := repository.NewUserAccountRepository(setting.GetDB())
	companyInfoRepository := repository.NewCompanyInformationRepository(setting.GetDB())
	studentRepository := repository.NewStudentRepository(setting.GetDB())
	adminRepository := repository.NewAdminRepository(setting.GetDB())
	notificationRepository := repository.NewNotificationRepository(setting.GetDB())

	boothRequestUsecase := usecase.NewBoothRequestUsecase(boothRepository, boothRequestRepository, companyInfoRepository)
	boothUsecase := usecase.NewBoothUsecase(boothRepository, companyInfoRepository)
	authenticationUsecase := usecase.NewAuthenticationUsecase(userAccountRepository)
	companyInfoUsecase := usecase.NewCompanyUsecase(companyInfoRepository)
	userAccountUsecase := usecase.NewUserAccountUsecase(userAccountRepository, companyInfoRepository, adminRepository, studentRepository)
	notificationUsecase := usecase.NewNotificationUsecase(notificationRepository, userAccountRepository)

	jobFairHandler := handler.NewHandler(boothRequestUsecase, boothUsecase, authenticationUsecase, companyInfoUsecase, userAccountUsecase, notificationUsecase)
	
	router := gin.Default()
		auth := router.Group("/auth")
		auth.GET("/google/authorize", jobFairHandler.GoogleAuthorize)
		auth.GET("/google/callback/", jobFairHandler.GoogleCallback)
		// auth.GET("/login/account/",)
		// auth.GET("/logout/",)
		auth.POST("/login", jobFairHandler.LoginWithAccount)
		auth.POST("/logout", jobFairHandler.Logout)

		booth := router.Group("/booth")
		booth.GET("/get-all-booth", middleware.AuthMiddleware(), jobFairHandler.GetAllBooths)
		booth.GET("/company-owned-booth", middleware.AuthMiddleware(), jobFairHandler.GetCompanyOwnedBoothIDs)
		booth.GET("/company", middleware.AuthMiddleware(), jobFairHandler.GetAllBoothCompany)
		booth.PUT("/", middleware.AuthMiddleware(), jobFairHandler.UpdateBooth)

		request := router.Group("/request")
		request.GET("/", middleware.AuthMiddleware(), jobFairHandler.GetRequest)
		request.GET("/get-all-request", middleware.AuthMiddleware(), jobFairHandler.GetAllRequests)
		request.GET("/company", middleware.AuthMiddleware(), jobFairHandler.GetCompanyRequests)
		request.POST("/", middleware.AuthMiddleware(), jobFairHandler.CreateRequest)
		request.PUT("/accept", middleware.AuthMiddleware(), jobFairHandler.AcceptRequest)
		request.PUT("/reject", middleware.AuthMiddleware(), jobFairHandler.RejectRequest)
		request.PUT("/finish", middleware.AuthMiddleware(), jobFairHandler.FinishRequest)
		request.PUT("/handle-list", middleware.AuthMiddleware(), jobFairHandler.HandleRequestList)
		request.DELETE("/:request_id", middleware.AuthMiddleware(), jobFairHandler.RemoveRequest)
		request.GET("/:request_id/payment", middleware.AuthMiddleware(), jobFairHandler.GetPayment)

		account := router.Group("/admin/account")
		account.GET("/get-all-info", middleware.AuthMiddleware(), jobFairHandler.GetAllUserInfo)
		account.DELETE("/:account_id", middleware.AuthMiddleware(), jobFairHandler.DeleteAccount)
		account.POST("/reset-password", middleware.AuthMiddleware(), jobFairHandler.ResetPassword)
		account.POST("/", middleware.AuthMiddleware(), jobFairHandler.CreateAccount)

		notification := router.Group("/notification")
		notification.POST("/", middleware.AuthMiddleware(), jobFairHandler.SendNotification)

		profile := router.Group("/profile")
		profile.PUT("/company", middleware.AuthMiddleware(), jobFairHandler.UpdateCompanyAccountInfo)
		profile.GET("/company", middleware.AuthMiddleware(), jobFairHandler.GetCompanyInfo)


	router.Run(":8080")
}
