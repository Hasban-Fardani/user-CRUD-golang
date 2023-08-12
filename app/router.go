package app

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hasban-fardani/user-CRUD-go/controller"
)

func AddRouter(engine *gin.Engine) {
	db := Connect()
	api := engine.Group("/api")

	// define controller
	userController := controller.NewUserController(db)

	// it's not working on vervel but working on local
	if v := os.Getenv("ON_VERCEL"); v == "" {
		viewsController := controller.NewViewsController(engine)

		// routes views
		engine.GET("/", viewsController.Index)
	}

	// routes API
	api.POST("/register", userController.Create)
	api.GET("/users", userController.GetAll)
	api.GET("/users/:id", userController.GetByID)
	api.PUT("/users/:id", userController.Update)
	api.DELETE("/users/:id", userController.DeleteByID)
}
