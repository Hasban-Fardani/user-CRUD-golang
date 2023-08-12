package model

type UserQuery struct {
	Name  string `form:"name" json:"name"`
	Email string `form:"email" json:"email"`
}
