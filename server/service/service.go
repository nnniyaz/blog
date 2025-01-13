package service

import (
	"github.com/nnniyaz/blog/domain/base/config"
	"github.com/nnniyaz/blog/pkg/email"
	"github.com/nnniyaz/blog/repo"
	articleService "github.com/nnniyaz/blog/service/article"
	authorService "github.com/nnniyaz/blog/service/author"
	bioService "github.com/nnniyaz/blog/service/bio"
	bookService "github.com/nnniyaz/blog/service/book"
	contactService "github.com/nnniyaz/blog/service/contact"
)

type Service struct {
	Article articleService.ApplicationService
	Contact contactService.ContactService
	Author  authorService.AuthorService
	Bio     bioService.BioService
	Book    bookService.BookService
}

func NewService(repos *repo.Repo, config *config.Config, emailService email.Email) *Service {
	return &Service{
		Article: articleService.NewArticleService(repos.RepoArticle),
		Contact: contactService.NewContactService(repos.RepoContact),
		Author:  authorService.NewAuthorService(repos.RepoAuthor),
		Bio:     bioService.NewBioService(repos.RepoBio),
		Book:    bookService.NewBookService(repos.RepoBook),
	}
}
