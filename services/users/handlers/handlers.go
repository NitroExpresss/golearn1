package handlers

import (
	"golearn1/services/users/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// Greeting greets user on index page
func Greeting(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, I am users service!")
}

//GetUserArticles makes a request to articles service
func GetUserArticles(c echo.Context) error {
	ua := new(models.UsrArts)
	id := c.Param("id")
	err := ua.GetUserArticles(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, ua)
}

// GetUser returns firstname field of chosen user
func GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := new(models.Usr)
	user, err := models.GetUser(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSONPretty(http.StatusOK, user, " ")
}

// GetUsers calls model to get list of users
func GetUsers(c echo.Context) error {

	users, err := models.GetUsers()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSONPretty(http.StatusOK, users, " ")
}

// AddUser compiles new user and calls model method
func AddUser(c echo.Context) error {
	user := new(models.Usr)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err := user.AddUser()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
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
	return c.JSON(http.StatusOK, "USER WAS SUCCESSFULLY DELETED")
}

// UpdateUser calls db to update user
func UpdateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := new(models.Usr)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := user.UpdateUser(id); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}
