package controller

import (
	"net/http"
	"strconv"
	"www/src/model"

	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
	i, _ := strconv.Atoi(c.Param("id"))
	id := uint(i)
	user := model.User{}
	user.FirstById(id)

	return c.JSON(http.StatusOK, user)
}

func CreateUser(c echo.Context) error {
	name := c.FormValue("name")
	a, _ := strconv.Atoi(c.FormValue("age"))
	age := uint(a)

	user := model.User{
		Name: name,
		Age:  age,
	}
	user.Create()

	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	i, _ := strconv.Atoi(c.Param("id"))
	id := uint(i)
	name := c.FormValue("name")
	a, _ := strconv.Atoi(c.FormValue("age"))
	age := uint(a)

	user := model.User{
		ID:   id,
		Name: name,
		Age:  age,
	}
	user.Updates()

	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	i, _ := strconv.Atoi(c.Param("id"))
	id := uint(i)
	user := model.User{}
	user.DeleteById(id)

	return c.JSON(http.StatusOK, user)
}
