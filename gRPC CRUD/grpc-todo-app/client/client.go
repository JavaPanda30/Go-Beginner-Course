package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "grpc-todo-app/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewTodoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	createResp, err := client.CreateTodo(ctx, &pb.CreateTodoRequest{
		Title:       "Buy Groceries",
		Description: "Milk, Eggs, Bread",
	})
	if err != nil {
		log.Fatalf("CreateTodo failed: %v", err)
	}

	fmt.Println("Todo Created with ID:", createResp.Id)

	// Get todos
	listResp, err := client.ListTodos(ctx, &pb.ListTodosRequest{})
	if err != nil {
		log.Fatalf("ListTodos failed: %v", err)
	}
	fmt.Println("List of Todos:", listResp.Todos)
}
