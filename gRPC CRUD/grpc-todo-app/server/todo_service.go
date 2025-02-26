package server

import (
	"context"
	"fmt"
	"sync"

	"grpc-todo-app/pb"
	"github.com/google/uuid"
)

type TodoService struct {
	pb.UnimplementedTodoServiceServer
	mu    sync.Mutex
	todos map[string]*pb.Todo
}

func NewTodoService() *TodoService {
	return &TodoService{
		todos: make(map[string]*pb.Todo),
	}
}

func (s *TodoService) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (*pb.CreateTodoResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := uuid.New().String()
	todo := &pb.Todo{
		Id:          id,
		Title:       req.Title,
		Description: req.Description,
		Completed:   false,
	}
	s.todos[id] = todo

	fmt.Println("Todo Created:", todo)
	return &pb.CreateTodoResponse{Id: id}, nil
}

func (s *TodoService) GetTodo(ctx context.Context, req *pb.GetTodoRequest) (*pb.GetTodoResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo, exists := s.todos[req.Id]
	if !exists {
		return nil, fmt.Errorf("todo not found")
	}
	return &pb.GetTodoResponse{Todo: todo}, nil
}

func (s *TodoService) ListTodos(ctx context.Context, req *pb.ListTodosRequest) (*pb.ListTodosResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var todoList []*pb.Todo
	for _, todo := range s.todos {
		todoList = append(todoList, todo)
	}

	return &pb.ListTodosResponse{Todos: todoList}, nil
}

func (s *TodoService) UpdateTodo(ctx context.Context, req *pb.UpdateTodoRequest) (*pb.UpdateTodoResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo, exists := s.todos[req.Id]
	if !exists {
		return nil, fmt.Errorf("todo not found")
	}

	todo.Title = req.Title
	todo.Description = req.Description
	todo.Completed = req.Completed
	s.todos[req.Id] = todo

	fmt.Println("Todo Updated:", todo)
	return &pb.UpdateTodoResponse{Success: true}, nil
}

func (s *TodoService) DeleteTodo(ctx context.Context, req *pb.DeleteTodoRequest) (*pb.DeleteTodoResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.todos[req.Id]
	if !exists {
		return nil, fmt.Errorf("todo not found")
	}
	delete(s.todos, req.Id)

	fmt.Println("Todo Deleted:", req.Id)
	return &pb.DeleteTodoResponse{Success: true}, nil
}
