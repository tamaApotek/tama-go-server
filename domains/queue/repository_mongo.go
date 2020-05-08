package queue

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type queueMongoRepo struct {
	col *mongo.Collection
}

func NewRepoMongo(db *mongo.Database) Repository {
	col := db.Collection("queues")

	return &queueMongoRepo{col}
}

func (qm *queueMongoRepo) Add(ctx context.Context, queue *Queue) (string, error) {
	res, err := qm.col.InsertOne(ctx, queue)

	if err != nil {
		return "", fmt.Errorf("Error add queue: %w", err)
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		queue.ID = &oid
	}

	return queue.ID.Hex(), nil

}