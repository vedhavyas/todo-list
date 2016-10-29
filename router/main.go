package main

import (
	"context"
	"log"

	"github.com/micro/go-micro/client"

	"encoding/json"
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

	data, _ := json.Marshal(res)
	log.Println(string(data))
	log.Println(res.Data)

}
