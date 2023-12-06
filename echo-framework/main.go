package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/users", GetAllUsers)
	e.GET("/users/:id", GetUserByID)
	e.POST("/users", CreateUser)
	e.PUT("/users/:id", UpdateUser)
	e.DELETE("/users/:id", DeleteUser)
	e.Logger.Fatal(e.Start(":1323"))
}

func GetAllUsers(ctx echo.Context) error {
	search := ctx.QueryParam("search")

	return ctx.JSON(http.StatusOK, map[string]string{"message": "Succesfully Get All Users", "filter": search})
}

func CreateUser(ctx echo.Context) error {
	var input struct {
		Name string `json:"name"`
	}
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Failed to Bind Input"})
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{"message": "Succesfully Create a User", "data": input})
}

func GetUserByID(ctx echo.Context) error {
	id := ctx.Param("id")
	return ctx.JSON(http.StatusOK, map[string]string{"message": fmt.Sprintf("Succesfully Get User By ID : %s", id)})
}

func UpdateUser(ctx echo.Context) error {
	id := ctx.Param("id")
	return ctx.JSON(http.StatusOK, map[string]string{"message": fmt.Sprintf("Succesfully Update User By ID : %s", id)})
}

func DeleteUser(ctx echo.Context) error {
	return ctx.JSON(http.StatusNoContent, nil)
}
