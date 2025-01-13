package repo

import (
	"context"
	"github.com/nnniyaz/blog/domain/article"
	"github.com/nnniyaz/blog/domain/base/uuid"
	mongoArticle "github.com/nnniyaz/blog/repo/mongo/article"
	"go.mongodb.org/mongo-driver/mongo"
)

type Article interface {
	Create(ctx context.Context, article *article.Article) error
	Update(ctx context.Context, article *article.Article) error
	Delete(ctx context.Context, id uuid.UUID) error
	Restore(ctx context.Context, id uuid.UUID) error
	FindById(ctx context.Context, id uuid.UUID) (*article.Article, error)
	FindAll(ctx context.Context, offset, limit int64, isDeleted bool, title string) ([]*article.Article, int64, error)
}

type Repo struct {
	RepoArticle Article
}

func NewRepo(client *mongo.Client) *Repo {
	return &Repo{
		RepoArticle: mongoArticle.NewRepoArticle(client),
	}
}
