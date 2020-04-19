package repository

import (
	"context"
	"fmt"

	"github.com/tamaApotek/tama-go-server/models"
	"github.com/tamaApotek/tama-go-server/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userMongo struct {
	col *mongo.Collection
}

func NewUserMongo(c *mongo.Client) user.Repository {
	col := c.Database("tama").Collection("user")
	return &userMongo{col}
}

func (um *userMongo) Create(ctx context.Context, user *models.User) (string, error) {
	res, err := um.col.InsertOne(ctx, user)

	if err != nil {
		return "", fmt.Errorf("Mongo Insert error: %w", models.ErrInternal)
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		user.ID = &oid
	}

	return user.ID.Hex(), nil
}

func (um *userMongo) UpdateOne(ctx context.Context, user *models.User) error {
	res := um.col.FindOneAndUpdate(
		ctx, bson.D{
			{
				Key: "_id", Value: user.ID,
			},
		}, user)

	err := res.Err()

	return err
}

func (um *userMongo) FindByID(ctx context.Context, UID string) (user *models.User, err error) {
	q := um.col.FindOne(ctx, bson.D{{Key: "uid", Value: UID}})
	err = q.Err()

	if err != nil {
		return nil, err
	}

	err = q.Decode(&user)
	return user, nil
}
