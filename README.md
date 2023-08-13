# GO CRUD Users

Table of contents

- [GO CRUD Users](#go-crud-users)
  - [About](#about)
  - [Getting Started](#getting-started)
  - [endpoint](#endpoint)

## About
This is a simple sample code to create a CRUD (Create, Read, Update, Delete) API using [Go](https://go.dev/), [gin framework](https://gin-gonic.com/), and [Go-Mysql-Driver](https://pkg.go.dev/github.com/go-sql-driver/mysql) .


## Getting Started
- Download [Go](https://go.dev/doc/install)
- Setup [GOPATH](https://dasarpemrogramangolang.novalagung.com/A-gopath-dan-workspace.html)
- clone The Project
  ```bash
  git clone https://github.com/Hasban-Fardani/user-CRUD-golang.git
  ```
- Run command:
  ```bash
  go run main.go
  ```

## endpoint
- GET
  - /api/users
  - /api/users/:id
- POST
  - /api/users
- PUT
  - /api/users/:id
- DELETE
  - /api/users/:id