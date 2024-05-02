package repositoryimpl

import (
	"app/src/domain/entity"
	"app/src/domain/repository"
	"app/src/domain/value_object"

	"gorm.io/gorm"
)

type authInfra struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) repository.AuthRepository {
	return &authInfra{db: db}
}

func (i *authInfra) CreateAuth(auth entity.Auth) (*entity.Auth, error) {
	hashedPass, err := value_object.NewPassword(auth.Password)
	if err != nil {
		return nil, err
	}
	authUser := entity.Auth{
		UserId:   auth.UserId,
		Email:    auth.Email,
		Password: hashedPass.Value(),
	}

	err = i.db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&authUser).Error
	})
	if err != nil {
		return nil, err
	}

	return &authUser, nil
}

func (i *authInfra) GetAuthByEmail(email string) (*entity.Auth, error) {
	var authUser entity.Auth
	err := i.db.Where("email = ?", email).First(&authUser).Error
	if err != nil {
		return nil, err
	}

	return &authUser, nil
}
