package services

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/nnniyaz/blog/internal/domain/base/config"
	"github.com/nnniyaz/blog/internal/repos"
	"github.com/nnniyaz/blog/internal/services/article"
	authService "github.com/nnniyaz/blog/internal/services/auth"
	"github.com/nnniyaz/blog/internal/services/author"
	"github.com/nnniyaz/blog/internal/services/bio"
	"github.com/nnniyaz/blog/internal/services/book"
	"github.com/nnniyaz/blog/internal/services/contact"
	"github.com/nnniyaz/blog/internal/services/project"
	sessionService "github.com/nnniyaz/blog/internal/services/session"
	uploadService "github.com/nnniyaz/blog/internal/services/upload"
	userService "github.com/nnniyaz/blog/internal/services/user"
	"github.com/nnniyaz/blog/pkg/email"
)

type Service struct {
	Article articleService.ArticleService
	Contact contactService.ContactService
	Author  authorService.AuthorService
	Bio     bioService.BioService
	Book    bookService.BookService
	Project projectService.ProjectService
	User    userService.UserService
	Session sessionService.SessionService
	Auth    authService.AuthService
	Upload  uploadService.UploadService
}

func NewService(repos *repos.Repo, config *config.Config, s3 *s3.S3, emailService email.Email) *Service {
	session := sessionService.NewSessionService(repos.RepoSession)
	return &Service{
		Article: articleService.NewArticleService(repos.RepoArticle),
		Contact: contactService.NewContactService(repos.RepoContact),
		Author:  authorService.NewAuthorService(repos.RepoAuthor),
		Bio:     bioService.NewBioService(repos.RepoBio),
		Book:    bookService.NewBookService(repos.RepoBook),
		Project: projectService.NewProjectService(repos.RepoProject),
		User:    userService.NewUserService(repos.RepoUser),
		Auth:    authService.NewAuthService(repos.RepoUser, session),
		Upload:  uploadService.NewUploadService(s3, config.GetSpaceBucket()),
	}
}
