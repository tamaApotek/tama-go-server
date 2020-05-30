package doctor

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type doctorMongo struct {
	col *mongo.Collection
}

var repo *doctorMongo

func NewRepoMongo(db *mongo.Database) Repository {
	col := db.Collection("doctors")

	if repo != nil {
		return repo
	}

	repo = &doctorMongo{col}
	return repo
}

func (dm *doctorMongo) Add(ctx context.Context, doctor *Doctor) (string, error) {
	res, err := dm.col.InsertOne(ctx, doctor)
	if err != nil {
		return "", err
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		doctor.ID = &oid
	}

	return doctor.ID.Hex(), nil
}

func (dm *doctorMongo) FindByID(ctx context.Context, id primitive.ObjectID) (*Doctor, error) {
	res := dm.col.FindOne(ctx, bson.M{"_id": id})
	err := res.Err()
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}

		return nil, err
	}

	var doctor Doctor
	err = res.Decode(&doctor)
	return &doctor, err
}
