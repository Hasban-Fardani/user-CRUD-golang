package app

import (
	"github.com/gin-gonic/gin"
	"github.com/hasban-fardani/user-CRUD-go/controller"
)

func AddRouter(engine *gin.Engine) {
	db := Connect()
	api := engine.Group("/api")

	// define controller
	userController := controller.NewUserController(db)
	viewsController := controller.NewViewsController(engine)

	// routes views
	engine.GET("/", viewsController.Index)

	// routes API
	api.POST("/register", userController.Create)
}
