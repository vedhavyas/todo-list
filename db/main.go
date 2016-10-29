package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/micro/go-micro/server"

	"github.com/vedhavyas/todo-list/auth/proto"
	"github.com/vedhavyas/todo-list/db/proto"
)

// Todo holds the details of a single todo
type Todo struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

// DB implements the rpc GetTodoList
type DB struct{}

// GetTodoList returns the todo-list of the given user
func (d *DB) GetTodoList(ctx context.Context, req *auth.AuthResponse, res *db.TodoListResponse) error {
	log.Println("Request for GetTodoList for user - " + req.Username)
	todoList, err := getTodoListOfUser(req.Username)
	if err != nil {
		log.Println("Failed to fetch the todolist..")
		return err
	}

	res.Username = req.Username
	for _, todo := range todoList {
		protoTodo := &db.TodoList{
			TodoName: todo.Name,
			Status:   todo.Status,
		}

		res.TodoList = append(res.TodoList, protoTodo)
	}
	log.Println("Successfully fetched todolist...")

	return nil
}

// getTodoListOfUser will return the todo-list of the given user
func getTodoListOfUser(username string) ([]Todo, error) {
	fileName := "db/todo-list.json"
	fh, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	var todoList []Todo
	data, err := ioutil.ReadAll(fh)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &todoList)
	if err != nil {
		return nil, err
	}

	return todoList, nil
}

func main() {
	// init the server
	server.Init(
		server.Name("DB"),
		server.Version("1.0"),
	)

	// register handler
	server.Handle(server.NewHandler(&DB{}))

	// start the server
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}
