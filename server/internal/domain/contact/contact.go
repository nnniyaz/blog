package contact

import (
	"github.com/nnniyaz/blog/internal/domain/base/uuid"
	"github.com/nnniyaz/blog/pkg/core"
	"strings"
	"time"
)

type Contact struct {
	id        uuid.UUID
	label     core.MlString
	link      string
	isDeleted bool
	createdAt time.Time
	updatedAt time.Time
}

func NewContact(label core.MlString, link string) (*Contact, error) {
	cleanedLabel, err := label.Clean()
	if err != nil {
		return nil, err
	}

	return &Contact{
		id:        uuid.NewUUID(),
		label:     cleanedLabel,
		link:      strings.TrimSpace(link),
		isDeleted: false,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}

func (c *Contact) GetId() uuid.UUID {
	return c.id
}

func (c *Contact) GetLabel() core.MlString {
	return c.label
}

func (c *Contact) GetLink() string {
	return c.link
}

func (c *Contact) GetIsDeleted() bool {
	return c.isDeleted
}

func (c *Contact) GetCreatedAt() time.Time {
	return c.createdAt
}

func (c *Contact) GetUpdatedAt() time.Time {
	return c.updatedAt
}

func (c *Contact) Update(label core.MlString, link string) error {
	cleanedLabel, err := label.Clean()
	if err != nil {
		return err
	}

	c.label = cleanedLabel
	c.link = strings.TrimSpace(link)
	c.updatedAt = time.Now()

	return nil
}

func UnmarshalContactFromDatabase(id uuid.UUID, label core.MlString, link string, isDeleted bool, createdAt, updatedAt time.Time) *Contact {
	return &Contact{
		id:        id,
		label:     label,
		link:      link,
		isDeleted: isDeleted,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}
