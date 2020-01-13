package main

import (

	"fmt"

	// Import the generated protobuf code
	pb "github.com/intet/shippy/consignment-service/proto/consignment"
	"github.com/micro/go-micro"
	"context"
	"log"
)

const (
	port = ":50051"
	defaultHost = "localhost:27017"
)

func main() {
	// Set-up micro instance
	srv := micro.NewService(
		micro.Name("shippy.service.consignment"),
	)

	srv.Init()

	uri := defaultHost
	if uri == "" {
		uri = defaultHost
	}
	client, err := CreateClient(uri)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.TODO())

	consignmentCollection := client.Database("shippy").Collection("consignments")

	repository := &MongoRepository{consignmentCollection}
	h := &handler{repository}

	// Register handlers
	pb.RegisterShippingServiceHandler(srv.Server(), h)

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
