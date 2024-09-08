package repository

type Transaction interface {
	Transaction(f func() error) error
}
