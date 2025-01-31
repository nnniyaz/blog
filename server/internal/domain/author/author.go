package author

import (
	"github.com/nnniyaz/blog/internal/domain/base/uuid"
	"github.com/nnniyaz/blog/pkg/core"
	"time"
)

type Author struct {
	id        uuid.UUID
	firstName core.MlString
	lastName  core.MlString
	avatarUri string
	createdAt time.Time
	updatedAt time.Time
	version   int
}

func NewAuthor(firstName, lastName core.MlString, avatarUri string) (*Author, error) {
	cleanedFirstName, err := firstName.Clean()
	if err != nil {
		return nil, err
	}

	cleanedLastName, err := lastName.Clean()
	if err != nil {
		return nil, err
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

func (b *Author) GetFirstName() core.MlString {
	return b.firstName
}

func (b *Author) GetLastName() core.MlString {
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

func (b *Author) Update(firstName, lastName core.MlString, avatarUri string) error {
	cleanedFirstName, err := firstName.Clean()
	if err != nil {
		return err
	}

	cleanedLastName, err := lastName.Clean()
	if err != nil {
		return nil
	}

	b.firstName = cleanedFirstName
	b.lastName = cleanedLastName
	b.avatarUri = avatarUri
	b.updatedAt = time.Now()
	b.version++

	return nil
}

func UnmarshalAuthorFromDatabase(id uuid.UUID, firstName, lastName core.MlString, avatarUri string, createdAt, updatedAt time.Time, version int) *Author {
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
