package queryservice

import "backend/domain/repository"

type IAuthQueriService interface {
	Get(sessionId string) (string, error)
}

type authQueryService struct {
	authRepo repository.IAuthRepository
}

func NewAuthQueryService(authRepo repository.IAuthRepository) IAuthQueriService {
	return &authQueryService{authRepo: authRepo}
}

func (q *authQueryService) Get(sessionId string) (string, error) {
	name, err := q.authRepo.Get(sessionId)
	return name, err
}
