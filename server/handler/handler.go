package handler

import (
	"jobfair2024/usecase"
)

type JobFairHandler struct {
	boothRequestUsecase usecase.BoothRequestUsecase
	boothUsecase usecase.BoothUsecase
	authenticationUsecase usecase.AuthenticationUsecase
	companyInfoUsecase usecase.CompanyUsecase
	userAccountUsecase usecase.UserAccountUsecase
	notificationUsecase usecase.NotificationUsecase
}

func NewHandler(
	boothRequestUsecase usecase.BoothRequestUsecase,
	boothUsecase usecase.BoothUsecase,
	authenticationUsecase usecase.AuthenticationUsecase,
	companyInfoUsecase usecase.CompanyUsecase,
	userAccountUsecase usecase.UserAccountUsecase,
	notificationUsecase usecase.NotificationUsecase,
) *JobFairHandler {
	return &JobFairHandler{
		userAccountUsecase: userAccountUsecase,
		boothRequestUsecase: boothRequestUsecase,
		boothUsecase: boothUsecase,
		authenticationUsecase: authenticationUsecase,
		companyInfoUsecase: companyInfoUsecase,
		notificationUsecase: notificationUsecase,
	}
}
