package contactRepo

import (
	"context"
	"github.com/nnniyaz/blog/internal/domain/base/uuid"
	"github.com/nnniyaz/blog/internal/domain/contact"
	"github.com/nnniyaz/blog/pkg/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type RepoContact struct {
	client *mongo.Client
}

func NewRepoContact(client *mongo.Client) *RepoContact {
	return &RepoContact{client: client}
}

func (r *RepoContact) Coll() *mongo.Collection {
	return r.client.Database("main").Collection("contacts")
}

type mongoContact struct {
	Id        uuid.UUID     `bson:"_id"`
	Label     core.MlString `bson:"label"`
	Link      string        `bson:"link"`
	IsDeleted bool          `bson:"isDeleted"`
	CreatedAt time.Time     `bson:"createdAt"`
	UpdatedAt time.Time     `bson:"updatedAt"`
}

func newFromContact(contact *contact.Contact) *mongoContact {
	return &mongoContact{
		Id:        contact.GetId(),
		Label:     contact.GetLabel(),
		Link:      contact.GetLink(),
		IsDeleted: contact.GetIsDeleted(),
		CreatedAt: contact.GetCreatedAt(),
		UpdatedAt: contact.GetUpdatedAt(),
	}
}

func (m *mongoContact) ToAggregate() *contact.Contact {
	return contact.UnmarshalContactFromDatabase(m.Id, m.Label, m.Link, m.IsDeleted, m.CreatedAt, m.UpdatedAt)
}

func (r *RepoContact) Create(ctx context.Context, contact *contact.Contact) error {
	_, err := r.Coll().InsertOne(ctx, newFromContact(contact))
	return err
}

func (r *RepoContact) Update(ctx context.Context, contact *contact.Contact) error {
	_, err := r.Coll().UpdateOne(ctx, bson.M{"_id": contact.GetId()}, bson.M{"$set": newFromContact(contact)})
	return err
}

func (r *RepoContact) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.Coll().UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{
		"isDeleted": true,
		"updatedAt": time.Now(),
	}})
	return err
}

func (r *RepoContact) Restore(ctx context.Context, id uuid.UUID) error {
	_, err := r.Coll().UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{
		"isDeleted": false,
		"updatedAt": time.Now(),
	}})
	return err
}

func (r *RepoContact) FindById(ctx context.Context, id uuid.UUID) (*contact.Contact, error) {
	var result mongoContact
	if err := r.Coll().FindOne(ctx, bson.M{"_id": id}).Decode(&result); err != nil {
		return nil, err
	}
	return result.ToAggregate(), nil
}

func (r *RepoContact) FindAll(ctx context.Context, offset, limit int64, isDeleted bool, search string) ([]*contact.Contact, int64, error) {
	filters := bson.D{
		{"isDeleted", isDeleted},
	}

	if search != "" {
		filters = append(filters, bson.E{"$or", bson.A{
			bson.M{"label": bson.M{"$regex": search, "$options": "i"}},
		}})
	}

	cur, err := r.Coll().Find(ctx, filters, &options.FindOptions{
		Skip:  &offset,
		Limit: &limit,
	})
	if err != nil {
		return nil, 0, err
	}
	defer cur.Close(ctx)

	count, err := r.Coll().CountDocuments(ctx, filters)
	if err != nil {
		return nil, 0, err
	}

	var contacts []*contact.Contact
	for cur.Next(ctx) {
		var contact mongoContact
		if err := cur.Decode(&contact); err != nil {
			return nil, 0, err
		}
		contacts = append(contacts, contact.ToAggregate())
	}
	return contacts, count, nil
}
