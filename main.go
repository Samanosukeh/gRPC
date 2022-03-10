package gRPC

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

// protoc --go_out=plugins=grpc:chat --go_opt=paths=source_relative chat.proto
// protoc --go_out=plugins=grpc:chat chat.proto
// protoc --go_out=plugins=grpc:chat --go_opt=paths=source_relative chat.proto
// protoc --go_out=plugins=grpc:./gen --go_opt=paths=source_relative chat.proto
// protoc --go_out=plugins=grpc:proto/gen --go_opt=paths=source_relative proto/user.proto
