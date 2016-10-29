package main

import (
	"context"
	"errors"
	"log"

	"github.com/micro/go-micro/server"

	"github.com/vedhavyas/todo-list/auth/proto"
)

// DummyToken is dummy value for a predefined JWT Token
const DummyToken = "123456789"

// DummyUsername is the dummy value supposedly extracted from JWT token
const DummyUsername = "vedhavyas"

// Auth implements the rpc interface
type Auth struct{}

// Authenticate authenticates the given JWT token
func (a *Auth) Authenticate(ctx context.Context, req *auth.AuthRequest, res *auth.AuthResponse) error {
	log.Println("Authentication request received...")
	username, err := extractUsernameFromToken(req.Token)
	if err != nil {
		log.Println("Authentication failed.")
		return err
	}

	log.Println("Authentication successful.")
	res.Username = username
	return nil
}

// getUsernameFromToken extracts the username from given JWT token
func extractUsernameFromToken(token string) (string, error) {
	if token != DummyToken {
		return "", errors.New("Invalid JWT Token")
	}

	return DummyUsername, nil
}

func main() {
	// init the server
	server.Init(
		server.Name("Auth"),
		server.Version("1.0"),
	)

	// register handler
	server.Handle(server.NewHandler(&Auth{}))

	// start the server
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}
