package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hasban-fardani/user-CRUD-go/app"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	s := gin.New()
	s.Use(gin.Logger(), gin.Recovery())

	app.AddRouter(s)
	
	fmt.Println("Server Berjalan di http://localhost:5000")
	s.Run("localhost:5000")
}
