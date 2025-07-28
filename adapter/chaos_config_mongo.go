package adapter

import (
	_interface "chaos-api/adapter/interface"
	"chaos-api/domain"
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"time"
)

type MongoDbChaosConfigAdapter struct {
	Collection *mongo.Collection
}

func NewMongoDbChaosConfigAdapter() _interface.ChaosConfigAdapter {
	//uri := os.Getenv("MONGODB_URI")
	client, _ := mongo.Connect(options.Client().ApplyURI("mongodb://root:example@localhost:27017"))

	return &MongoDbChaosConfigAdapter{
		client.Database("chaos").Collection("config"),
	}
}

func (a *MongoDbChaosConfigAdapter) UpsertChaosConfig(c *domain.ChaosConfig) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r, err := a.Collection.UpdateOne(ctx, bson.M{"project_id": c.ProjectId}, bson.M{"$set": c})
	if err != nil {
		return err
	}

	if r.MatchedCount == 0 {
		_, err := a.Collection.InsertOne(ctx, c)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *MongoDbChaosConfigAdapter) GetChaosConfigByProjectId(projectId string) ([]domain.ChaosConfig, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	find, err := a.Collection.Find(ctx, bson.M{"project_id": projectId})
	if err != nil {
		return nil, err
	}

	var configs []domain.ChaosConfig
	err = find.Decode(&configs)
	if err != nil {
		return nil, err
	}

	return configs, nil
}

func (a *MongoDbChaosConfigAdapter) GetChaosConfigByService(projectId string, service string) (*domain.ChaosConfig, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var config domain.ChaosConfig
	err := a.Collection.FindOne(ctx, bson.M{"project_id": projectId, "name": service}).Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (a *MongoDbChaosConfigAdapter) ResetConfig(projectId string, service string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := a.Collection.DeleteOne(ctx, bson.M{"project_id": projectId, "name": service})
	if err != nil {
		return err
	}

	return nil
}
