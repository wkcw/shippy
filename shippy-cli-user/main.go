package main

import (
	"log"
	"os"

	microclient "github.com/micro/go-micro/client"
	pb "github.com/wkcw/shippy/shippy-service-user/proto/user"
	"golang.org/x/net/context"
)

func main() {

	client := pb.NewUserServiceClient("go.micro.srv.user", microclient.DefaultClient)

	name := "wkcw"
	email := "wkcw@gmail.com"
	password := "test123"
	company := "mycompany"

	r, err := client.Create(context.TODO(), &pb.User{
		Name:     name,
		Email:    email,
		Password: password,
		Company:  company,
	})
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %s", r.User.Id)

	getAll, err := client.GetAll(context.Background(), &pb.request{})
	if err != nil {
		log.Fatalf("Could not list users: %v", err)
	}
	for _, v := range getAll.Users {
		log.Println(v)
	}

	authResponse, err := client.Auth(contexzt.TODO(), &pb.User{
		Email:    email,
		Password: password,
	})

	if err != nil {
		log.Fatalf("Could not authenticate user: %s error: %v\n", email, err)
	}

	log.Printf("Your access token is : %s \n", authResponse.Token)

	os.Exit(0)
}
