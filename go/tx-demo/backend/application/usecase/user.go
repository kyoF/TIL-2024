package usecase

import (
	"backend/application/dto"
	"backend/domain/repository"
	"fmt"
)

type User interface {
	Get(userId string) (dto.User, error)
	Insert(userId, name string, age int) error
	UpdateName(userId, name string) error
	UpdateAge(userId string, age int) error
}

type user struct {
	userRepository repository.User
}

func NewUser(userRepository repository.User) User {
	return &user{userRepository: userRepository}
}

func (u *user) Get(userId string) (dto.User, error) {
	user, err := u.userRepository.Get(userId)
	if err != nil {
		return dto.User{}, err
	}

	return dto.User{
		UserId: user.UserId,
		Name:   user.Name,
		Age:    user.Age,
	}, nil
}

func (u *user) Insert(userId, name string, age int) error {
	return u.userRepository.Transaction(func() error {
		return u.userRepository.Insert(userId, name, age)
	})
}

func (u *user) UpdateName(userId, name string) error {
	return u.userRepository.Transaction(func() error {
		return u.userRepository.UpdateName(userId, name)
	})
}

func (u *user) UpdateAge(userId string, age int) error {
	var err error

	err = u.userRepository.Transaction(func() error {
		err := u.userRepository.UpdateAge(userId, age)
		if err != nil {
			return err
		}
		user, err := u.userRepository.Get(userId)
		if err != nil {
			return err
		}
		fmt.Println(user)
		return nil
	})
	if err != nil {
		return err
	}

	err = u.userRepository.Transaction(func() error {
		err := u.userRepository.UpdateAge(userId, 26)
		if err != nil {
			return err
		}
		user, err := u.userRepository.Get(userId)
		if err != nil {
			return err
		}
		fmt.Println(user)
		return nil
	})
	if err != nil {
		return err
	}

	user, err := u.userRepository.Get(userId)
	if err != nil {
		return err
	}
	fmt.Println(user)

	return nil
}
