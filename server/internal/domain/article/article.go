package article

import (
	"github.com/nnniyaz/blog/internal/domain/article/exceptions"
	"github.com/nnniyaz/blog/internal/domain/base/uuid"
	"github.com/nnniyaz/blog/pkg/core"
	"time"
)

type Article struct {
	id        uuid.UUID
	title     core.MlString
	content   core.MlString
	isDeleted bool
	createdAt time.Time
	updatedAt time.Time
	version   int
}

func NewArticle(title, content core.MlString) (*Article, error) {
	cleanedTitle, err := title.Clean()
	if err != nil {
		return nil, err
	}

	mMap := map[core.Lang]string(cleanedTitle)
	for _, v := range mMap {
		if len(v) > 100 {
			return nil, exceptions.ErrTitleTooLong
		}
	}

	cleanedContent, err := content.Clean()
	if err != nil {
		return nil, err
	}

	return &Article{
		id:        uuid.NewUUID(),
		title:     cleanedTitle,
		content:   cleanedContent,
		isDeleted: false,
		createdAt: time.Now(),
		updatedAt: time.Now(),
		version:   1,
	}, nil
}

func (a *Article) GetId() uuid.UUID {
	return a.id
}

func (a *Article) GetTitle() core.MlString {
	return a.title
}

func (a *Article) GetContent() core.MlString {
	return a.content
}

func (a *Article) GetIsDeleted() bool {
	return a.isDeleted
}

func (a *Article) GetCreatedAt() time.Time {
	return a.createdAt
}

func (a *Article) GetUpdatedAt() time.Time {
	return a.updatedAt
}

func (a *Article) GetVersion() int {
	return a.version
}

func (a *Article) Update(title, content core.MlString) error {
	cleanedTitle, err := title.Clean()
	if err != nil {
		return err
	}

	mMap := map[core.Lang]string(cleanedTitle)
	for _, v := range mMap {
		if len(v) > 100 {
			return exceptions.ErrTitleTooLong
		}
	}

	cleanedContent, err := content.Clean()
	if err != nil {
		return err
	}

	a.title = cleanedTitle
	a.content = cleanedContent
	a.updatedAt = time.Now()
	a.version++

	return nil
}

func UnmarshalArticleFromDatabase(id uuid.UUID, title, content core.MlString, isDeleted bool, createdAt, updatedAt time.Time, version int) *Article {
	return &Article{
		id:        id,
		title:     title,
		content:   content,
		isDeleted: isDeleted,
		createdAt: createdAt,
		updatedAt: updatedAt,
		version:   version,
	}
}
