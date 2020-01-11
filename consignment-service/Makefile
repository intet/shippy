build:
	protoc -I. --go_out=plugins=grpc:. \
	  proto/consignment/consignment.proto
	protoc -I. --go_out=plugins=grpc:. \
	  proto/playlist/playlist.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t consignment-service .

run:
    docker run -p 50051:50051 shippy-service-consignment
