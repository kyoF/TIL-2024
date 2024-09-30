package repository

type DBClient interface {
	Transaction(f func() error) error
}
