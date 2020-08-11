package main

import (
	"context"
	"errors"
	"log"

	pb "github.com/wkcw/shippy/shippy-service-vessel/proto/vessel"
	"github.com/micro/go-micro/v2"
)

type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

type VesselRepository struct {
	vessels []*pb.Vessel
}

func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	for _, vessel := range repo.vessels {
		if vessel.Capacity >= spec.Capacity && vessel.MaxWeight >= spec.MaxWeight {
			vessel.Capacity -= spec.Capacity
			vessel.MaxWeight -= spec.MaxWeight
			return vessel, nil
		}
	}
	return nil, errors.New("No vessel found by that spec. ")
}

type vesselService struct {
	repo Repository
}

func (s *vesselService) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {

	// Find the next available vessel
	vessel, err := s.repo.FindAvailable(req)
	if err != nil {
		return err
	}

	// Set the vessel as part of the response message type
	res.Vessel = vessel
	return nil
}

func main() {
	vessels := []*pb.Vessel{
		&pb.Vessel{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
	}
	repo := &VesselRepository{vessels}

	service := micro.NewService(
		micro.Name("shippy.service.vessel"),
	)

	service.Init()

	// Register our implementation with
	if err := pb.RegisterVesselServiceHandler(service.Server(), &vesselService{repo}); err != nil {
		log.Panic(err)
	}

	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
