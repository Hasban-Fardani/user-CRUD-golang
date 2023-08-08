package app

import (
	"github.com/gin-gonic/gin"
)

func AddRouter(engine *gin.Engine) {
	engine.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
}
