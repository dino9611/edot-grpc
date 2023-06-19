xprotos:
	protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. proto/*.proto

build-grpc-server:
	go build -o ./build/grpcserver ./grpcserver/cmd/main.go

build-gateway:
	go build -o ./build/gateway ./gateway/cmd/main.go

dev-grpc-server:
	nodemon --exec go run ./grpcserver/cmd/main.go --signal SIGTERM

dev-gateway:
	nodemon --exec go run ./gateway/cmd/main.go --signal SIGTERM

start-grpc-server:
	go run ./grpcserver/cmd/main.go 

start-gateway:
	go run ./gateway/cmd/main.go 

run-grpcserver:
	./build/grpcserver

run-gateway:
	./build/gateway
