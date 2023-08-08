package controller

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	DB *sql.DB
}

func NewUserController(db *sql.DB) *UserController {
	return &UserController{DB: db}
}

func (UserController) Create(c *gin.Context) {

}

func (UserController) Read(c *gin.Context) {

}

func (UserController) Update(c *gin.Context) {

}

func (UserController) Delete(c *gin.Context) {

}
