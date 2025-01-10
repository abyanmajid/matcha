package main

import (
	"net/http"

	"github.com/abyanmajid/matcha"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type ApiResources struct {
	login *matcha.Resource
}

func createApiResources() *ApiResources {
	return &ApiResources{
		login: matcha.NewResource("/login", LoginHandler),
	}
}

func LoginHandler(r *http.Request, body LoginRequest) *matcha.Response[LoginResponse] {
	return &matcha.Response[LoginResponse]{
		Response: LoginResponse{
			Token: "hello",
		},
	}
}

func main() {
	mux := matcha.New()

	resources := createApiResources()

	mux.Post(resources.login.Route, resources.login.Handler)
}
