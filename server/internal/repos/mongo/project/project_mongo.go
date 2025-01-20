package projectRepo

import (
	"context"
	"github.com/nnniyaz/blog/internal/domain/base/uuid"
	"github.com/nnniyaz/blog/internal/domain/project"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type ProjectRepo struct {
	client *mongo.Client
}

func NewProjectRepo(client *mongo.Client) *ProjectRepo {
	return &ProjectRepo{client: client}
}

func (r *ProjectRepo) Coll() *mongo.Collection {
	return r.client.Database("main").Collection("projects")
}

type mongoProject struct {
	Id             uuid.UUID `bson:"_id"`
	Name           string    `bson:"name"`
	Description    string    `bson:"description"`
	CoverUri       string    `bson:"coverUri"`
	AppLink        string    `bson:"appLink"`
	SourceCodeLink string    `bson:"sourceCodeLink"`
	IsDeleted      bool      `bson:"isDeleted"`
	CreatedAt      time.Time `bson:"createdAt"`
	UpdatedAt      time.Time `bson:"updatedAt"`
}

func newFromProject(project *project.Project) *mongoProject {
	return &mongoProject{
		Id:             project.GetId(),
		Name:           project.GetName(),
		Description:    project.GetDescription(),
		CoverUri:       project.GetCoverUri(),
		AppLink:        project.GetAppLink(),
		SourceCodeLink: project.GetSourceCodeLink(),
		IsDeleted:      project.GetIsDeleted(),
		CreatedAt:      project.GetCreatedAt(),
		UpdatedAt:      project.GetUpdatedAt(),
	}
}

func (m *mongoProject) ToAggregate() *project.Project {
	return project.UnmarshalProjectFromDatabase(m.Id, m.Name, m.Description, m.CoverUri, m.AppLink, m.SourceCodeLink, m.IsDeleted, m.CreatedAt, m.UpdatedAt)
}

func (r *ProjectRepo) Create(ctx context.Context, project *project.Project) error {
	_, err := r.Coll().InsertOne(ctx, newFromProject(project))
	return err
}

func (r *ProjectRepo) Update(ctx context.Context, project *project.Project) error {
	_, err := r.Coll().UpdateOne(ctx, bson.M{"_id": project.GetId()}, bson.M{"$set": newFromProject(project)})
	return err
}

func (r *ProjectRepo) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.Coll().UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{
		"isDeleted": true,
		"updatedAt": time.Now(),
	}})
	return err
}

func (r *ProjectRepo) Restore(ctx context.Context, id uuid.UUID) error {
	_, err := r.Coll().UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{
		"isDeleted": false,
		"updatedAt": time.Now(),
	}})
	return err
}

func (r *ProjectRepo) FindById(ctx context.Context, id uuid.UUID) (*project.Project, error) {
	var project mongoProject
	if err := r.Coll().FindOne(ctx, bson.M{"_id": id}).Decode(&project); err != nil {
		return nil, err
	}
	return project.ToAggregate(), nil
}

func (r *ProjectRepo) FindAll(ctx context.Context, offset, limit int64, isDeleted bool, search string) ([]*project.Project, int64, error) {
	filters := bson.D{
		{"isDeleted", isDeleted},
	}

	if search != "" {
		filters = append(filters, bson.E{"$or", bson.A{
			bson.M{"name": bson.M{"$regex": search, "$options": "i"}},
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

	var projects []*project.Project
	for cur.Next(ctx) {
		var project mongoProject
		if err := cur.Decode(&project); err != nil {
			return nil, 0, err
		}
		projects = append(projects, project.ToAggregate())
	}
	return projects, count, nil
}
