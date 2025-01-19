package projectService

import (
	"context"
	"github.com/nnniyaz/blog/internal/domain/base/uuid"
	"github.com/nnniyaz/blog/internal/domain/project"
	"github.com/nnniyaz/blog/internal/repos"
)

type ProjectService interface {
	Create(ctx context.Context, name, description, coverUri, appLink, sourceCodeLink string) error
	Update(ctx context.Context, id, name, description, coverUri, appLink, sourceCodeLink string) error
	Delete(ctx context.Context, id string) error
	Restore(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (*project.Project, error)
	FindAll(ctx context.Context, offset, limit int64, isDeleted bool, search string) ([]*project.Project, int64, error)
}

type projectService struct {
	repo repos.Project
}

func NewProjectService(projectRepo repos.Project) ProjectService {
	return &projectService{repo: projectRepo}
}

func (s *projectService) Create(ctx context.Context, name, description, coverUri, appLink, sourceCodeLink string) error {
	project, err := project.NewProject(name, description, coverUri, appLink, sourceCodeLink)
	if err != nil {
		return err
	}
	return s.repo.Create(ctx, project)
}

func (s *projectService) Update(ctx context.Context, id, name, description, coverUri, appLink, sourceCodeLink string) error {
	convertedId, err := uuid.UUIDFromString(id)
	if err != nil {
		return err
	}

	foundProject, err := s.repo.FindById(ctx, convertedId)
	if err != nil {
		return err
	}

	err = foundProject.Update(name, description, coverUri, appLink, sourceCodeLink)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, foundProject)
}

func (s *projectService) Delete(ctx context.Context, id string) error {
	convertedId, err := uuid.UUIDFromString(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, convertedId)
}

func (s *projectService) Restore(ctx context.Context, id string) error {
	convertedId, err := uuid.UUIDFromString(id)
	if err != nil {
		return err
	}
	return s.repo.Restore(ctx, convertedId)
}

func (s *projectService) FindById(ctx context.Context, id string) (*project.Project, error) {
	convertedId, err := uuid.UUIDFromString(id)
	if err != nil {
		return nil, err
	}
	return s.repo.FindById(ctx, convertedId)
}

func (s *projectService) FindAll(ctx context.Context, offset, limit int64, isDeleted bool, search string) ([]*project.Project, int64, error) {
	return s.repo.FindAll(ctx, offset, limit, isDeleted, search)
}
