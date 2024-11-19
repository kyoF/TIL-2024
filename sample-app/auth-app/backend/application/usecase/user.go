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

type User interface {
	Login(name, password string) (string, error)
	Logout(sessionId string) error
	Signup(name, password string) error
	Get(sessionId string) (string, error)
}

type user struct {
	userRepository      repository.User
	sessionRepository   repository.Session
	sessionQueryService queryservice.Session
}

func NewUser(
	userRepository repository.User,
	sessionRepository repository.Session,
	sessionQueryService queryservice.Session,
) User {
	return &user{
		userRepository:      userRepository,
		sessionRepository:   sessionRepository,
		sessionQueryService: sessionQueryService,
	}
}

func (uc *user) Login(name, password string) (string, error) {
	passwordvo := valueobject.NewPassword(password)

	storedPassword, err := uc.userRepository.Get(name)
	if err != nil {
		return "", err
	}

	if storedPassword != passwordvo.Hash() {
		return "", errors.New("password is uncorrect")
	}

	sessionId := fmt.Sprintf("%x", sha256.Sum256([]byte(name+time.Now().String())))

	err = uc.sessionRepository.Set(sessionId, name, 24*time.Hour)
	if err != nil {
		return "", err
	}

	return sessionId, nil
}

func (uc *user) Logout(sessinId string) error {
	err := uc.sessionRepository.Delete(sessinId)
	return err
}

func (uc *user) Signup(name, password string) error {
	passwordvo := valueobject.NewPassword(password)

	err := uc.userRepository.Insert(name, passwordvo.Hash())

	return err
}

func (uc *user) Get(sessionId string) (string, error) {
	name, err := uc.sessionQueryService.Get(sessionId)
	return name, err
}
