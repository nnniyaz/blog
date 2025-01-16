package project

import (
	"github.com/nnniyaz/blog/internal/domain/base/uuid"
	"github.com/nnniyaz/blog/internal/domain/project/exceptions"
	"strings"
	"time"
)

type Project struct {
	id             uuid.UUID
	name           string
	description    string
	coverUri       string
	appLink        string
	sourceCodeLink string
	isDeleted      bool
	createdAt      time.Time
	updatedAt      time.Time
}

func NewProject(name, description, coverUri, appLink, sourceCodeLink string) (*Project, error) {
	cleanedName := name
	if cleanedName == "" {
		return nil, exceptions.ErrProjectNameIsEmpty
	}

	return &Project{
		id:             uuid.NewUUID(),
		name:           cleanedName,
		description:    strings.TrimSpace(description),
		coverUri:       coverUri,
		appLink:        appLink,
		sourceCodeLink: sourceCodeLink,
		isDeleted:      false,
		createdAt:      time.Now(),
		updatedAt:      time.Now(),
	}, nil
}

func (p *Project) GetId() uuid.UUID {
	return p.id
}

func (p *Project) GetName() string {
	return p.name
}

func (p *Project) GetDescription() string {
	return p.description
}

func (p *Project) GetCoverUri() string {
	return p.coverUri
}

func (p *Project) GetAppLink() string {
	return p.appLink
}

func (p *Project) GetSourceCodeLink() string {
	return p.sourceCodeLink
}

func (p *Project) GetIsDeleted() bool {
	return p.isDeleted
}

func (p *Project) GetCreatedAt() time.Time {
	return p.createdAt
}

func (p *Project) GetUpdatedAt() time.Time {
	return p.updatedAt
}

func (p *Project) Update(name, description, coverUri, appLink, sourceCodeLink string) error {
	cleanedName := name
	if cleanedName == "" {
		return exceptions.ErrProjectNameIsEmpty
	}

	p.name = cleanedName
	p.description = strings.TrimSpace(description)
	p.coverUri = coverUri
	p.appLink = appLink
	p.sourceCodeLink = sourceCodeLink
	p.updatedAt = time.Now()
	return nil
}

func UnmarshalProjectFromDatabase(id uuid.UUID, name, description, coverUri, appLink, sourceCodeLink string, isDeleted bool, createdAt, updatedAt time.Time) *Project {
	return &Project{
		id:             id,
		name:           name,
		description:    description,
		coverUri:       coverUri,
		appLink:        appLink,
		sourceCodeLink: sourceCodeLink,
		isDeleted:      isDeleted,
		createdAt:      createdAt,
		updatedAt:      updatedAt,
	}
}
