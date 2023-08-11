package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleErr(err error, c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"msg": err.Error(),
	})
}
