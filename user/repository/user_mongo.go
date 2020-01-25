package repository

import (
	"context"

	"github.com/tamaApotek/tama-go-server/models"
	"github.com/tamaApotek/tama-go-server/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userMongo struct {
	col *mongo.Collection
}

func NewUserMongo(c *mongo.Client) user.Repository {
	col := c.Database("tama").Collection("user")
	return &userMongo{col}
}

func (um *userMongo) FindByUID(ctx context.Context, UID string) (user *models.User, err error) {
	q := um.col.FindOne(ctx, bson.M{"uid": UID})
	err = q.Err()

	if err != nil {
		return nil, err
	}

	err = q.Decode(&user)
	return user, nil
}
