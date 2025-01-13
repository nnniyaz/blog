package book

import (
	"context"
	"github.com/nnniyaz/blog/domain/base/uuid"
	"github.com/nnniyaz/blog/domain/book"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type BookRepo struct {
	client *mongo.Client
}

func NewBookRepo(client *mongo.Client) *BookRepo {
	return &BookRepo{client: client}
}

func (r *BookRepo) Coll() *mongo.Collection {
	return r.client.Database("main").Collection("books")
}

type mongoBook struct {
	Id          uuid.UUID `bson:"_id"`
	Title       string    `bson:"title"`
	Author      string    `bson:"author"`
	Description string    `bson:"description"`
	CoverUri    string    `bson:"coverUri"`
	EBookUri    string    `bson:"eBookUri"`
	IsDeleted   bool      `bson:"isDeleted"`
	CreatedAt   time.Time `bson:"createdAt"`
	UpdatedAt   time.Time `bson:"updatedAt"`
}

func newFromBook(book *book.Book) *mongoBook {
	return &mongoBook{
		Id:          book.GetId(),
		Title:       book.GetTitle(),
		Author:      book.GetAuthor(),
		Description: book.GetDescription(),
		CoverUri:    book.GetCoverUri(),
		EBookUri:    book.GetEBookUri(),
		IsDeleted:   book.GetIsDeleted(),
		CreatedAt:   book.GetCreatedAt(),
		UpdatedAt:   book.GetUpdatedAt(),
	}
}

func (m *mongoBook) ToAggregate() *book.Book {
	return book.UnmarshalBookFromDatabase(m.Id, m.Title, m.Author, m.Description, m.CoverUri, m.EBookUri, m.IsDeleted, m.CreatedAt, m.UpdatedAt)
}

func (r *BookRepo) Create(ctx context.Context, book *book.Book) error {
	_, err := r.Coll().InsertOne(ctx, newFromBook(book))
	return err
}

func (r *BookRepo) Update(ctx context.Context, book *book.Book) error {
	_, err := r.Coll().UpdateOne(ctx, bson.M{"_id": book.GetId()}, bson.M{"$set": newFromBook(book)})
	return err
}

func (r *BookRepo) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.Coll().UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{
		"isDeleted": true,
		"updatedAt": time.Now(),
	}})
	return err
}

func (r *BookRepo) Restore(ctx context.Context, id uuid.UUID) error {
	_, err := r.Coll().UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{
		"isDeleted": false,
		"updatedAt": time.Now(),
	}})
	return err
}

func (r *BookRepo) FindById(ctx context.Context, id uuid.UUID) (*book.Book, error) {
	var m mongoBook
	if err := r.Coll().FindOne(ctx, bson.M{"_id": id}).Decode(&m); err != nil {
		return nil, err
	}
	return m.ToAggregate(), nil
}

func (r *BookRepo) FindAll(ctx context.Context, offset, limit int64, isDeleted bool, search string) ([]*book.Book, int64, error) {
	filters := bson.D{
		{"isDeleted", isDeleted},
	}
	if search != "" {
		filters = append(filters, bson.E{"$or", bson.A{
			bson.M{"title": bson.M{"$regex": search, "$options": "i"}},
		}})
	}
	cursor, err := r.Coll().Find(ctx, filters, &options.FindOptions{
		Skip:  &offset,
		Limit: &limit,
	})
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	count, err := r.Coll().CountDocuments(ctx, filters)
	if err != nil {
		return nil, 0, err
	}

	var books []*book.Book
	for cursor.Next(ctx) {
		var m mongoBook
		if err := cursor.Decode(&m); err != nil {
			return nil, 0, err
		}
		books = append(books, m.ToAggregate())
	}

	return books, count, nil
}
