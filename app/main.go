package main

import (
	"context"
	"log"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"

	"github.com/vedhavyas/todo-list/auth/proto"
	"github.com/vedhavyas/todo-list/db/proto"
)

// App implements the rpc GetTodoList
type App struct{}

// GetTodoList authenticates the request and returns back the response
func (a *App) GetTodoList(ctx context.Context, req *auth.AuthRequest, res *db.TodoListResponse) error {
	log.Println("GetTodolist request received...")

	// Authenticate the request
	authReq := client.NewRequest("Auth", "Auth.Authenticate", req)
	authRes := &auth.AuthResponse{}
	if err := client.Call(ctx, authReq, authRes); err != nil {
		return err
	}
	log.Println("Authentication Successful...")

	// Get todoList
	todoReq := client.NewRequest("DB", "DB.GetTodoList", authRes)
	if err := client.Call(ctx, todoReq, res); err != nil {
		return err
	}
	log.Println("Fetched todo list succesfully...")

	return nil
}

func main() {
	server.Init(
		server.Name("App"),
		server.Version("1.0"),
	)

	server.Handle(server.NewHandler(&App{}))
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
