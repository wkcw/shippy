package main

import (
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
	host := os.Getenv("DB_HOST")

	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)
	defer session.Close()

	if err != nil {
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}

	srv := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	vesselClient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())

	repo = repository
	h = &handler{ , vesselClient}
}