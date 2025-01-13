package authorService

import (
	"context"
	"github.com/nnniyaz/blog/domain/author"
	"github.com/nnniyaz/blog/domain/base/uuid"
	"github.com/nnniyaz/blog/repo"
)

type AuthorService interface {
	Create(ctx context.Context, firstName, lastName, avatarUri string) error
	Update(ctx context.Context, id, firstName, lastName, avatarUri string) error
	Delete(ctx context.Context, id string) error
	Restore(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (*author.Author, error)
	FindAll(ctx context.Context) ([]*author.Author, int64, error)
}

type authorService struct {
	repo repo.Author
}

func NewAuthorService(repo repo.Author) AuthorService {
	return &authorService{repo: repo}
}

func (s *authorService) Create(ctx context.Context, firstName, lastName, avatarUri string) error {
	newAuthor, err := author.NewAuthor(firstName, lastName, avatarUri)
	if err != nil {
		return err
	}
	return s.repo.Create(ctx, newAuthor)
}

func (s *authorService) Update(ctx context.Context, id, firstName, lastName, avatarUri string) error {
	convertedId, err := uuid.UUIDFromString(id)
	if err != nil {
		return err
	}

	author, err := s.repo.FindById(ctx, convertedId)
	if err != nil {
		return err
	}

	err = author.Update(firstName, lastName, avatarUri)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, author)
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

func (s *authorService) FindAll(ctx context.Context) ([]*author.Author, int64, error) {
	return s.repo.FindAll(ctx)
}
