package bioRepo

import (
	"context"
	"github.com/nnniyaz/blog/internal/domain/base/uuid"
	"github.com/nnniyaz/blog/internal/domain/bio"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type RepoBio struct {
	client *mongo.Client
}

func NewRepoBio(client *mongo.Client) *RepoBio {
	return &RepoBio{client: client}
}

func (r *RepoBio) Coll() *mongo.Collection {
	return r.client.Database("main").Collection("bio")
}

type mongoBio struct {
	Id        uuid.UUID `bson:"_id"`
	Bio       string    `bson:"bio"`
	Active    bool      `bson:"active"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
	Version   int       `bson:"version"`
}

func newFromBio(bio *bio.Bio) *mongoBio {
	return &mongoBio{
		Id:        bio.GetId(),
		Bio:       bio.GetBio(),
		Active:    bio.GetActive(),
		CreatedAt: bio.GetCreatedAt(),
		UpdatedAt: bio.GetUpdatedAt(),
		Version:   bio.GetVersion(),
	}
}

func (m *mongoBio) ToAggregate() *bio.Bio {
	return bio.UnmarshalBioFromDatabase(m.Id, m.Bio, m.Active, m.CreatedAt, m.UpdatedAt, m.Version)
}

func (r *RepoBio) Create(ctx context.Context, bio *bio.Bio) error {
	_, err := r.Coll().InsertOne(ctx, newFromBio(bio))
	return err
}

func (r *RepoBio) Update(ctx context.Context, bio *bio.Bio) error {
	_, err := r.Coll().UpdateOne(ctx, bson.M{"_id": bio.GetId()}, bson.M{"$set": newFromBio(bio)})
	return err
}

func (r *RepoBio) SetActive(ctx context.Context, id uuid.UUID) error {
	_, err := r.Coll().UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{
		"active": true,
	}})
	if err != nil {
		return err
	}
	_, err = r.Coll().UpdateMany(ctx, bson.M{"_id": bson.M{"$ne": id}}, bson.M{"$set": bson.M{
		"active": false,
	}})
	return err
}

func (r *RepoBio) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.Coll().UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{
		"isDeleted": true,
		"updatedAt": time.Now(),
	}})
	return err
}

func (r *RepoBio) Restore(ctx context.Context, id uuid.UUID) error {
	_, err := r.Coll().UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{
		"isDeleted": false,
		"updatedAt": time.Now(),
	}})
	return err
}

func (r *RepoBio) FindById(ctx context.Context, id uuid.UUID) (*bio.Bio, error) {
	var result mongoBio
	if err := r.Coll().FindOne(ctx, bson.M{"_id": id}).Decode(&result); err != nil {
		return nil, err
	}
	return result.ToAggregate(), nil
}

func (r *RepoBio) FindByActive(ctx context.Context) (*bio.Bio, error) {
	var result mongoBio
	if err := r.Coll().FindOne(ctx, bson.M{"active": true}).Decode(&result); err != nil {
		return nil, err
	}
	return result.ToAggregate(), nil
}

func (r *RepoBio) FindAll(ctx context.Context, offset, limit int64) ([]*bio.Bio, int64, error) {
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

	var bios []*bio.Bio
	for cursor.Next(ctx) {
		var bio mongoBio
		if err := cursor.Decode(&bio); err != nil {
			return nil, 0, err
		}
		bios = append(bios, bio.ToAggregate())
	}
	return bios, count, nil
}
