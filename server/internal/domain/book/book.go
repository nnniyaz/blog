package book

import (
	"github.com/nnniyaz/blog/internal/domain/base/uuid"
	"github.com/nnniyaz/blog/pkg/core"
	"time"
)

type Book struct {
	id          uuid.UUID
	title       core.MlString
	author      core.MlString
	description core.MlString
	coverUri    string
	eBookUri    string
	isDeleted   bool
	createdAt   time.Time
	updatedAt   time.Time
}

func NewBook(title, author, description core.MlString, coverUri, eBookUri string) (*Book, error) {
	cleanedTitle, err := title.Clean()
	if err != nil {
		return nil, err
	}

	cleanedAuthor, err := author.Clean()
	if err != nil {
		return nil, err
	}

	return &Book{
		id:          uuid.NewUUID(),
		title:       cleanedTitle,
		author:      cleanedAuthor,
		description: description,
		coverUri:    coverUri,
		eBookUri:    eBookUri,
		isDeleted:   false,
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
	}, nil
}

func (b *Book) GetId() uuid.UUID {
	return b.id
}

func (b *Book) GetTitle() core.MlString {
	return b.title
}

func (b *Book) GetAuthor() core.MlString {
	return b.author
}

func (b *Book) GetDescription() core.MlString {
	return b.description
}

func (b *Book) GetCoverUri() string {
	return b.coverUri
}

func (b *Book) GetEBookUri() string {
	return b.eBookUri
}

func (b *Book) GetIsDeleted() bool {
	return b.isDeleted
}

func (b *Book) GetCreatedAt() time.Time {
	return b.createdAt
}

func (b *Book) GetUpdatedAt() time.Time {
	return b.updatedAt
}

func (b *Book) Update(title, author, description core.MlString, coverUri, eBookUri string) error {
	cleanedTitle, err := title.Clean()
	if err != nil {
		return err
	}
	cleanedAuthor, err := author.Clean()
	if err != nil {
		return err
	}

	b.title = cleanedTitle
	b.author = cleanedAuthor
	b.description = description
	b.coverUri = coverUri
	b.eBookUri = eBookUri
	b.updatedAt = time.Now()

	return nil
}

func UnmarshalBookFromDatabase(id uuid.UUID, title, author, description core.MlString, coverUri, eBookUri string, isDeleted bool, createdAt, updatedAt time.Time) *Book {
	return &Book{
		id:          id,
		title:       title,
		author:      author,
		description: description,
		coverUri:    coverUri,
		eBookUri:    eBookUri,
		isDeleted:   isDeleted,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}
}
