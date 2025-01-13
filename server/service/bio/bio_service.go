package bioService

import (
	"context"
	"github.com/nnniyaz/blog/domain/base/uuid"
	"github.com/nnniyaz/blog/domain/bio"
	"github.com/nnniyaz/blog/repo"
)

type BioService interface {
	Create(ctx context.Context, bio string) error
	Update(ctx context.Context, id, bio string) error
	Delete(ctx context.Context, id string) error
	Restore(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (*bio.Bio, error)
	FindAll(ctx context.Context) ([]*bio.Bio, error)
}

type bioService struct {
	repo repo.Bio
}

func NewBioService(repo repo.Bio) BioService {
	return &bioService{repo: repo}
}

func (s *bioService) Create(ctx context.Context, bioContent string) error {
	newBio, err := bio.NewBio(bioContent)
	if err != nil {
		return err
	}
	return s.repo.Create(ctx, newBio)
}

func (s *bioService) Update(ctx context.Context, id, bioContent string) error {
	convertedId, err := uuid.UUIDFromString(id)
	if err != nil {
		return err
	}

	bio, err := s.repo.FindById(ctx, convertedId)
	if err != nil {
		return err
	}

	err = bio.Update(bioContent)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, bio)
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

func (s *bioService) FindAll(ctx context.Context) ([]*bio.Bio, error) {
	return s.repo.FindAll(ctx)
}
