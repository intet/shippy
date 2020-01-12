build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t consignment-cli .

run:
	docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns consignment-cli
