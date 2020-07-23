package handlers

import (
	"golearn/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// Greeting greets user on index page
func Greeting(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, this is crooked representation of CRUD on Golang! Enjoy")
}

// GetUser returns firstname field of chosen user
func GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := new(models.User)
	user, err := models.GetUser(id)

	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

// GetUsers calls model to get list of users
func GetUsers(c echo.Context) error {

	users, err := models.GetUsers()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

// AddUser compiles new user and calls model method
func AddUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	err := user.AddUser()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

// DeleteUser calls model method to delete user from db
func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.DeleteUser(id)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "USER WAS SUCCESSFULLY DELETED")
}

// UpdateUser calls db to update user
func UpdateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := new(models.User)

	if err := c.Bind(user); err != nil {
		return err
	}
	models.UpdateUser(user, id)

	return c.JSON(http.StatusOK, user)
}
