package sessionService

import (
	"context"
	"github.com/nnniyaz/blog/internal/domain/base/uuid"
	"github.com/nnniyaz/blog/internal/domain/session"
	"github.com/nnniyaz/blog/internal/repos"
)

type SessionService interface {
	Create(ctx context.Context, session *session.Session) error
	DeleteBySession(ctx context.Context, session uuid.UUID) error
	DeleteByUserId(ctx context.Context, userId uuid.UUID) error
	FindAll(ctx context.Context, offset, limit int64) ([]*session.Session, int64, error)
	FindBySession(ctx context.Context, session uuid.UUID) (*session.Session, error)
	FindByUserId(ctx context.Context, userId uuid.UUID) ([]*session.Session, error)
}

type sessionService struct {
	repo repos.Session
}

func NewSessionService(sessionRepo repos.Session) SessionService {
	return &sessionService{repo: sessionRepo}
}

func (s *sessionService) Create(ctx context.Context, session *session.Session) error {
	return s.repo.Create(ctx, session)
}

func (s *sessionService) DeleteBySession(ctx context.Context, session uuid.UUID) error {
	return s.repo.DeleteBySession(ctx, session)
}

func (s *sessionService) DeleteByUserId(ctx context.Context, userId uuid.UUID) error {
	return s.repo.DeleteByUserId(ctx, userId)
}

func (s *sessionService) FindAll(ctx context.Context, offset, limit int64) ([]*session.Session, int64, error) {
	return s.repo.FindAll(ctx, offset, limit)
}

func (s *sessionService) FindBySession(ctx context.Context, session uuid.UUID) (*session.Session, error) {
	return s.repo.FindBySession(ctx, session)
}

func (s *sessionService) FindByUserId(ctx context.Context, userId uuid.UUID) ([]*session.Session, error) {
	return s.repo.FindByUserId(ctx, userId)
}
