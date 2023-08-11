package controller

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hasban-fardani/user-CRUD-go/helper"
	"github.com/hasban-fardani/user-CRUD-go/model"
)

type UserController struct {
	DB *sql.DB
}

func (usr *UserController) GetLastId() int {
	var lastId int

	res := usr.DB.QueryRow("SELECT MAX(id) FROM Users;")
	res.Scan(&lastId)

	return lastId
}

func NewUserController(db *sql.DB) *UserController {
	return &UserController{DB: db}
}

func (usr *UserController) Create(c *gin.Context) {
	var data model.RequestCreateUser
	var id int

	id = usr.GetLastId() + 1

	if err := c.BindJSON(&data); err != nil {
		helper.HandleErr(err, c)
	}

	_, err := usr.DB.Exec(
		"INSERT INTO Users (id, name, email, password) values (?, ?, ?, ?)",
		id, data.Name, data.Email, data.Password)

	if err != nil {
		helper.HandleErr(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
		"id":  id,
	})
}

func (usr *UserController) GetByID(c *gin.Context) {
	id := c.Param("id")
	user := model.User{}

	res := usr.DB.QueryRow("SELECT * FROM Users where id=?", id)
	err := res.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		helper.HandleErr(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": user,
	})
}

func (usr *UserController) GetAll(c *gin.Context) {
	var allUsr []model.User
	var usrTmp model.User

	rows, err := usr.DB.Query("SELECT * FROM Users")

	if err != nil {
		helper.HandleErr(err, c)
		return
	}

	for rows.Next() {
		if err := rows.Err(); err != nil {
			helper.HandleErr(err, c)
			return
		}
		rows.Scan(&usrTmp.Id, &usrTmp.Name, &usrTmp.Email, &usrTmp.Password, &usrTmp.CreatedAt, &usrTmp.UpdatedAt)
		allUsr = append(allUsr, usrTmp)
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": allUsr,
	})
}

func (UserController) Update(c *gin.Context) {

}

func (UserController) Delete(c *gin.Context) {

}
