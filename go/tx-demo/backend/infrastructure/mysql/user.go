package mysqlinfra

import (
	"backend/domain/entity"
	"backend/domain/repository"
	"backend/infrastructure/mysql/models"
)

type user struct {
	*dbClient
}

func NewUser(dbClient *dbClient) repository.User {
	return &user{
		dbClient: dbClient,
	}
}

func (u *user) Get(userId string) (*entity.User, error) {
	var user models.User
	err := u.db.Where("user_id = ?", userId).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &entity.User{
		UserId: user.UserID,
		Name:   user.Name,
		Age:    user.Age,
	}, nil
}

func (u *user) Insert(userId, name string, age int) error {
	user := &models.User{
		UserID: userId,
		Name:   name,
		Age:    age,
	}

	err := u.dbClient.DB().Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (u *user) UpdateName(userId, name string) error {
	err := u.dbClient.DB().Model(&models.User{}).Where("user_id = ?", userId).Update("name", name).Error
	return err
}

func (u *user) UpdateAge(userId string, age int) error {
	err := u.dbClient.DB().Model(&models.User{}).Where("user_id = ?", userId).Update("age", age).Error
	return err
}
