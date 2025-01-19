package repos

import (
	"context"
	"github.com/nnniyaz/blog/internal/domain/article"
	"github.com/nnniyaz/blog/internal/domain/author"
	"github.com/nnniyaz/blog/internal/domain/base/uuid"
	"github.com/nnniyaz/blog/internal/domain/bio"
	"github.com/nnniyaz/blog/internal/domain/book"
	"github.com/nnniyaz/blog/internal/domain/contact"
	"github.com/nnniyaz/blog/internal/domain/project"
	"github.com/nnniyaz/blog/internal/domain/session"
	"github.com/nnniyaz/blog/internal/domain/user"
	mongoArticle "github.com/nnniyaz/blog/internal/repos/mongo/article"
	mongoAuthor "github.com/nnniyaz/blog/internal/repos/mongo/author"
	mongoBio "github.com/nnniyaz/blog/internal/repos/mongo/bio"
	mongoBook "github.com/nnniyaz/blog/internal/repos/mongo/book"
	mongoContact "github.com/nnniyaz/blog/internal/repos/mongo/contact"
	mongoProject "github.com/nnniyaz/blog/internal/repos/mongo/project"
	mongoSession "github.com/nnniyaz/blog/internal/repos/mongo/session"
	mongoUser "github.com/nnniyaz/blog/internal/repos/mongo/user"
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

type User interface {
	Create(ctx context.Context, u *user.User) error
	Update(ctx context.Context, u *user.User) error
	Delete(ctx context.Context, id uuid.UUID) error
	Restore(ctx context.Context, id uuid.UUID) error
	FindById(ctx context.Context, id uuid.UUID) (*user.User, error)
	FindAll(ctx context.Context, offset, limit int64, isDeleted bool, search string) ([]*user.User, int64, error)
}

type Session interface {
	Create(ctx context.Context, s *session.Session) error
	DeleteBySession(ctx context.Context, session uuid.UUID) error
	DeleteByUserId(ctx context.Context, userId uuid.UUID) error
	FindAll(ctx context.Context, offset, limit int64) ([]*session.Session, int64, error)
	FindBySession(ctx context.Context, session uuid.UUID) (*session.Session, error)
	FindByUserId(ctx context.Context, userId uuid.UUID) ([]*session.Session, error)
}

type Repo struct {
	RepoArticle Article
	RepoContact Contact
	RepoAuthor  Author
	RepoBio     Bio
	RepoBook    Book
	RepoProject Project
	RepoUser    User
	RepoSession Session
}

func NewRepo(client *mongo.Client) *Repo {
	return &Repo{
		RepoArticle: mongoArticle.NewRepoArticle(client),
		RepoContact: mongoContact.NewRepoContact(client),
		RepoAuthor:  mongoAuthor.NewRepoAuthor(client),
		RepoBio:     mongoBio.NewRepoBio(client),
		RepoBook:    mongoBook.NewBookRepo(client),
		RepoProject: mongoProject.NewProjectRepo(client),
		RepoUser:    mongoUser.NewUserRepo(client),
		RepoSession: mongoSession.NewRepoSession(client),
	}
}
