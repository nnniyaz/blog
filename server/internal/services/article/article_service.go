package articleService

import (
	"context"
	"github.com/nnniyaz/blog/internal/domain/article"
	"github.com/nnniyaz/blog/internal/domain/base/uuid"
	"github.com/nnniyaz/blog/internal/repos"
	"github.com/nnniyaz/blog/pkg/core"
)

type ArticleService interface {
	Create(ctx context.Context, title, content core.MlString) error
	Update(ctx context.Context, id string, title, content core.MlString) error
	Delete(ctx context.Context, id string) error
	Restore(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (*article.Article, error)
	FindAll(ctx context.Context, offset, limit int64, isDeleted bool, search string) ([]*article.Article, int64, error)
}

type articleService struct {
	repo repos.Article
}

func NewArticleService(r repos.Article) ArticleService {
	return &articleService{repo: r}
}

func (s *articleService) Create(ctx context.Context, title, content core.MlString) error {
	newArticle, err := article.NewArticle(title, content)
	if err != nil {
		return err
	}
	return s.repo.Create(ctx, newArticle)
}

func (s *articleService) Update(ctx context.Context, id string, title, content core.MlString) error {
	convertedId, err := uuid.UUIDFromString(id)
	if err != nil {
		return err
	}

	foundArticle, err := s.repo.FindById(ctx, convertedId)
	if err != nil {
		return err
	}

	err = foundArticle.Update(title, content)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, foundArticle)
}

func (s *articleService) Delete(ctx context.Context, id string) error {
	convertedId, err := uuid.UUIDFromString(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, convertedId)
}

func (s *articleService) Restore(ctx context.Context, id string) error {
	convertedId, err := uuid.UUIDFromString(id)
	if err != nil {
		return err
	}
	return s.repo.Restore(ctx, convertedId)
}

func (s *articleService) FindById(ctx context.Context, id string) (*article.Article, error) {
	convertedId, err := uuid.UUIDFromString(id)
	if err != nil {
		return nil, err
	}
	return s.repo.FindById(ctx, convertedId)
}

func (s *articleService) FindAll(ctx context.Context, offset, limit int64, isDeleted bool, search string) ([]*article.Article, int64, error) {
	return s.repo.FindAll(ctx, offset, limit, isDeleted, search)
}
