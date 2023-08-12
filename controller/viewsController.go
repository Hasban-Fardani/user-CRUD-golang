package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hasban-fardani/user-CRUD-go/views"
)

type M map[string]string

type ViewsController struct {
	engine *gin.Engine
}

func NewViewsController(engine *gin.Engine) *ViewsController {
	views.InitViews(engine)
	return &ViewsController{}
}

func (v *ViewsController) Index(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{"name": "Hasban"})
}
