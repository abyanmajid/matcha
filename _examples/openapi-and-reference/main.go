package main

import (
	"net/http"

	"github.com/abyanmajid/matcha"
	"github.com/abyanmajid/matcha/ctx"
	"github.com/abyanmajid/matcha/logger"
	"github.com/abyanmajid/matcha/openapi"
	"github.com/abyanmajid/matcha/reference"
)

type LoginRequest struct {
	Email    string `json:"email_address"`
	Password string
}

type LoginResponse struct {
	Token string
}

func LoginHandler(c *ctx.Request[LoginRequest]) *ctx.Response[LoginResponse] {
	return &ctx.Response[LoginResponse]{
		Response: LoginResponse{
			Token: "a secure token",
		},
		StatusCode: http.StatusOK,
		Error:      nil,
	}
}

func LoginResource() (*openapi.Resource, error) {
	requestSchema, err := openapi.NewSchema(LoginRequest{})
	if err != nil {
		return nil, err
	}

	responseSchema, err := openapi.NewSchema(LoginResponse{})
	if err != nil {
		return nil, err
	}

	doc := openapi.ResourceDoc{
		Summary:     "Log in a user by issuing a token",
		Description: "Check if there's a matching user, compare password with its hash, sign and return a JSON Web Token (JWT)",
		Schema: openapi.Schema{
			RequestBody: openapi.RequestBody{
				Content: openapi.Json(requestSchema),
			},
			Responses: map[int]openapi.Response{
				http.StatusOK: {
					Description: "Successfully logged user in",
					Content:     openapi.Json(responseSchema),
				},
				http.StatusUnauthorized: {
					Description: "Invalid credentials.",
					Content:     openapi.Json(openapi.SimpleErrorSchema()),
				},
			},
		},
	}

	resource := openapi.NewResource("Login", doc, LoginHandler)

	return &resource, nil
}

type ApiResources struct {
	Login *openapi.Resource
}

func createApiResources() (*ApiResources, error) {
	login, err := LoginResource()
	if err != nil {
		return nil, err
	}

	return &ApiResources{
		Login: login,
		// ...
	}, nil
}

func main() {
	app := matcha.New()

	app.Documentation("/docs", openapi.Meta{
		OpenAPI:        "3.0.0",
		PackageName:    "My API",
		PackageVersion: "0.1.0",
	})

	app.Reference("/reference", &reference.Options{
		Source: "/docs",
	})

	resources, err := createApiResources()
	if err != nil {
		logger.Fatal("Failed to create resources: %v", err)
	}

	app.Get("/login", resources.Login)

	app.Serve(":8080")
}
