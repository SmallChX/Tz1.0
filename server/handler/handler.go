package handler

import (
	"jobfair2024/usecase"
)

type JobFairHandler struct {
	boothRequestUsecase usecase.BoothRequestUsecase
	boothUsecase usecase.BoothUsecase
	authenticationUsecase usecase.AuthenticationUsecase
}

func NewHandler(
	boothRequestUsecase usecase.BoothRequestUsecase,
	boothUsecase usecase.BoothUsecase,
	authenticationUsecase usecase.AuthenticationUsecase,
) *JobFairHandler {
	return &JobFairHandler{
		boothRequestUsecase: boothRequestUsecase,
		boothUsecase: boothUsecase,
		authenticationUsecase: authenticationUsecase,
	}
}
