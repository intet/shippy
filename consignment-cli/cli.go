package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"context"

	pb "github.com/intet/shippy/consignment-service/proto/consignment"
	micro "github.com/micro/go-micro"
)

const (
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	log.Println("try to start")

	// Set up a connection to the server.
	service := micro.NewService(micro.Name("shippy.consignment.cli"))
	service.Init()
	client := pb.NewShippingServiceClient("shippy.service.consignment", service.Client())

	// Contact the server and print out its response.
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	getAll, err := client.GetConsignment(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}

	log.Printf("Created: %t", r.Created)
}
