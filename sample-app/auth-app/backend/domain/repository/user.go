package repository

type IUserRepository interface {
	Insert(name, password string) error
	Get(name string) (string, error)
}
