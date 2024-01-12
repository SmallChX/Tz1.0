package usecase

type BoothRequestUsecase interface {
	GetRequest()
	GetAllRequest()
	CreateRequest()
	DeleteRequest()
}
