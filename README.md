<div align="center">
  <img src="https://github.com/user-attachments/assets/8f032a92-491e-45fd-9aff-a8e34761e9a5" width="50%" alt="matcha">
</div>

<div align="center">
  <h1>Matcha: A Schema-First REST Framework</h1>
</div>

<div align="center">
    Matcha is an ergonomic schema-first REST framework built on top of the 
    <a href="https://github.dev/go-chi/chi">Chi</a> 
    Golang router, designed to overclock your productivity in building type-safe, well-documented, and principled REST APIs.
</div>

<br>

<div align="center">
  <a href="https://matcha.abydyl.net" target="_blank" rel="noopener noreferrer">
  <img
    src="https://img.shields.io/badge/DOCS-MATCHA.ABYDYL.NET-555555?style=for-the-badge&logo=read-the-docs&logoColor=white&labelColor=333333&color=777777"
    alt="Docs"
  />
</a>
</div>

<br>

## Overview

**Matcha is an opinionated approach to building web servers:** You are encouraged to design and document your API resources before implementing the handler logic. The `MatchaOpenAPI` (`matcha.New()`) router has chaining path-defining methods (e.g., `GET`) which, for every route, requires you to (1) strictly explicitly define the schema for the request and response payloads, and (2) write an OpenAPI documentation in Go code.

**Opting out of OpenAPI:** You can alternatively use the base `Matcha` (`matcha.NewBase()`) router, which doesn't require you to write OpenAPI documentation for your routes.

## Packages

Being a direct fork, Matcha ([matcha](https://github.com/abyanmajid/matcha)) inherits all functionalities of [Chi](https://github.com/go-chi/chi) and [chi/middleware](https://github.com/go-chi/chi/tree/master/middleware). However, Matcha also wires together much of the tooling you might need to create a production-ready server:

- OpenAPI specification builder ([matcha/openapi](https://github.com/abyanmajid/matcha/tree/master/openapi)),
- API reference using Scalar ([matcha/reference](https://github.com/abyanmajid/matcha/tree/master/reference))
- In-memory and Redis caching ([matcha/cache](https://github.com/abyanmajid/matcha/tree/master/cache))
- SMTP emails ([matcha/email](https://github.com/abyanmajid/matcha/tree/master/email))
- Security; JSON Web Tokens, encryption, and hashing ([matcha/security](https://github.com/abyanmajid/matcha/tree/master/security))
- Type-safe Env ([matcha/env](https://github.com/abyanmajid/matcha/tree/master/env))
- Logging ([matcha/logger](https://github.com/abyanmajid/matcha/tree/master/logger))

## Usage

Add Matcha to your Go project:

```
go get -u github.com/abyanmajid/matcha
```

Create a new `Matcha` instance and register (1) a handler serving our OpenAPI documentation, and (2) a handler serving our API reference, which points to the OpenAPI handler as the content soure:

```go
package main

import (
  "github.com/abyanmajid/matcha"
  "github.com/abyanmajid/matcha/openapi"
  "github.com/abyanmajid/matcha/reference"
)

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

  // We'll define our resources here
  // ...

  app.Serve(":8080")
}
```

Let's kickoff our API by creating a user login resource. The Matcha multiplexer requires you to define an OpenAPI documentation for each route. You can compose such documentation purely in Go code.

Start by defining the request and response schema:

```go
// The request body schema
type LoginRequest struct {
  // Note: Matcha will use a json tag as the
  // property name in the OpenAPI specification.
  Email string `json:"email_address"`
  // If you don't provide any, Matcha will parse
  // the field name as is e.g., "password"
  Password string
}

// The successful response schema
//
// You need this success case schema for the
// Matcha handler, however you can choose to
// define a schema for other status codes as well
type LoginResponse struct {
  Token string
}
```

Create a new login resource, returning an `Resource` instance from the [matcha/openapi](#) package. In it, you'll want to start by parsing your types into a `ContentSchema`:

```go
import (
  // ...
  "github.com/abyanmajid/matcha/openapi"
  // ...
)

func LoginResource() (*openapi.Resource, error) {
  requestSchema, err := openapi.NewSchema(LoginRequest{})
  if err != nil {
    return nil, err
  }

  responseSchema, err := openapi.NewSchema(LoginResponse{})
  if err != nil {
    return nil, err
  }

  // ...
}
```

Onto defining the documentation!

```go
import (
  // ...
  "github.com/abyanmajid/matcha/ctx"
  // ...
)

func LoginResource() (*openapi.Resource, error) {
  // ...

  doc := openapi.ResourceDoc{
    Summary: "Log in a user by issuing a token",
    Description: "Check if there's a matching user, compare password with its hash, sign and return a JSON Web Token (JWT)",
    Schema: openapi.Schema{
      RequestBody: openapi.RequestBody{
        Content: openapi.Json(requestSchema),
      },
      Responses: map[int]openapi.Response{
        http.StatusOK: {
          Description: "Successfully logged user in"
          Content: openapi.Json(responseSchema),
        },
        http.StatusUnauthorized: {
          Description: "Invalid credentials.",
          Content: openapi.Json(openapi.SimpleErrorSchema()),
        },
      },
    },
  }

  //...
}
```

Create your typed handler:

```go
func LoginHandler(c *ctx.Request[LoginRequest]) *ctx.Response[LoginResponse] {
  // specify your handler logic here...
  return &ctx.Response[LoginResponse]{
    Response: LoginResponse{
      Token: "a secure token",
    },
    StatusCode: http.StatusOK,
    Error: nil,
  }
}
```

Compile everything into a new `Resource`:

```go
func LoginResource() (*openapi.Resource, error) {
  // ...

  resource := openapi.NewResource("Login", doc, LoginHandler)

  return &resource, nil
}
```

Now let's make this resource available by passing it to the router. As an aside, you might also find it useful to define a utility function to populate all of your resources:

```go
type ApiResources struct {
  Login *openapi.Resource
}

func createApiResources() (*ApiResources, error) {
  login, err := LoginResource()
  if err != nil {
    return nil, err
  }

  // Instantiate your other resources here
  // ...

  return &ApiResources{
    Login: login,
    // ...
  }, nil
}
```

Finally, create a `POST` route and pass in our resource!

```go
import (
  // ...
  "github.com/abyanmajid/matcha/logger"
  // ...
)

func main() {
  // ...

  resources, err := createApiResources()
  if err != nil {
    logger.Fatal("Failed to create resources: %v", err)
  }

  app.Post("/login", resources.Login)

  // ...
}
```

At this point, you should've created a resource at `/login` of which specification is documented at `/docs` and can be viewed from a reference interface at `/reference`. A rather elegant design, I hope you agree! 😎

## License

Matcha is licensed under GPL 3.0.
