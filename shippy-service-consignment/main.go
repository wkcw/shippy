package main

import (
	pb "github.com/wkcw/shippy-service-consignment/proto/consignment"
	"log"

	// 使用 go-micro
	micro "github.com/micro/go-micro/v2"
	"golang.org/x/net/context"
)

const (
	port = ":50051"
)

type IRepository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// Repository - Simulate a database, it will be replaced with a true one.
type Repository struct {
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error){
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

type service struct {
	repo IRepository
}

// CreateConsignment
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	res.Created = true
	res.Consignment = consignment
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

func main() {
	repo := &Repository{}

	srv := micro.NewService(
		micro.Name("shippy.consignment.service"),
		micro.Version("latest"),
	)

	srv.Init()

	if err := pb.RegisterShippingServiceHandler(srv.Server(), &service{repo}); err != nil {
		log.Panic(err)
	}

	if err := srv.Run(); err != nil {
		log.Panic(err)
	}
}

