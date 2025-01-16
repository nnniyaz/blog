package articleRepo

import (
	"context"
	"github.com/nnniyaz/blog/internal/domain/article"
	"github.com/nnniyaz/blog/internal/domain/base/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type RepoArticle struct {
	client *mongo.Client
}

func NewRepoArticle(client *mongo.Client) *RepoArticle {
	return &RepoArticle{client: client}
}

func (r *RepoArticle) Coll() *mongo.Collection {
	return r.client.Database("main").Collection("articles")
}

type mongoArticle struct {
	Id        uuid.UUID `bson:"_id"`
	Title     string    `bson:"title"`
	Content   string    `bson:"content"`
	IsDeleted bool      `bson:"isDeleted"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
	Version   int       `bson:"version"`
}

func newFromArticle(article *article.Article) *mongoArticle {
	return &mongoArticle{
		Id:        article.GetId(),
		Title:     article.GetTitle(),
		Content:   article.GetContent(),
		IsDeleted: article.GetIsDeleted(),
		CreatedAt: article.GetCreatedAt(),
		UpdatedAt: article.GetUpdatedAt(),
		Version:   article.GetVersion(),
	}
}

func (m *mongoArticle) ToAggregate() *article.Article {
	return article.UnmarshalArticleFromDatabase(m.Id, m.Title, m.Content, m.IsDeleted, m.CreatedAt, m.UpdatedAt, m.Version)
}

func (r *RepoArticle) Create(ctx context.Context, article *article.Article) error {
	_, err := r.Coll().InsertOne(ctx, newFromArticle(article))
	return err
}

func (r *RepoArticle) Update(ctx context.Context, article *article.Article) error {
	_, err := r.Coll().UpdateOne(ctx, bson.M{"_id": article.GetId()}, bson.M{"$set": newFromArticle(article)})
	return err
}

func (r *RepoArticle) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.Coll().UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{
		"isDeleted": true,
		"updatedAt": time.Now(),
	}})
	return err
}

func (r *RepoArticle) Restore(ctx context.Context, id uuid.UUID) error {
	_, err := r.Coll().UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{
		"isDeleted": false,
		"updatedAt": time.Now(),
	}})
	return err
}

func (r *RepoArticle) FindById(ctx context.Context, id uuid.UUID) (*article.Article, error) {
	var result mongoArticle
	if err := r.Coll().FindOne(ctx, bson.M{"_id": id}).Decode(&result); err != nil {
		return nil, err
	}
	return result.ToAggregate(), nil
}

func (r *RepoArticle) FindAll(ctx context.Context, offset, limit int64, isDeleted bool, search string) ([]*article.Article, int64, error) {
	filters := bson.D{
		{"isDeleted", isDeleted},
	}

	if search != "" {
		filters = append(filters, bson.E{"$or", bson.A{
			bson.M{"title": bson.M{"$regex": search, "$options": "i"}},
		}})
	}

	cur, err := r.Coll().Find(ctx, filters, &options.FindOptions{
		Skip:  &offset,
		Limit: &limit,
	})
	if err != nil {
		return nil, 0, err
	}

	count, err := r.Coll().CountDocuments(ctx, filters)
	if err != nil {
		return nil, 0, err
	}

	var articles []*article.Article
	for cur.Next(ctx) {
		var result mongoArticle
		if err := cur.Decode(&result); err != nil {
			return nil, 0, err
		}
		articles = append(articles, result.ToAggregate())
	}
	return articles, count, nil
}
