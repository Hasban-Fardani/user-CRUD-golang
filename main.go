package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hasban-fardani/user-CRUD-go/app"
)

func main() {
	fmt.Println("Hello World!")

	s := gin.New()
	s.Use(gin.Logger(), gin.Recovery())
	app.AddRouter(s)
	s.Run("localhost:8080")
}
