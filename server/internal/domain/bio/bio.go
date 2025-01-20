package bio

import (
	"github.com/nnniyaz/blog/internal/domain/base/uuid"
	"github.com/nnniyaz/blog/internal/domain/bio/exceptions"
	"strings"
	"time"
)

type Bio struct {
	id        uuid.UUID
	bio       string
	active    bool
	createdAt time.Time
	updatedAt time.Time
	version   int
}

func NewBio(bio string) (*Bio, error) {
	cleanedEmptyBio := strings.TrimSpace(bio)
	if cleanedEmptyBio == "" {
		return nil, exceptions.ErrBioEmpty
	}

	return &Bio{
		id:        uuid.NewUUID(),
		bio:       cleanedEmptyBio,
		active:    true,
		createdAt: time.Now(),
		updatedAt: time.Now(),
		version:   1,
	}, nil
}

func (b *Bio) GetId() uuid.UUID {
	return b.id
}

func (b *Bio) GetBio() string {
	return b.bio
}

func (b *Bio) GetActive() bool {
	return b.active
}

func (b *Bio) GetCreatedAt() time.Time {
	return b.createdAt
}

func (b *Bio) GetUpdatedAt() time.Time {
	return b.updatedAt
}

func (b *Bio) GetVersion() int {
	return b.version
}

func (b *Bio) Update(bio string) error {
	cleanedEmptyBio := strings.TrimSpace(bio)
	if cleanedEmptyBio == "" {
		return exceptions.ErrBioEmpty
	}

	b.bio = cleanedEmptyBio
	b.updatedAt = time.Now()
	b.version++

	return nil
}

func UnmarshalBioFromDatabase(id uuid.UUID, bio string, active bool, createdAt, updatedAt time.Time, version int) *Bio {
	return &Bio{
		id:        id,
		bio:       bio,
		active:    active,
		createdAt: createdAt,
		updatedAt: updatedAt,
		version:   version,
	}
}
