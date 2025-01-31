package bioService

import (
	"context"
	"github.com/nnniyaz/blog/internal/domain/base/uuid"
	"github.com/nnniyaz/blog/internal/domain/bio"
	"github.com/nnniyaz/blog/internal/repos"
	"github.com/nnniyaz/blog/pkg/core"
)

type BioService interface {
	Create(ctx context.Context, bio core.MlString) error
	Update(ctx context.Context, id string, bio core.MlString) error
	Delete(ctx context.Context, id string) error
	Restore(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (*bio.Bio, error)
	FindByActive(ctx context.Context) (*bio.Bio, error)
	FindAll(ctx context.Context, offset, limit int64) ([]*bio.Bio, int64, error)
}

type bioService struct {
	repo repos.Bio
}

func NewBioService(repo repos.Bio) BioService {
	return &bioService{repo: repo}
}

func (s *bioService) Create(ctx context.Context, bioContent core.MlString) error {
	newBio, err := bio.NewBio(bioContent)
	if err != nil {
		return err
	}
	return s.repo.Create(ctx, newBio)
}

func (s *bioService) Update(ctx context.Context, id string, bioContent core.MlString) error {
	convertedId, err := uuid.UUIDFromString(id)
	if err != nil {
		return err
	}

	foundBio, err := s.repo.FindById(ctx, convertedId)
	if err != nil {
		return err
	}

	err = foundBio.Update(bioContent)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, foundBio)
}

func (s *bioService) Delete(ctx context.Context, id string) error {
	convertedId, err := uuid.UUIDFromString(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, convertedId)
}

func (s *bioService) Restore(ctx context.Context, id string) error {
	convertedId, err := uuid.UUIDFromString(id)
	if err != nil {
		return err
	}
	return s.repo.Restore(ctx, convertedId)
}

func (s *bioService) SetActive(ctx context.Context, id string) error {
	convertedId, err := uuid.UUIDFromString(id)
	if err != nil {
		return err
	}
	return s.repo.SetActive(ctx, convertedId)
}

func (s *bioService) FindById(ctx context.Context, id string) (*bio.Bio, error) {
	convertedId, err := uuid.UUIDFromString(id)
	if err != nil {
		return nil, err
	}
	return s.repo.FindById(ctx, convertedId)
}

func (s *bioService) FindByActive(ctx context.Context) (*bio.Bio, error) {
	return s.repo.FindByActive(ctx)
}

func (s *bioService) FindAll(ctx context.Context, offset, limit int64) ([]*bio.Bio, int64, error) {
	return s.repo.FindAll(ctx, offset, limit)
}
