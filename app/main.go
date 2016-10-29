package main

import (
	"context"
	"log"

	"github.com/micro/go-micro"

	"github.com/vedhavyas/todo-list/app/proto"
)

type TodoList struct{}

func (t *TodoList) Get(ctx context.Context, req *app.ListRequest, res *app.ListResponse) error {
	res.Data = "Hello, World!!"
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("App"),
		micro.Version("1.0"),
	)

	service.Init()
	service.Server().Handle(service.Server().NewHandler(&TodoList{}))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
