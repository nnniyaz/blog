package service

import (
	"github.com/nnniyaz/blog/domain/base/config"
	"github.com/nnniyaz/blog/pkg/email"
	"github.com/nnniyaz/blog/pkg/logger"
	"github.com/nnniyaz/blog/repo"
	articleService "github.com/nnniyaz/blog/service/article"
)

type Service struct {
	Article articleService.ApplicationService
}

func NewService(repos *repo.Repo, config *config.Config, l logger.Logger, emailService email.Email) *Service {
	return &Service{
		Article: articleService.NewArticleService(l, repos.RepoArticle),
	}
}
