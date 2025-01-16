package contactService

import (
	"context"
	"github.com/nnniyaz/blog/internal/domain/base/uuid"
	"github.com/nnniyaz/blog/internal/domain/contact"
	"github.com/nnniyaz/blog/internal/repos"
)

type ContactService interface {
	Create(ctx context.Context, label, link string) error
	Update(ctx context.Context, id string, label, link string) error
	Delete(ctx context.Context, id string) error
	Restore(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (*contact.Contact, error)
	FindAll(ctx context.Context, offset, limit int64, isDeleted bool, search string) ([]*contact.Contact, int64, error)
}

type contactService struct {
	repo repos.Contact
}

func NewContactService(r repos.Contact) ContactService {
	return &contactService{repo: r}
}

func (s *contactService) Create(ctx context.Context, label, link string) error {
	newContact, err := contact.NewContact(label, link)
	if err != nil {
		return err
	}
	return s.repo.Create(ctx, newContact)
}

func (s *contactService) Update(ctx context.Context, id, label, link string) error {
	convertedId, err := uuid.UUIDFromString(id)
	if err != nil {
		return err
	}
	foundContact, err := s.repo.FindById(ctx, convertedId)
	if err != nil {
		return err
	}
	err = foundContact.Update(label, link)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, foundContact)
}

func (s *contactService) Delete(ctx context.Context, id string) error {
	convertedId, err := uuid.UUIDFromString(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, convertedId)
}

func (s *contactService) Restore(ctx context.Context, id string) error {
	convertedId, err := uuid.UUIDFromString(id)
	if err != nil {
		return err
	}
	return s.repo.Restore(ctx, convertedId)
}

func (s *contactService) FindById(ctx context.Context, id string) (*contact.Contact, error) {
	convertedId, err := uuid.UUIDFromString(id)
	if err != nil {
		return nil, err
	}
	return s.repo.FindById(ctx, convertedId)
}

func (s *contactService) FindAll(ctx context.Context, offset, limit int64, isDeleted bool, search string) ([]*contact.Contact, int64, error) {
	return s.repo.FindAll(ctx, offset, limit, isDeleted, search)
}
