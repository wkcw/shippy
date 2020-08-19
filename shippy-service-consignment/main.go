package main

import (
	"context"
	"github.com/micro/go-micro/v2"
	"log"
	"os"

	pb "github.com/wkcw/shippy/shippy-service-consignment/proto/consignment"
	vesselProto "github.com/wkcw/shippy/shippy-service-vessel/proto/vessel"
)

const (
	defaultHost = "datastore:27017"
)

func main() {
	uri := os.Getenv("DB_HOST")

	if uri == "" {
		uri = defaultHost
	}

	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panicf("Could not connect to datastore with host %s - %v", uri, err)
	}
	defer client.Disconnect(context.Background())
	consignmentCollection := client.Database("shippy").Collection("consignments")

	repository := &MongoRepository{consignmentCollection}


	srv := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	vesselClient := vesselProto.NewVesselService("go.micro.srv.vessel", srv.Client())

	h := &handler{repository, vesselClient}
	err = pb.RegisterShippingServiceHandler(srv.Server(), h)
	if err != nil {
		log.Panic(err)
	}

	//Run the server
	if err = srv.Run(); err != nil {
		log.Panic(err)
	}

}