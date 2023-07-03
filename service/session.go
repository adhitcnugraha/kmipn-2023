package service

import (
	"kmipn-2023/model"
	repo "kmipn-2023/repository"
)

type SessionService interface {
	GetSessionByEmail(email string) (model.Session, error)
}

type sessionService struct {
	sessionRepo repo.SessionRepository
}

func NewSessionService(sessionRepo repo.SessionRepository) *sessionService {
	return &sessionService{sessionRepo}
}

func (c *sessionService) GetSessionByEmail(email string) (model.Session, error) {
	session, err := c.sessionRepo.SessionAvailEmail(email)
	if err != nil {
		return model.Session{}, nil
	}
	return session, nil
}
