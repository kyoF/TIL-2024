package usecase

import (
	queryservice "backend/application/queryservice/interface"
	"backend/domain/repository"
	"backend/domain/valueobject"
	"crypto/sha256"
	"errors"
	"fmt"
	"time"
)

type IUserUsecase interface {
	Login(name, password string) (string, error)
	Logout(sessionId string) error
	Signup(name, password string) error
    Get(sessionId string) (string, error)
}

type userUsecase struct {
	userRepo         repository.IUserRepository
	authRepo         repository.IAuthRepository
	authQueryService queryservice.Auth
}

func NewUserUsecase(
	userRepo repository.IUserRepository,
	authRepo repository.IAuthRepository,
	authQueryService queryservice.Auth,
) IUserUsecase {
	return &userUsecase{
		userRepo:         userRepo,
		authRepo:         authRepo,
		authQueryService: authQueryService,
	}
}

func (uc *userUsecase) Login(name, password string) (string, error) {
	passwordvo := valueobject.NewPassword(password)

	storedPassword, err := uc.userRepo.Get(name)
	if err != nil {
		return "", err
	}

	if storedPassword != passwordvo.Hash() {
		return "", errors.New("password is uncorrect")
	}

	sessionId := fmt.Sprintf("%x", sha256.Sum256([]byte(name+time.Now().String())))

	err = uc.authRepo.Set(sessionId, name, 24*time.Hour)
	if err != nil {
		return "", err
	}

	return sessionId, nil
}

func (uc *userUsecase) Logout(sessinId string) error {
	err := uc.authRepo.Delete(sessinId)
	return err
}

func (uc *userUsecase) Signup(name, password string) error {
	passwordvo := valueobject.NewPassword(password)

	err := uc.userRepo.Insert(name, passwordvo.Hash())

	return err
}

func (uc *userUsecase) Get(sessionId string) (string, error) {
    name, err := uc.authQueryService.Get(sessionId)
    return name, err
}
