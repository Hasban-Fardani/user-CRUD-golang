package views

import "github.com/gin-gonic/gin"

func InitViews(engine *gin.Engine) {
	engine.LoadHTMLGlob("/*")
}
