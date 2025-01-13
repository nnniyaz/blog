package mongo

import (
	"context"
	"github.com/nnniyaz/blog/pkg/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Client interface {
	Close(ctx context.Context) error
	StartSession() (mongo.Session, error)
}

func New(ctx context.Context, mongoUri string) (*mongo.Client, error) {
	mongoOptions := options.Client().ApplyURI(mongoUri).SetRegistry(uuid.MongoRegistry)
	client, err := mongo.Connect(ctx, mongoOptions)
	if err != nil {
		return nil, err
	}
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	return client, nil
}
