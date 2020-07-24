package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	micro "github.com/micro/go-micro/v2"
	pb "github.com/wkcw/shippy/shippy-service-consignment"
	"golang.org/x/net/context"
)

const (
	address = "localhost:50051"
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error){
	var consignment * pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json_err := json.Unmarshal(data, &consignment)
	if json_err != nil {
		return nil, json_err
	}
	return consignment, err
}

func main() {

	service := micro.NewService(micro.Name("shippy.consignment.cli"))
	service.Init()

	client := pb.NewShippingService("shippy.consignment.service", service.Client())


	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.TODO(), consignment)

	if err != nil {
		log.Fatalf("Could not create consignment: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}

