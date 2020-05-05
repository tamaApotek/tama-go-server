package doctor

import (
	"context"
	"github.com/tamaApotek/tama-go-server/domains/query"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type doctorMongo struct {
	col *mongo.Collection
}

func NewRepoMongo(db *mongo.Database) Repository {
	col := db.Collection("doctors")

	return &doctorMongo{col}
}

func (dm *doctorMongo) Add(ctx context.Context, doctor *Doctor) (string, error) {
	res, err := dm.col.InsertOne(ctx, doctor)
	if err != nil {
		return "", query.NewErrorQuery("Internal server error", query.ErrorEnum.Internal, err)
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		doctor.ID = &oid
	}

	return doctor.ID.Hex(), nil
}
