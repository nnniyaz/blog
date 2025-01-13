package article

import (
	"github.com/nnniyaz/blog/domain/article/exceptions"
	"github.com/nnniyaz/blog/domain/base/uuid"
	"strings"
	"time"
)

type Article struct {
	id        uuid.UUID
	title     string
	content   string
	isDeleted bool
	createdAt time.Time
	updatedAt time.Time
	version   int
}

func NewArticle(title, content string) (*Article, error) {
	cleanedTitle := strings.TrimSpace(title)
	if cleanedTitle == "" {
		return nil, exceptions.ErrTitleIsEmpty
	}

	if len(cleanedTitle) > 100 {
		return nil, exceptions.ErrTitleIsTooLong
	}

	cleanedContent := strings.TrimSpace(content)
	if cleanedContent == "" {
		return nil, exceptions.ErrContentIsEmpty
	}

	return &Article{
		id:        uuid.NewUUID(),
		title:     strings.TrimSpace(cleanedTitle),
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

func (a *Article) GetTitle() string {
	return a.title
}

func (a *Article) GetContent() string {
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

func (a *Article) Update(title, content string) error {
	cleanedTitle := strings.TrimSpace(title)
	if cleanedTitle == "" {
		return exceptions.ErrTitleIsEmpty
	}

	if len(cleanedTitle) > 100 {
		return exceptions.ErrTitleIsTooLong
	}

	cleanedContent := strings.TrimSpace(content)
	if cleanedContent == "" {
		return exceptions.ErrContentIsEmpty
	}

	a.title = cleanedTitle
	a.content = cleanedContent
	a.updatedAt = time.Now()
	a.version++

	return nil
}

func UnmarshalArticleFromDatabase(id uuid.UUID, title, content string, isDeleted bool, createdAt, updatedAt time.Time, version int) *Article {
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
