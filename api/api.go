package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hasban-fardani/user-CRUD-go/app"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := gin.Default()
	app.AddRouter(server)
	server.ServeHTTP(w, r)
}
