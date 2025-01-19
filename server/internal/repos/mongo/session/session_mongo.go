package session

import (
	"context"
	"github.com/nnniyaz/blog/internal/domain/base/uuid"
	"github.com/nnniyaz/blog/internal/domain/session"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type SessionRepo struct {
	client *mongo.Client
}

func NewRepoSession(client *mongo.Client) *SessionRepo {
	return &SessionRepo{client: client}
}

func (r *SessionRepo) Coll() *mongo.Collection {
	return r.client.Database("blog").Collection("sessions")
}

type mongoSession struct {
	Id        uuid.UUID `bson:"_id"`
	UserId    uuid.UUID `bson:"user_id"`
	Session   uuid.UUID `bson:"session"`
	UserAgent string    `bson:"user_agent"`
	CreatedAt time.Time `bson:"created_at"`
}

func newMongoSession(s *session.Session) *mongoSession {
	return &mongoSession{
		Id:        s.GetId(),
		UserId:    s.GetUserId(),
		Session:   s.GetSession(),
		UserAgent: s.GetUserAgent().String(),
		CreatedAt: s.GetCreatedAt(),
	}
}

func (m *mongoSession) ToAggregate() *session.Session {
	return session.UnmarshalSessionFromDatabase(m.Id, m.UserId, m.Session, m.UserAgent, m.CreatedAt)
}

func (r *SessionRepo) Create(ctx context.Context, s *session.Session) error {
	_, err := r.Coll().InsertOne(ctx, newMongoSession(s))
	return err
}

func (r *SessionRepo) DeleteBySession(ctx context.Context, session uuid.UUID) error {
	_, err := r.Coll().DeleteOne(ctx, map[string]interface{}{"session": session})
	return err
}

func (r *SessionRepo) DeleteByUserId(ctx context.Context, userId uuid.UUID) error {
	_, err := r.Coll().DeleteMany(ctx, map[string]interface{}{"user_id": userId})
	return err
}

func (r *SessionRepo) FindAll(ctx context.Context, offset, limit int64) ([]*session.Session, int64, error) {
	cur, err := r.Coll().Find(ctx, bson.D{{}}, &options.FindOptions{
		Skip:  &offset,
		Limit: &limit,
	})
	if err != nil {
		return nil, 0, err
	}

	count, err := r.Coll().CountDocuments(ctx, bson.D{{}})
	if err != nil {
		return nil, 0, err
	}

	var sessions []*session.Session
	for cur.Next(ctx) {
		var session mongoSession
		if err := cur.Decode(&session); err != nil {
			return nil, 0, err
		}
		sessions = append(sessions, session.ToAggregate())
	}
	return sessions, count, nil
}

func (r *SessionRepo) FindBySession(ctx context.Context, session uuid.UUID) (*session.Session, error) {
	var s mongoSession
	err := r.Coll().FindOne(ctx, map[string]interface{}{"session": session}).Decode(&s)
	if err != nil {
		return nil, err
	}
	return s.ToAggregate(), nil
}

func (r *SessionRepo) FindByUserId(ctx context.Context, userId uuid.UUID) ([]*session.Session, error) {
	cursor, err := r.Coll().Find(ctx, map[string]interface{}{"user_id": userId})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(nil)
	var sessions []*session.Session
	for cursor.Next(nil) {
		var s mongoSession
		err := cursor.Decode(&s)
		if err != nil {
			return nil, err
		}
		sessions = append(sessions, s.ToAggregate())
	}
	return sessions, nil
}
