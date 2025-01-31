package bookService

import (
	"context"
	"github.com/nnniyaz/blog/internal/domain/base/uuid"
	"github.com/nnniyaz/blog/internal/domain/book"
	"github.com/nnniyaz/blog/internal/repos"
	"github.com/nnniyaz/blog/pkg/core"
)

type BookService interface {
	Create(ctx context.Context, title, description, author core.MlString, coverUri, eBookUri string) error
	Update(ctx context.Context, id string, title, description, author core.MlString, coverUri, eBookUri string) error
	Delete(ctx context.Context, id string) error
	Restore(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (*book.Book, error)
	FindAll(ctx context.Context, offset, limit int64, isDeleted bool, search string) ([]*book.Book, int64, error)
}

type bookService struct {
	repo repos.Book
}

func NewBookService(repo repos.Book) BookService {
	return &bookService{repo: repo}
}

func (s *bookService) Create(ctx context.Context, title, description, author core.MlString, coverUri, eBookUri string) error {
	newBook, err := book.NewBook(title, description, author, coverUri, eBookUri)
	if err != nil {
		return err
	}
	return s.repo.Create(ctx, newBook)
}

func (s *bookService) Update(ctx context.Context, id string, title, description, author core.MlString, coverUri, eBookUri string) error {
	convertedId, err := uuid.UUIDFromString(id)
	if err != nil {
		return err
	}

	foundBook, err := s.repo.FindById(ctx, convertedId)
	if err != nil {
		return err
	}

	err = foundBook.Update(title, description, author, coverUri, eBookUri)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, foundBook)
}

func (s *bookService) Delete(ctx context.Context, id string) error {
	convertedId, err := uuid.UUIDFromString(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, convertedId)
}

func (s *bookService) Restore(ctx context.Context, id string) error {
	convertedId, err := uuid.UUIDFromString(id)
	if err != nil {
		return err
	}
	return s.repo.Restore(ctx, convertedId)
}

func (s *bookService) FindById(ctx context.Context, id string) (*book.Book, error) {
	convertedId, err := uuid.UUIDFromString(id)
	if err != nil {
		return nil, err
	}
	return s.repo.FindById(ctx, convertedId)
}

func (s *bookService) FindAll(ctx context.Context, offset, limit int64, isDeleted bool, search string) ([]*book.Book, int64, error) {
	return s.repo.FindAll(ctx, offset, limit, isDeleted, search)
}
