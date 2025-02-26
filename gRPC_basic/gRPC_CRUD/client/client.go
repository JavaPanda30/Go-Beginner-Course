package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "example.com/crud/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to Connect: %v ", err)
	}
	defer conn.Close()

	client := pb.NewTodoServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	createResp, err := client.CreateTodo(ctx, &pb.CreateTodoRequest{
		Title:       "Buy Something",
		Description: "Buy someother thindd ",
	})
	if err != nil {
		log.Fatalf("CreateTodo failed: %v", err)
	}

	fmt.Println("Todo Created with ID:", createResp.Id)

	listResp, err := client.ListTodo(ctx, &pb.ListTodoRequest{})
	if err != nil {
		log.Fatalf("ListTodo failed: %v", err)
	}
	fmt.Println("List of Todo:", listResp.Todos)
}
