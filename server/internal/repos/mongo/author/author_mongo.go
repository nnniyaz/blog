package authorRepo

import (
	"context"
	"github.com/nnniyaz/blog/internal/domain/author"
	"github.com/nnniyaz/blog/internal/domain/base/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type RepoAuthor struct {
	client *mongo.Client
}

func NewRepoAuthor(client *mongo.Client) *RepoAuthor {
	return &RepoAuthor{client: client}
}

func (r *RepoAuthor) Coll() *mongo.Collection {
	return r.client.Database("main").Collection("author")
}

type mongoAuthor struct {
	Id        uuid.UUID `bson:"_id"`
	FirstName string    `bson:"firstName"`
	LastName  string    `bson:"lastName"`
	AvatarUri string    `bson:"avatarUri"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
	Version   int       `bson:"version"`
}

func newFromAuthor(author *author.Author) *mongoAuthor {
	return &mongoAuthor{
		Id:        author.GetId(),
		FirstName: author.GetFirstName(),
		LastName:  author.GetLastName(),
		AvatarUri: author.GetAvatarUri(),
		CreatedAt: author.GetCreatedAt(),
		UpdatedAt: author.GetUpdatedAt(),
		Version:   author.GetVersion(),
	}
}

func (m *mongoAuthor) ToAggregate() *author.Author {
	return author.UnmarshalAuthorFromDatabase(m.Id, m.FirstName, m.LastName, m.AvatarUri, m.CreatedAt, m.UpdatedAt, m.Version)
}

func (r *RepoAuthor) Create(ctx context.Context, author *author.Author) error {
	_, err := r.Coll().InsertOne(ctx, newFromAuthor(author))
	return err
}

func (r *RepoAuthor) Update(ctx context.Context, author *author.Author) error {
	_, err := r.Coll().UpdateOne(ctx, bson.M{"_id": author.GetId()}, bson.M{"$set": newFromAuthor(author)})
	return err
}

func (r *RepoAuthor) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.Coll().UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{
		"isDeleted": true,
		"updatedAt": time.Now(),
	}})
	return err
}

func (r *RepoAuthor) Restore(ctx context.Context, id uuid.UUID) error {
	_, err := r.Coll().UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{
		"isDeleted": false,
		"updatedAt": time.Now(),
	}})
	return err
}

func (r *RepoAuthor) FindById(ctx context.Context, id uuid.UUID) (*author.Author, error) {
	var result mongoAuthor
	if err := r.Coll().FindOne(ctx, bson.M{"_id": id}).Decode(&result); err != nil {
		return nil, err
	}
	return result.ToAggregate(), nil
}

func (r *RepoAuthor) FindAll(ctx context.Context, offset, limit int64) ([]*author.Author, int64, error) {
	cursor, err := r.Coll().Find(ctx, bson.M{}, &options.FindOptions{
		Skip:  &offset,
		Limit: &limit,
	})
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	count, err := r.Coll().CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}

	var authors []*author.Author
	for cursor.Next(ctx) {
		var author mongoAuthor
		if err := cursor.Decode(&author); err != nil {
			return nil, 0, err
		}
		authors = append(authors, author.ToAggregate())
	}
	return authors, count, nil
}
