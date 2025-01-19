package session

import (
	"github.com/nnniyaz/blog/internal/domain/base/uuid"
	"github.com/nnniyaz/blog/internal/domain/session/valueobject"
	"time"
)

type Session struct {
	id        uuid.UUID
	userId    uuid.UUID
	session   uuid.UUID
	userAgent valueobject.UserAgent
	createdAt time.Time
}

func NewSession(userId uuid.UUID, userAgent string) (*Session, error) {
	ua, err := valueobject.NewUserAgent(userAgent)
	if err != nil {
		return nil, err
	}
	return &Session{
		id:        uuid.NewUUID(),
		userId:    userId,
		session:   uuid.NewUUID(),
		userAgent: ua,
		createdAt: time.Now(),
	}, nil
}

func (s *Session) GetId() uuid.UUID {
	return s.id
}

func (s *Session) GetUserId() uuid.UUID {
	return s.userId
}

func (s *Session) GetSession() uuid.UUID {
	return s.session
}

func (s *Session) GetUserAgent() valueobject.UserAgent {
	return s.userAgent
}

func (s *Session) GetCreatedAt() time.Time {
	return s.createdAt
}

func UnmarshalSessionFromDatabase(id uuid.UUID, userId uuid.UUID, session uuid.UUID, userAgent string, createdAt time.Time) *Session {
	return &Session{
		id:        id,
		userId:    userId,
		session:   session,
		userAgent: valueobject.UserAgent(userAgent),
		createdAt: createdAt,
	}
}
