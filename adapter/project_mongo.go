package adapter

import (
	_interface "chaos-api/adapter/interface"
	"chaos-api/domain"
	"context"
	"errors"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"time"
)

type MongoDbProjectAdapter struct {
	Collection *mongo.Collection
}

func NewMongoDbProjectAdapter(client *mongo.Client) _interface.ProjectAdapter {
	return &MongoDbProjectAdapter{
		Collection: client.Database("chaos").Collection("projects"),
	}
}

func (a *MongoDbProjectAdapter) GetProjects() ([]domain.Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	find, err := a.Collection.Find(ctx, bson.M{})
	if errors.Is(err, mongo.ErrNoDocuments) {
		return []domain.Project{}, nil
	}

	if err != nil {
		return nil, err
	}

	var result []domain.Project
	err = find.Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *MongoDbProjectAdapter) CreateProject(name string) (*domain.Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	project := domain.Project{
		Id:   id.String(),
		Name: name,
	}
	_, err = a.Collection.InsertOne(ctx, project)
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (a *MongoDbProjectAdapter) UpdateProject(projectId, name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r, err := a.Collection.UpdateOne(ctx, bson.M{"id": projectId}, bson.M{"$set": domain.Project{
		Id:   projectId,
		Name: name,
	}})
	if err != nil {
		return err
	}

	if r.MatchedCount == 0 {
		return errors.New("project not found")
	}

	return nil
}

func (a *MongoDbProjectAdapter) DeleteProject(projectId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := a.Collection.DeleteOne(ctx, bson.M{"id": projectId})
	if err != nil {
		return err
	}

	return nil
}
