package repository

type User interface {
	Insert(name, password string) error
	Get(name string) (string, error)
}
