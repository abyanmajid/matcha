<div align="center">
  <img src="https://github.com/user-attachments/assets/185986c2-a956-490b-920d-0e06ebf6204b" width="40%" alt="matcha">
</div>

<div align="center">
  <h1>Matcha: A schema-first REST framework</h1>
</div>

<div align="center">
    Matcha is an ergonomic schema-first REST framework built on top of the 
    <a href="https://github.dev/go-chi/chi">Chi</a> 
    Golang router, designed to overclock your productivity in building type-safe, well-documented, and principled REST APIs.
</div>

## Overview

**Matcha is an opinionated approach to building web servers:** You are encouraged to design and document your API resources before implementing the handler logic. The `MatchaOpenAPI` (`matcha.NewOpenAPI()`) router has chaining path-defining methods (e.g., `GET`) which, for every route, requires you to (1) strictly explicitly define the schema for the request and response payloads, and (2) write an OpenAPI documentation in Go code.

**Opting out of OpenAPI:** You can alternatively use the base `Matcha` (`matcha.New()`) router, which doesn't require you to write OpenAPI documentation for your routes.

## Packages

Being a direct fork, Matcha ([matcha](https://github.com/abyanmajid/matcha)) inherits all functionalities of [Chi](https://github.com/go-chi/chi) and [chi/middleware](https://github.com/go-chi/chi/tree/master/middleware). However, Matcha also wires together much of the tooling you might need to create a production-ready server:

- OpenAPI specification ([matcha/openapi](https://github.com/abyanmajid/matcha/tree/master/openapi)),
- API reference ([matcha/reference](#))
- Input validation ([matcha/validation](https://github.com/abyanmajid/matcha/tree/master/validation))
- In-memory and Redis caching ([matcha/cache](https://github.com/abyanmajid/matcha/tree/master/cache))
- SMTP emails ([matcha/email](https://github.com/abyanmajid/matcha/tree/master/email))
- JSON Web Tokens ([matcha/jwt](https://github.com/abyanmajid/matcha/tree/master/jwt))
- Logging ([matcha/logging](https://github.com/abyanmajid/matcha/tree/master/logging))

## Usage

WIP.

## API Reference

WIP.

## License

Matcha is licensed under GPL 3.0.
