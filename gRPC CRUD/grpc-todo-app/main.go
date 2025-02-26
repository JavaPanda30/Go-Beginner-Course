package main

import (
	"fmt"
	"log"
	"net"

	pb "grpc-todo-app/pb"
	"grpc-todo-app/server"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	todoService := server.NewTodoService()
	pb.RegisterTodoServiceServer(grpcServer, todoService)

	fmt.Println("gRPC server started on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
