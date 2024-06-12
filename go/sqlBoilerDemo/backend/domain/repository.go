package domain

type IRepository interface {
	GetUserProfiles() ([]UserProfile, error)
}
