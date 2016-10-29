package main

import (
	"context"
	"log"

	"github.com/micro/go-micro/server"
	"github.com/vedhavyas/todo-list/app/proto"
)

type App struct{}

func (a *App) Get(ctx context.Context, req *app.ListRequest, res *app.ListResponse) error {
	res.Data = "Hello, World!!"
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
