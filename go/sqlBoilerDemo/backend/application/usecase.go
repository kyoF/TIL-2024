package application

import "backend/domain"

type IUsecase interface {
	GetUsers()
}

type usecase struct {
	repository domain.IRepository
}

func NewUsecase(repository domain.IRepository) IUsecase {
	return &usecase{
		repository: repository,
	}
}

func (u *usecase) GetUsers() {}
