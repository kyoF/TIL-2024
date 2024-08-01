package application

import "backend/domain"

type IUsecase interface {
	GetUserProfiles() ([]domain.UserProfile, error)
}

type usecase struct {
	repository domain.IRepository
}

func NewUsecase(repository domain.IRepository) IUsecase {
	return &usecase{
		repository: repository,
	}
}

func (u *usecase) GetUserProfiles() ([]domain.UserProfile, error) {
	return u.repository.GetUserProfiles()
}
