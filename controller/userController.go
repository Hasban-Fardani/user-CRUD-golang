package controller

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hasban-fardani/user-CRUD-go/model"
)

type UserController struct {
	DB *sql.DB
}

func NewUserController(db *sql.DB) *UserController {
	return &UserController{DB: db}
}

func (usr *UserController) Create(c *gin.Context) {
	var data model.RequestCreateUser

	if err := c.BindJSON(&data); err != nil {
		panic(err)
	}

	fmt.Println("bind data berhasil:")
}

func (UserController) Read(c *gin.Context) {

}

func (UserController) Update(c *gin.Context) {

}

func (UserController) Delete(c *gin.Context) {

}
