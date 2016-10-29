package main

import (
	"log"

	"github.com/micro/go-micro/client"
	"github.com/vedhavyas/todo-list/app/proto"
	"golang.org/x/net/context"
)

func main() {

	req := client.NewRequest("App", "App.Get", &app.ListRequest{
		Name: "Vedhavyas",
	})

	ctx := context.Background()

	res := &app.ListResponse{}

	if err := client.Call(ctx, req, res); err != nil {
		log.Fatal(err)
	}

	log.Println(res.Data)

}
