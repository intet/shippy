package main

import (
	"context"
	"log"
	"net"
	"sync"

	// Import the generated protobuf code
	pb "github.com/intet/shippy/consignment-service/proto/consignment"
	pbm "github.com/intet/shippy/consignment-service/proto/playlist"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// Repository - Dummy repository, this simulates the use of a datastore
// of some kind. We'll replace this with a real implementation later on.
type Repository struct {
	mu           sync.RWMutex
	consignments []*pb.Consignment
}

// Create a new consignment
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.mu.Lock()
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	repo.mu.Unlock()
	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type service struct {
	repo repository
}

// CreateConsignment - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {

	// Save our consignment
	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}
	log.Println("some:", req.Weight)

	// Return matching the `Response` message we created in our
	// protobuf definition.
	return &pb.Response{Created: true, Consignment: consignment}, nil
}
func (s *service) GetConsignment(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	all := s.repo.GetAll()
	return &pb.Response{Created: false, Consignments: all}, nil
}

type musicRepository interface {
	create(rq *pbm.CreatePlayListRq) (*pbm.CreatePlayListRs, error)
}

type MusicRepository struct {
	mu       sync.RWMutex
	playList []*pbm.CreatePlayListRq
}

func (repo *MusicRepository) create(rq *pbm.CreatePlayListRq) (*pbm.CreatePlayListRs, error) {
	repo.mu.Lock()
	updated := append(repo.playList, rq)
	repo.playList = updated
	repo.mu.Unlock()
	return &pbm.CreatePlayListRs{Name: rq.Name, Size: int32(len(rq.Tracks))}, nil
}

type playListService struct {
	repo musicRepository
}

func (s *playListService) CreatePlayList(ctx context.Context, req *pbm.CreatePlayListRq) (*pbm.CreatePlayListRs, error) {
	result, err := s.repo.create(req)
	if err != nil {
		return nil, err
	}
	log.Println("some:", req.Name)

	return result, nil
}

func main() {

	repo := &Repository{}
	musicRepo := &MusicRepository{}

	// Set-up our gRPC server.
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// Register our service with the gRPC server, this will tie our
	// implementation into the auto-generated interface code for our
	// protobuf definition.
	pb.RegisterShippingServiceServer(s, &service{repo})
	pbm.RegisterPlayListServiceServer(s, &playListService{musicRepo})
	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.Println("Running on port:", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
