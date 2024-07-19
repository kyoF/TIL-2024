package usecase

import (
	"backend/domain/repository"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"time"
)

type IUserUsecase interface {
	Login(name, password string) (string, error)
	Logout() error
	Signup(name, password string) error
}

type userUsecase struct {
	userRepo    repository.IUserRepository
	sessionRepo repository.ISessionRepository
}

func NewUserUsecase(
	userRepo repository.IUserRepository,
	sessionRepo repository.ISessionRepository,
) IUserUsecase {
	return &userUsecase{userRepo: userRepo, sessionRepo: sessionRepo}
}

func (uc *userUsecase) Login(name, password string) (string, error) {
	hashedPassword := hashPassword(password)

	user, err := uc.userRepo.Get(name)
	if err != nil {
		return "", err
	}

	if user.Password != hashedPassword {
		return "", errors.New("password is uncorrect")
	}

	sessionId := fmt.Sprintf("%x", sha256.Sum256([]byte(user.Name+time.Now().String())))

	err = uc.sessionRepo.Insert(sessionId, user.Name, 24*time.Hour)
	if err != nil {
		return "", err
	}

	return sessionId, nil
}

func (uc *userUsecase) Logout() error {
	return nil
}

func (uc *userUsecase) Signup(name, password string) error {
	return nil
}

func hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}
