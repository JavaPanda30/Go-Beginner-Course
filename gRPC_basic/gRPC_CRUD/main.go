package main

import (
	"fmt"
	"log"
	"net"

	"example.com/crud/pb"
	"example.com/crud/server"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Fail to Listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	todoService := server.NewTodoService()
	pb.RegisterTodoServiceServer(grpcServer, todoService)

	fmt.Println("Server Started")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve: %v", err)
	}
}
