package author

import (
	"github.com/nnniyaz/blog/domain/author/exceptions"
	"github.com/nnniyaz/blog/domain/base/uuid"
	"strings"
	"time"
)

type Author struct {
	id        uuid.UUID
	firstName string
	lastName  string
	avatarUri string
	createdAt time.Time
	updatedAt time.Time
	version   int
}

func NewAuthor(firstName, lastName, avatarUri string) (*Author, error) {
	cleanedFirstName := strings.TrimSpace(firstName)
	if cleanedFirstName == "" {
		return nil, exceptions.ErrFirstNameIsEmpty
	}

	cleanedLastName := strings.TrimSpace(lastName)
	if cleanedLastName == "" {
		return nil, exceptions.ErrLastNameIsEmpty
	}

	return &Author{
		id:        uuid.NewUUID(),
		firstName: cleanedFirstName,
		lastName:  cleanedLastName,
		avatarUri: avatarUri,
		createdAt: time.Now(),
		updatedAt: time.Now(),
		version:   1,
	}, nil
}

func (b *Author) GetId() uuid.UUID {
	return b.id
}

func (b *Author) GetFirstName() string {
	return b.firstName
}

func (b *Author) GetLastName() string {
	return b.lastName
}

func (b *Author) GetAvatarUri() string {
	return b.avatarUri
}

func (b *Author) GetCreatedAt() time.Time {
	return b.createdAt
}

func (b *Author) GetUpdatedAt() time.Time {
	return b.updatedAt
}

func (b *Author) GetVersion() int {
	return b.version
}

func (b *Author) Update(firstName, lastName, avatarUri string) error {
	cleanedFirstName := strings.TrimSpace(firstName)
	if cleanedFirstName == "" {
		return exceptions.ErrFirstNameIsEmpty
	}

	cleanedLastName := strings.TrimSpace(lastName)
	if cleanedLastName == "" {
		return exceptions.ErrLastNameIsEmpty
	}

	b.firstName = cleanedFirstName
	b.lastName = cleanedLastName
	b.avatarUri = avatarUri
	b.updatedAt = time.Now()
	b.version++
	return nil
}

func UnmarshalAuthorFromDatabase(id uuid.UUID, firstName, lastName, avatarUri string, createdAt, updatedAt time.Time, version int) *Author {
	return &Author{
		id:        id,
		firstName: firstName,
		lastName:  lastName,
		avatarUri: avatarUri,
		createdAt: createdAt,
		updatedAt: updatedAt,
		version:   version,
	}
}
