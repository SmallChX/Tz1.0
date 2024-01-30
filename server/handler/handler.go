package handler

import (
	"jobfair2024/usecase"
)

type JobFairHandler struct {
	boothRequestUsecase usecase.BoothRequestUsecase
	boothUsecase usecase.BoothUsecase
}

func NewHandler(
	boothRequestUsecase usecase.BoothRequestUsecase,
	boothUsecase usecase.BoothUsecase,
) *JobFairHandler {
	return &JobFairHandler{
		boothRequestUsecase: boothRequestUsecase,
		boothUsecase: boothUsecase,
	}
}
