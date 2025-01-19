package authService

import (
	"context"
	"errors"
	"github.com/nnniyaz/blog/internal/domain/base/email"
	"github.com/nnniyaz/blog/internal/domain/base/uuid"
	"github.com/nnniyaz/blog/internal/domain/session"
	exceptions2 "github.com/nnniyaz/blog/internal/domain/session/exceptions"
	"github.com/nnniyaz/blog/internal/domain/user"
	"github.com/nnniyaz/blog/internal/domain/user/exceptions"
	"github.com/nnniyaz/blog/internal/repos"
	sessionService "github.com/nnniyaz/blog/internal/services/session"
	"go.mongodb.org/mongo-driver/mongo"
	"sort"
)

const maxSessionsCount = 5

type AuthService interface {
	Login(ctx context.Context, email, password, userAgent string) (uuid.UUID, error)
	Logout(ctx context.Context, session string) error
	UserCheck(ctx context.Context, session string, userAgent string) (*user.User, error)
}

type authService struct {
	userRepo       repos.User
	sessionService sessionService.SessionService
}

func NewAuthService(userRepo repos.User, sessionService sessionService.SessionService) AuthService {
	return &authService{userRepo: userRepo, sessionService: sessionService}
}

func (a *authService) Login(ctx context.Context, rawEmail, password, userAgent string) (uuid.UUID, error) {
	// Parse email
	convertedEmail, err := email.NewEmail(rawEmail)
	if err != nil {
		return uuid.Nil, err
	}

	// Find user by email
	user, err := a.userRepo.FindByEmail(ctx, convertedEmail)
	if err != nil {
		// If user not found
		if errors.Is(err, mongo.ErrNoDocuments) {
			return uuid.Nil, exceptions.ErrUserEmailNotFound
		}

		// If error occurred
		return uuid.Nil, err
	}

	// Compare password
	if !user.GetPassword().Compare(password) {
		return uuid.Nil, exceptions.ErrUserPasswordInvalid
	}

	// Find user sessions
	sessions, err := a.sessionService.FindByUserId(ctx, user.GetId())
	if err != nil {
		return uuid.Nil, err
	}

	// If sessions count is more than maxSessionsCount
	if len(sessions) >= maxSessionsCount {
		// Sort sessions by createdAt
		sort.Slice(sessions, func(i, j int) bool {
			return sessions[i].GetCreatedAt().Before(sessions[j].GetCreatedAt())
		})

		// Delete oldest session
		if err = a.sessionService.DeleteBySession(ctx, sessions[0].GetId()); err != nil {
			return uuid.Nil, err
		}
	}

	// Create new session instance
	newSession, err := session.NewSession(user.GetId(), userAgent)
	if err != nil {
		return uuid.Nil, err
	}

	// Create new session in database
	if err = a.sessionService.Create(ctx, newSession); err != nil {
		return uuid.Nil, err
	}

	// Return session id
	return newSession.GetSession(), nil
}

func (a *authService) Logout(ctx context.Context, sessionId string) error {
	convertedSessionId, err := uuid.UUIDFromString(sessionId)
	if err != nil {
		return err
	}
	return a.sessionService.DeleteBySession(ctx, convertedSessionId)
}

func (a *authService) UserCheck(ctx context.Context, sessionId string, userAgent string) (*user.User, error) {
	convertedSessionId, err := uuid.UUIDFromString(sessionId)
	if err != nil {
		return nil, err
	}

	// Find session by id
	session, err := a.sessionService.FindBySession(ctx, convertedSessionId)
	if err != nil {
		return nil, err
	}

	// Check if user agent is equal
	if session.GetUserAgent().String() != userAgent {
		return nil, exceptions2.ErrUserAgentInvalid
	}

	// Find user by id
	return a.userRepo.FindById(ctx, session.GetUserId())
}
