package domain

type IRepository interface {
	GetUserProfiles() ([]Profile, error)
}
