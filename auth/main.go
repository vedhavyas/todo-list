package main

import (
	"context"
	"errors"

	"log"

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
	log.Println("Authntication request received...")
	username, err := getUsernameFromToken(req.Token)
	if err != nil {
		log.Println("Authentication failed.")
		return err
	}
	res.Username = username
	return nil
}

// getUsernameFromToken extracts the username from given JWT token
func getUsernameFromToken(token string) (string, error) {
	if token != DummyToken {
		return "", errors.New("Invalid JWT Token")
	}

	return DummyUsername, nil
}
