package user

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userMongo struct {
	col *mongo.Collection
}

func NewRepoMongo(db *mongo.Database) Repository {
	col := db.Collection("users")
	return &userMongo{col}
}

func (um *userMongo) Create(ctx context.Context, user *User) (string, error) {
	res, err := um.col.InsertOne(ctx, user)

	if err != nil {
		return "", err
	}

	fmt.Printf("created user ====> %+v\n", res)
	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		fmt.Printf("user oid ======> %+v\n", oid)
		user.ID = oid
	}

	return user.ID.Hex(), nil
}

func (um *userMongo) UpdateOne(ctx context.Context, user *User) error {
	res := um.col.FindOneAndUpdate(
		ctx, bson.D{
			{
				Key: "_id", Value: user.ID,
			},
		}, user)

	err := res.Err()

	return err
}

func (um *userMongo) FindByID(ctx context.Context, UID string) (user *User, err error) {
	q := um.col.FindOne(ctx, bson.D{{Key: "uid", Value: UID}})

	err = q.Err()

	if err != nil {
		return nil, err
	}

	err = q.Decode(&user)
	return user, err
}

func (um *userMongo) FindByEmail(ctx context.Context, email string) (*User, error) {
	q := um.col.FindOne(ctx, bson.D{
		{
			Key: "email", Value: email,
		},
	})

	err := q.Err()
	if err != nil {
		return nil, err
	}

	user := new(User)
	err = q.Decode(user)

	return user, err
}

func (um *userMongo) SearchText(ctx context.Context, queryString string) ([]*User, error) {
	filter := bson.D{
		{
			Key: "$text", Value: bson.D{
				{Key: "$search", Value: queryString},
			},
		},
	}
	cursor, err := um.col.Find(ctx, filter)
	defer func() {
		if cursor != nil {
			_ = cursor.Close(ctx)
		}
	}()

	if err != nil {
		return nil, err
	}

	var users []*User
	for cursor.Next(ctx) {
		var user *User

		err = cursor.Decode(user)
		if err != nil {
			continue
		}

		users = append(users, user)
	}

	return users, nil
}
