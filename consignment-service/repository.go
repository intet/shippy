package main
import (
	"context"
	pb "github.com/intet/shippy/consignment-service/proto/consignment"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository interface {
	Create(ctx context.Context, consignment *pb.Consignment) error
	GetAll(ctx context.Context) ([]*pb.Consignment, error)
}

// MongoRepository implementation
type MongoRepository struct {
	collection *mongo.Collection
}

// Create -
func (repository *MongoRepository) Create(
	ctx context.Context,
	consignment *pb.Consignment,
) error {
	_, err := repository.collection.InsertOne(ctx, consignment)
	return err
}

// GetAll -
func (repository *MongoRepository) GetAll(ctx context.Context) ([]*pb.Consignment, error) {
	cur, err := repository.collection.Find(ctx, nil, nil)
	var consignments []*pb.Consignment
	for cur.Next(ctx) {
		var consignment *pb.Consignment
		if err := cur.Decode(&consignment); err != nil {
			return nil, err
		}
		consignments = append(consignments, consignment)
	}
	return consignments, err
}