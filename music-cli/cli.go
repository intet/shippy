package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"context"

	pbm "github.com/intet/shippy/consignment-service/proto/playlist"
	"google.golang.org/grpc"
)

const (
	address         = "localhost:50051"
	defaultFilename = "music.json"
)

func parseFile(file string) (*pbm.CreatePlayListRq, error) {
	var consignment *pbm.CreatePlayListRq
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := pbm.NewPlayListServiceClient(conn)

	// Contact the server and print out its response.
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreatePlayList(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Println(r)
}
