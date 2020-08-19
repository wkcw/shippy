package main

import (
	"context"
	"log"
	"os"

	"github.com/micro/go-micro/v2"
	pb "github.com/wkcw/shippy/shippy-service-vessel/proto/vessel"
)


const (
	defaultHost = "datastore:27017"
)

func main(){

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = defaultHost
	}

	client, err := CreateClient(context.Background(), host, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	vesselsCollection := client.Database("shippy").Collection("vessels")

	repo := &MongoRepository{vesselsCollection}

	h := &handler{repo}

	service := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)

	err = pb.RegisterVesselServiceHandler(service.Server(), h)
	if err != nil {
		log.Panic(err)
	}

	//Run the server
	if err = service.Run(); err != nil {
		log.Panic(err)
	}
}
