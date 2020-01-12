build:
	protoc -I. --go_out=plugins=micro:. \
	  proto/consignment/consignment.proto
	protoc -I. --go_out=plugins=micro:. \
	  proto/playlist/playlist.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t consignment-service-intetm .

run:
	docker run -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns consignment-service-intetm
