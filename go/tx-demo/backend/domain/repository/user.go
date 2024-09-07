package repository

type User interface {
	Update(userId string) error
}
