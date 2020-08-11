package main

import (
	"context"
	pb "github.com/wkcw/shippy/shippy-service-vessel/proto/vessel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository interface {
	FindAvailable(ctx context.Context, spec *Specification) (*Vessel, error)
	Create(ctx context.Context, vessel *Vessel) error
}

type MongoRepository struct {
	collection *mongo.Collection
}

type Specification struct {
	Capacity int32
	MaxWeight int32
}

type Vessel struct {
	id string
	capacity int32
	max_weight int32
	name string
	available bool
	owner_id string
}

func MarshalSpecification(spec *pb.Specification) *Specification{
	return &Specification{
		MaxWeight: spec.MaxWeight,
		Capacity: spec.Capacity,
	}
}

func UnmarshalSpecification(spec *Specification) *pb.Specification{
	return &pb.Specification{
		MaxWeight: spec.MaxWeight,
		Capacity: spec.Capacity,
	}
}

func MarshalVessel(vessel *pb.Vessel) *Vessel{
	return &Vessel{
		id: vessel.Id,
		capacity: vessel.Capacity,
		max_weight: vessel.MaxWeight,
		name: vessel.Name,
		available: vessel.Available,
		owner_id: vessel.OwnerId,
	}
}

func UnmarshalVessel(vessel *Vessel) *pb.Vessel{
	return &pb.Vessel{
		Id: vessel.id,
		Capacity: vessel.capacity,
		MaxWeight: vessel.max_weight,
		Name: vessel.name,
		Available: vessel.available,
		OwnerId: vessel.owner_id,
	}
}

func (repo *MongoRepository) Create(ctx context.Context, vessel *Vessel) error {
	_, err := repo.collection.InsertOne(ctx, vessel)
	return err
}

func (repo *MongoRepository) FindAvailable(ctx context.Context, spec *Specification) (*Vessel, error) {
	filter := bson.M{
		"capacity": bson.M{"gte": spec.Capacity},
		"maxweight": bson.M{"gte": spec.MaxWeight},
	}

	vessel := &Vessel{}
	if err := repo.collection.FindOne(ctx, filter).Decode(vessel); err != nil {
		return nil, err
	}
	return vessel, nil
}