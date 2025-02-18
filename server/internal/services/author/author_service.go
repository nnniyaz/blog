package authorService

import (
	"context"
	"github.com/nnniyaz/blog/internal/domain/author"
	"github.com/nnniyaz/blog/internal/domain/base/uuid"
	"github.com/nnniyaz/blog/internal/repos"
	"github.com/nnniyaz/blog/pkg/core"
)

type AuthorService interface {
	Create(ctx context.Context, firstName, lastName core.MlString, avatarUri string) error
	Update(ctx context.Context, id string, firstName, lastName core.MlString, avatarUri string) error
	Delete(ctx context.Context, id string) error
	Restore(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (*author.Author, error)
	FindAll(ctx context.Context, offset, limit int64) ([]*author.Author, int64, error)
}

type authorService struct {
	repo repos.Author
}

func NewAuthorService(repo repos.Author) AuthorService {
	return &authorService{repo: repo}
}

func (s *authorService) Create(ctx context.Context, firstName, lastName core.MlString, avatarUri string) error {
	newAuthor, err := author.NewAuthor(firstName, lastName, avatarUri)
	if err != nil {
		return err
	}
	return s.repo.Create(ctx, newAuthor)
}

func (s *authorService) Update(ctx context.Context, id string, firstName, lastName core.MlString, avatarUri string) error {
	convertedId, err := uuid.UUIDFromString(id)
	if err != nil {
		return err
	}

	foundAuthor, err := s.repo.FindById(ctx, convertedId)
	if err != nil {
		return err
	}

	err = foundAuthor.Update(firstName, lastName, avatarUri)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, foundAuthor)
}

func (s *authorService) Delete(ctx context.Context, id string) error {
	convertedId, err := uuid.UUIDFromString(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, convertedId)
}

func (s *authorService) Restore(ctx context.Context, id string) error {
	convertedId, err := uuid.UUIDFromString(id)
	if err != nil {
		return err
	}
	return s.repo.Restore(ctx, convertedId)
}

func (s *authorService) FindById(ctx context.Context, id string) (*author.Author, error) {
	convertedId, err := uuid.UUIDFromString(id)
	if err != nil {
		return nil, err
	}
	return s.repo.FindById(ctx, convertedId)
}

func (s *authorService) FindAll(ctx context.Context, offset, limit int64) ([]*author.Author, int64, error) {
	return s.repo.FindAll(ctx, offset, limit)
}
