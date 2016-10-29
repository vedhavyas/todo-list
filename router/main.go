package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-web"

	"github.com/vedhavyas/todo-list/auth/proto"
	"github.com/vedhavyas/todo-list/db/proto"
)

func main() {
	// create service
	service := web.NewService(
		web.Name("router"),
		web.Version("1.0"),
		web.Address(":8080"),
	)

	// initiate handlers
	service.Handle("/todos", http.HandlerFunc(todoHandler))
	service.Handle("/", http.HandlerFunc(genericHandler))

	// start service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

// todoHandler handles all the GetTodoList requests
func todoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		// handle all the other requests
		genericHandler(w, r)
		return
	}

	// get the JWT token from the header
	log.Println("Extracting token from request...")
	token := r.Header.Get("X-AUTH-TOKEN")
	if token == "" {
		// no token received from request
		// send as bad request
		log.Println("Token not found in the request...")
		sendErrorResponse(w, "JWT token not found", http.StatusBadRequest)
		return
	}

	// create a new todo request
	todoReq := client.NewRequest("App", "App.GetTodoList", &auth.AuthRequest{
		Token: token,
	})

	// create a new todo response
	todoRes := &db.TodoListResponse{}

	// create a context for the call
	ctx := context.Background()

	if err := client.Call(ctx, todoReq, todoRes); err != nil {
		log.Println(err)
		sendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Convert to JSON
	log.Println("Successfully fetched the todo list...")
	data, err := json.MarshalIndent(todoRes, "", " ")
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Sending the response...")
	w.Write(data)
}

// genericHandler handles all the other requests and responds back 404
func genericHandler(w http.ResponseWriter, r *http.Request) {
	sendErrorResponse(w, "Page not found", http.StatusNotFound)
}

// sendErrorResponse will send back the errormessage with appropriate status
func sendErrorResponse(w http.ResponseWriter, errorMessage string, status int) {
	log.Println("received unexpected request")
	var response struct {
		Error string `json:"error"`
	}

	response.Error = errorMessage

	// marshall response
	data, err := json.MarshalIndent(response, "", " ")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// send response back
	w.WriteHeader(status)
	w.Write(data)
}
