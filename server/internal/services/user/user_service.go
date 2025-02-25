package userService

import (
	"context"
	"errors"
	"github.com/nnniyaz/blog/internal/domain/base/email"
	"github.com/nnniyaz/blog/internal/domain/base/uuid"
	"github.com/nnniyaz/blog/internal/domain/user"
	"github.com/nnniyaz/blog/internal/domain/user/exceptions"
	"github.com/nnniyaz/blog/internal/repos"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	Create(ctx context.Context, email, password, role string) error
	UpdateEmail(ctx context.Context, id, email string) error
	UpdatePassword(ctx context.Context, id, password string) error
	UpdateRole(ctx context.Context, id, role string) error
	Delete(ctx context.Context, id string) error
	Restore(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (*user.User, error)
	FindByEmail(ctx context.Context, email string) (*user.User, error)
	FindAll(ctx context.Context, offset, limit int64, isDeleted bool, search string) ([]*user.User, int64, error)
}

type userService struct {
	repo repos.User
}

func NewUserService(userRepo repos.User) UserService {
	return &userService{repo: userRepo}
}

func (s *userService) Create(ctx context.Context, rawEmail, password, role string) error {
	foundUser, err := s.repo.FindByEmail(ctx, email.Email(rawEmail))
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return err
	}
	if foundUser != nil {
		return exceptions.ErrUserWithEmailExists
	}

	u, err := user.NewUser(rawEmail, password, role)
	if err != nil {
		return err
	}
	return s.repo.Create(ctx, u)
}

func (s *userService) UpdateEmail(ctx context.Context, id, email string) error {
	convertedId, err := uuid.UUIDFromString(id)
	foundUser, err := s.repo.FindById(ctx, convertedId)
	if err != nil {
		return err
	}

	err = foundUser.UpdateEmail(email)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, foundUser)
}

func (s *userService) UpdatePassword(ctx context.Context, id, password string) error {
	convertedId, err := uuid.UUIDFromString(id)
	foundUser, err := s.repo.FindById(ctx, convertedId)
	if err != nil {
		return err
	}

	err = foundUser.UpdatePassword(password)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, foundUser)
}

func (s *userService) UpdateRole(ctx context.Context, id, role string) error {
	convertedId, err := uuid.UUIDFromString(id)
	foundUser, err := s.repo.FindById(ctx, convertedId)
	if err != nil {
		return err
	}

	err = foundUser.UpdateRole(role)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, foundUser)
}

func (s *userService) Delete(ctx context.Context, id string) error {
	convertedId, err := uuid.UUIDFromString(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, convertedId)
}

func (s *userService) Restore(ctx context.Context, id string) error {
	convertedId, err := uuid.UUIDFromString(id)
	if err != nil {
		return err
	}
	return s.repo.Restore(ctx, convertedId)
}

func (s *userService) FindById(ctx context.Context, id string) (*user.User, error) {
	convertedId, err := uuid.UUIDFromString(id)
	if err != nil {
		return nil, err
	}
	return s.repo.FindById(ctx, convertedId)
}

func (s *userService) FindByEmail(ctx context.Context, rawEmail string) (*user.User, error) {
	convertedEmail, err := email.NewEmail(rawEmail)
	if err != nil {
		return nil, err
	}
	return s.repo.FindByEmail(ctx, convertedEmail)
}

func (s *userService) FindAll(ctx context.Context, offset, limit int64, isDeleted bool, search string) ([]*user.User, int64, error) {
	return s.repo.FindAll(ctx, offset, limit, isDeleted, search)
}
