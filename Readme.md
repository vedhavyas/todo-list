#Todo List using Go-Micro

This project contains 4 micro-services built using go-micro.

1. Router - Router route the requests to appropriate service(App) for results.
2. App - App authenticates the given request and respond back with todo-list of the user.
3. Auth - Auth authenticates a given service.
4. DB - DB responds back with all the todos for a given user

## pre-requisites
1. Golang 1.6+
2. Consul 
3. Rest Client (Postman - https://www.getpostman.com/) or curl

## Starting consul
1. Install Consul from - https://www.consul.io/intro/getting-started/install.html
2. Run consul - `consul agent -dev -advertise=127.0.0.1`

## Starting all the 4 services
1. open 4 terminals
2. `cd` to root of the project in each terminal

### Run the following 4 commands one in each terminal
1. `go run router/main.go`
2. `go run app/main.go`
3. `go run auth/main.go`
4. `go run db/main.go`

## Sending the request
1. Only one endpoint `GET /todos` is exposed for the entire service. This endpoint can be used to fetch the todos of the user.
2. Request sent to the endpoint should contain a JWT token in the request header with following key - `X-AUTH-TOKEN`
3. For testing, JWT token `123456789` is hardcoded to user `vedhavyas`

### Using Rest client
If using a rest client like postman, send a `GET` request to `localhost:8080` with JWT token in the header.

### Using curl
If using curl, simply copy the curl command to fetch the results 

`curl -X GET -H "X-AUTH-TOKEN: 123456789" "http://localhost:8080/todos"`
