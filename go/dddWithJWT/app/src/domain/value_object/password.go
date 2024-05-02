package value_object

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type password struct {
	value string
}

func NewPassword(pass string) (*password, error) {
	err := validation(pass)
	if err != nil {
		return nil, err
	}
	return &password{value: pass}, nil
}

func validation(pass string) error {
	if len(pass) < 8 {
		return errors.New("pass length must be at least 8 characters: %s")
	}
	return nil
}

func (p *password) Value() string {
	return p.value
}

func (p *password) Hashed() (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(p.value), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPass), nil
}

func (p *password) Check(inputPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(p.value), []byte(inputPass))
}
