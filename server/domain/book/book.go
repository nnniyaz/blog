package book

import (
	"github.com/nnniyaz/blog/domain/base/uuid"
	"github.com/nnniyaz/blog/domain/book/exceptions"
	"strings"
	"time"
)

type Book struct {
	id          uuid.UUID
	title       string
	author      string
	description string
	coverUri    string
	eBookUri    string
	isDeleted   bool
	createdAt   time.Time
	updatedAt   time.Time
}

func NewBook(title, author, description, coverUri, eBookUri string) (*Book, error) {
	cleanedTitle := strings.TrimSpace(title)
	if cleanedTitle == "" {
		return nil, exceptions.ErrBookTitleIsEmpty
	}

	cleanedAuthor := strings.TrimSpace(author)
	if cleanedAuthor == "" {
		return nil, exceptions.ErrBookAuthorIsEmpty
	}

	return &Book{
		id:          uuid.NewUUID(),
		title:       cleanedTitle,
		author:      cleanedAuthor,
		description: strings.TrimSpace(description),
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

func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) GetAuthor() string {
	return b.author
}

func (b *Book) GetDescription() string {
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

func (b *Book) Update(title, author, description, coverUri, eBookUri string) error {
	cleanedTitle := strings.TrimSpace(title)
	if cleanedTitle == "" {
		return exceptions.ErrBookTitleIsEmpty
	}

	cleanedAuthor := strings.TrimSpace(author)
	if cleanedAuthor == "" {
		return exceptions.ErrBookAuthorIsEmpty
	}

	b.title = cleanedTitle
	b.author = cleanedAuthor
	b.description = strings.TrimSpace(description)
	b.coverUri = coverUri
	b.eBookUri = eBookUri
	b.updatedAt = time.Now()
	return nil
}

func UnmarshalBookFromDatabase(id uuid.UUID, title, author, description, coverUri, eBookUri string, isDeleted bool, createdAt, updatedAt time.Time) *Book {
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
