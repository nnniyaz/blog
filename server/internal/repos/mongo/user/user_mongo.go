package userRepo

import (
	"context"
	"github.com/nnniyaz/blog/internal/domain/base/email"
	"github.com/nnniyaz/blog/internal/domain/base/uuid"
	"github.com/nnniyaz/blog/internal/domain/user"
	"github.com/nnniyaz/blog/internal/domain/user/valueobject"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type UserRepo struct {
	client *mongo.Client
}

func NewUserRepo(client *mongo.Client) *UserRepo {
	return &UserRepo{client: client}
}

func (r *UserRepo) Coll() *mongo.Collection {
	return r.client.Database("main").Collection("users")
}

type passwordMongo struct {
	Hash string `bson:"hash"`
	Salt string `bson:"salt"`
}

func newPasswordMongo(p valueobject.Password) passwordMongo {
	return passwordMongo{
		Hash: p.GetHash(),
		Salt: p.GetSalt(),
	}
}

type userMongo struct {
	Id        uuid.UUID     `bson:"_id"`
	Email     string        `bson:"email"`
	Password  passwordMongo `bson:"password"`
	Role      string        `bson:"role"`
	IsDeleted bool          `bson:"isDeleted"`
	CreatedAt time.Time     `bson:"createdAt"`
	UpdatedAt time.Time     `bson:"updatedAt"`
}

func newUserMongo(u user.User) userMongo {
	return userMongo{
		Id:        u.GetId(),
		Email:     u.GetEmail().String(),
		Password:  newPasswordMongo(u.GetPassword()),
		Role:      u.GetRole().String(),
		IsDeleted: u.GetIsDeleted(),
		CreatedAt: u.GetCreatedAt(),
		UpdatedAt: u.GetUpdatedAt(),
	}
}

func (m *userMongo) ToAggregate() *user.User {
	return user.UnmarshalUserFromDatabase(m.Id, m.Email, valueobject.UnmarshalPasswordFromDatabase(m.Password.Hash, m.Password.Salt), m.Role, m.IsDeleted, m.CreatedAt, m.UpdatedAt)
}

func (r *UserRepo) Create(ctx context.Context, u *user.User) error {
	_, err := r.Coll().InsertOne(ctx, newUserMongo(*u))
	return err
}

func (r *UserRepo) Update(ctx context.Context, u *user.User) error {
	_, err := r.Coll().UpdateOne(ctx, bson.M{"_id": u.GetId()}, bson.M{"$set": newUserMongo(*u)})
	return err
}

func (r *UserRepo) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.Coll().UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{
		"isDeleted": true,
		"updatedAt": time.Now(),
	}})
	return err
}

func (r *UserRepo) Restore(ctx context.Context, id uuid.UUID) error {
	_, err := r.Coll().UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{
		"isDeleted": false,
		"updatedAt": time.Now(),
	}})
	return err
}

func (r *UserRepo) FindById(ctx context.Context, id uuid.UUID) (*user.User, error) {
	var user userMongo
	if err := r.Coll().FindOne(ctx, bson.M{"_id": id}).Decode(&user); err != nil {
		return nil, err
	}
	return user.ToAggregate(), nil
}

func (r *UserRepo) FindByEmail(ctx context.Context, email email.Email) (*user.User, error) {
	var user userMongo
	if err := r.Coll().FindOne(ctx, bson.M{"email": email.String()}).Decode(&user); err != nil {
		return nil, err
	}
	return user.ToAggregate(), nil
}

func (r *UserRepo) FindAll(ctx context.Context, offset, limit int64, isDeleted bool, search string) ([]*user.User, int64, error) {
	filters := bson.D{
		{"isDeleted", isDeleted},
	}

	if search != "" {
		filters = append(filters, bson.E{"$or", bson.A{
			bson.M{"email": bson.M{"$regex": search, "$options": "i"}},
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

	var users []*user.User
	for cur.Next(ctx) {
		var user userMongo
		if err := cur.Decode(&user); err != nil {
			return nil, 0, err
		}
		users = append(users, user.ToAggregate())
	}
	return users, count, nil
}
