package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hasban-fardani/user-CRUD-go/app"
)

var server *gin.Engine

func init() {
	gin.SetMode(gin.ReleaseMode)
	server = gin.New()
	server.Use(gin.Logger(), gin.Recovery())
}

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println(os.Chdir("."))
	app.AddRouter(server)
	server.ServeHTTP(w, r)
}
