package controllers

import (
	"fmt"
	"net/http"
	"revel-web-api/app/models"

	"github.com/revel/revel"
)

type User struct {
	*revel.Controller
}

func (c User) Index() revel.Result {
	return c.RenderText("GET /users")
}

func (c User) SignUp() revel.Result {

	var newUser models.UserModel
	err := c.Params.BindJSON(&newUser)

	if err != nil {
		c.Response.SetStatus(http.StatusBadRequest)
		return c.RenderText("JSON cannot be parsed")
	}

	fmt.Println(newUser)

	return c.RenderText("GET /users/signup")
}
