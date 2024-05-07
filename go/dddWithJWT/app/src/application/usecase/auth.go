package usecase

import (
	"app/src/domain/entity"
	"app/src/domain/repository"
	"app/src/domain/value_object"
	"errors"
)

type IAuthUsecase interface {
	Signup(username, email, password string) (*entity.Auth, error)
	Login(email, password string) (string, *entity.Auth, error)
}

type authUsecase struct {
	authRepo repository.IAuthRepository
}

func NewUseCase(authRepo repository.IAuthRepository) IAuthUsecase {
	return &authUsecase{
		authRepo: authRepo,
	}
}

func (uc *authUsecase) Signup(userId, email, password string) (*entity.Auth, error) {
	existUser, err := uc.authRepo.GetAuthByEmail(email)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	if existUser.UserId == "" {
		return nil, errors.New(err.Error())
	}

	hashedPassword, err := value_object.NewPassword(password)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	a := &entity.Auth{
		UserId:   userId,
		Email:    email,
		Password: hashedPassword.Value(),
	}

	authUser, err := uc.authRepo.CreateAuth(*a)
	if err != nil {
		return nil, err
	}

	return authUser, nil
}

func (uc *authUsecase) Login(email, pass string) (string, *entity.Auth, error) {
	user, err := uc.authRepo.GetAuthByEmail(email)
	if err != nil {
		return "", nil, errors.New(err.Error())
	}
	if user.UserId == "" {
		return "", nil, errors.New(err.Error())
	}

	hashedPassword, err := value_object.NewPassword(pass)
	if err != nil {
		return "", nil, errors.New(err.Error())
	}
	err = hashedPassword.Check(user.Password)
	if err != nil {
		return "", nil, errors.New(err.Error())
	}

	signedString, err := utils.GenerateSignedString(user.UserId, "")
	if err != nil {
		return "", nil, &myerror.InternalServerError{Err: err}
	}

	return signedString, user, nil
}
