package repo

import (
	"context"
	"github.com/nnniyaz/blog/domain/article"
	"github.com/nnniyaz/blog/domain/author"
	"github.com/nnniyaz/blog/domain/base/uuid"
	"github.com/nnniyaz/blog/domain/bio"
	"github.com/nnniyaz/blog/domain/book"
	"github.com/nnniyaz/blog/domain/contact"
	"github.com/nnniyaz/blog/domain/project"
	mongoArticle "github.com/nnniyaz/blog/repo/mongo/article"
	mongoAuthor "github.com/nnniyaz/blog/repo/mongo/author"
	mongoBio "github.com/nnniyaz/blog/repo/mongo/bio"
	mongoBook "github.com/nnniyaz/blog/repo/mongo/book"
	mongoContact "github.com/nnniyaz/blog/repo/mongo/contact"
	mongoProject "github.com/nnniyaz/blog/repo/mongo/project"
	"go.mongodb.org/mongo-driver/mongo"
)

type Article interface {
	Create(ctx context.Context, article *article.Article) error
	Update(ctx context.Context, article *article.Article) error
	Delete(ctx context.Context, id uuid.UUID) error
	Restore(ctx context.Context, id uuid.UUID) error
	FindById(ctx context.Context, id uuid.UUID) (*article.Article, error)
	FindAll(ctx context.Context, offset, limit int64, isDeleted bool, search string) ([]*article.Article, int64, error)
}

type Contact interface {
	Create(ctx context.Context, contact *contact.Contact) error
	Update(ctx context.Context, contact *contact.Contact) error
	Delete(ctx context.Context, id uuid.UUID) error
	Restore(ctx context.Context, id uuid.UUID) error
	FindById(ctx context.Context, id uuid.UUID) (*contact.Contact, error)
	FindAll(ctx context.Context, offset, limit int64, isDeleted bool, search string) ([]*contact.Contact, int64, error)
}

type Author interface {
	Create(ctx context.Context, author *author.Author) error
	Update(ctx context.Context, author *author.Author) error
	Delete(ctx context.Context, id uuid.UUID) error
	Restore(ctx context.Context, id uuid.UUID) error
	FindById(ctx context.Context, id uuid.UUID) (*author.Author, error)
	FindAll(ctx context.Context) ([]*author.Author, int64, error)
}

type Bio interface {
	Create(ctx context.Context, bio *bio.Bio) error
	Update(ctx context.Context, bio *bio.Bio) error
	Delete(ctx context.Context, id uuid.UUID) error
	Restore(ctx context.Context, id uuid.UUID) error
	SetActive(ctx context.Context, id uuid.UUID) error
	FindById(ctx context.Context, id uuid.UUID) (*bio.Bio, error)
	FindByActive(ctx context.Context) (*bio.Bio, error)
	FindAll(ctx context.Context) ([]*bio.Bio, int64, error)
}

type Book interface {
	Create(ctx context.Context, book *book.Book) error
	Update(ctx context.Context, book *book.Book) error
	Delete(ctx context.Context, id uuid.UUID) error
	Restore(ctx context.Context, id uuid.UUID) error
	FindById(ctx context.Context, id uuid.UUID) (*book.Book, error)
	FindAll(ctx context.Context, offset, limit int64, isDeleted bool, search string) ([]*book.Book, int64, error)
}

type Project interface {
	Create(ctx context.Context, project *project.Project) error
	Update(ctx context.Context, project *project.Project) error
	Delete(ctx context.Context, id uuid.UUID) error
	Restore(ctx context.Context, id uuid.UUID) error
	FindById(ctx context.Context, id uuid.UUID) (*project.Project, error)
	FindAll(ctx context.Context, offset, limit int64, isDeleted bool, search string) ([]*project.Project, int64, error)
}

type Repo struct {
	RepoArticle Article
	RepoContact Contact
	RepoAuthor  Author
	RepoBio     Bio
	RepoBook    Book
	RepoProject Project
}

func NewRepo(client *mongo.Client) *Repo {
	return &Repo{
		RepoArticle: mongoArticle.NewRepoArticle(client),
		RepoContact: mongoContact.NewRepoContact(client),
		RepoAuthor:  mongoAuthor.NewRepoAuthor(client),
		RepoBio:     mongoBio.NewRepoBio(client),
		RepoBook:    mongoBook.NewBookRepo(client),
		RepoProject: mongoProject.NewProjectRepo(client),
	}
}
