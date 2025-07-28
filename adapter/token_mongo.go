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

type MongoDbTokenAdapter struct {
	Collection *mongo.Collection
}

func NewMongoDbTokenAdapter(client *mongo.Client) _interface.TokenAdapter {
	return &MongoDbTokenAdapter{
		Collection: client.Database("chaos").Collection("Tokens"),
	}
}

func (a *MongoDbTokenAdapter) GetTokens() ([]domain.Token, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	find, err := a.Collection.Find(ctx, bson.M{})
	if errors.Is(err, mongo.ErrNoDocuments) {
		return []domain.Token{}, nil
	}

	if err != nil {
		return nil, err
	}

	var result []domain.Token
	err = find.Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *MongoDbTokenAdapter) GenerateToken(projectId, name string) (*domain.Token, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	token, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	result := domain.Token{
		Id:        id.String(),
		ProjectId: projectId,
		Value:     token.String(),
		Name:      name,
	}
	_, err = a.Collection.InsertOne(ctx, result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (a *MongoDbTokenAdapter) DeleteToken(tokenId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := a.Collection.DeleteOne(ctx, bson.M{"id": tokenId})
	if err != nil {
		return err
	}

	return nil
}
