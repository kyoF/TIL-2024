package repositoryimpl

import (
	"dddWithJWT/domain/entity"
	"dddWithJWT/domain/repository"
	value_objects "dddWithJWT/domain/value_object"
	"hash"
	"log"

	"gorm.io/gorm"
)

type authInfra struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) repository.AuthRepository {
	return &authInfra{db: db}
}

func (i *authInfra) CreateAuth(auth entity.Auth) (*entity.Auth, error) {
    hashedPass, err := value_objects.NewPassword(auth.Password)
    if err != nil {
        log.Fatal(err)
    }
    user := entity.Auth{
        UserId: auth.UserId,
        Email: auth.Email,
        Password: hashedPass.Value(),
    }
	var lastInsertId int
	query := "INSERT INTO users(username, email, password) VALUES ($1, $2, $3) returning id"
	err := ri.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Password).Scan(&lastInsertId)
	if err != nil {
		return &model.User{}, err
	}

	user.ID = int64(lastInsertId)
	return user, nil
}

func (i *authInfra) GetAuthByEmail(email string) (*entity.Auth, error) {
	u := model.User{}
	query := "SELECT id, username, email, password FROM users WHERE email = $1"
	err := ri.db.QueryRowContext(ctx, query, email).Scan(&u.ID, &u.Username, &u.Email, &u.Password)
	if err != nil {
		return &model.User{}, nil
	}

	return &u, nil
}
