package infrastructure

import (
	"backend/domain"
	"backend/infrastructure/boil/models"
	"context"
	"database/sql"
)

type mysqlInfrastructure struct {
	db *sql.DB
}

func NewInfrastructure(db *sql.DB) domain.IRepository {
	return &mysqlInfrastructure{
		db: db,
	}
}

func (i *mysqlInfrastructure) GetUserProfiles() ([]domain.Profile, error) {
	profiles := []domain.Profile{}
	users, err := models.Users().All(context.Background(), i.db)
	if err != nil {
		return []domain.Profile{}, nil
	}
	for _, user := range users {
		profiles = append(profiles, domain.Profile{
			UserId:  user.UserID,
			Name:    user.Name.String,
			Profile: user.Profile.String,
		})
	}
	return []domain.Profile{}, nil
}
