package controller

import (
	"github.com/gin-gonic/gin"
)

type M map[string]string

type ViewsController struct {
	engine *gin.Engine
}

func NewViewsController(engine *gin.Engine) *ViewsController {
	engine.LoadHTMLGlob("public/*")
	return &ViewsController{}
}

func (v *ViewsController) Index(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{"name": "Hasban"})
}
