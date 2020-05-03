package repository

import (
	"context"
	error2 "github.com/tamaApotek/tama-go-server/query"

	"github.com/tamaApotek/tama-go-server/models"
	"github.com/tamaApotek/tama-go-server/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userMongo struct {
	col *mongo.Collection
}

func NewUserMongo(db *mongo.Database) user.Repository {
	col := db.Collection("user")
	return &userMongo{col}
}

func (um *userMongo) Create(ctx context.Context, user *models.User) (string, error) {
	res, err := um.col.InsertOne(ctx, user)

	if err != nil {
		return "", error2.NewErrorQuery("Failed creating new document", error2.ErrorEnum.Internal, err)
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
	if err != nil {
		return error2.NewErrorQuery("Failed updating document", error2.ErrorEnum.Internal, err)
	}

	return nil
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

func (um *userMongo) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	q := um.col.FindOne(ctx, bson.D{
		{
			Key: "email", Value: email,
		},
	})

	err := q.Err()
	if err != nil {
		return nil, error2.NewErrorQuery("Failed finding document", error2.ErrorEnum.Internal, err)
	}

	user := new(models.User)
	err = q.Decode(user)
	if err != nil {
		return nil, error2.NewErrorQuery("Unknown query occured", error2.ErrorEnum.Internal, err)
	}

	return user, nil
}

func (um *userMongo) SearchText(ctx context.Context, queryString string) ([]*models.User, error) {
	filter := bson.D{
		{
			Key: "$text", Value: bson.D{
				{Key: "$search", Value: queryString},
			},
		},
	}

	q, err := um.col.Find(ctx, filter)
	defer q.Close(ctx)

	if err != nil {
		return nil, error2.NewErrorQuery("Unknown query occured", error2.ErrorEnum.Internal, err)
	}

	var users []*models.User
	for q.Next(ctx) {
		var user *models.User

		err = q.Decode(user)
		if err != nil {
			continue
		}

		users = append(users, user)
	}

	return users, nil
}
