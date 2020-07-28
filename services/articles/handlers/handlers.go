package handlers

import (
	"golearn1/services/articles/models"
	"net/http"

	"strconv"

	"github.com/labstack/echo"
)

// Greeting greets Article on index page
func Greeting(c echo.Context) error {
	return c.String(http.StatusOK, "Hello!articles")
}

// GetArticles calls model to get list of Articles
func GetArticles(c echo.Context) error {
	articles, err := models.GetArticles()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSONPretty(http.StatusOK, articles, " ")
}

//GetUserArticles calls model to get list of Articles for user
func GetUserArticles(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	articles, err := models.GetUserArticles(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSONPretty(http.StatusOK, articles, " ")
}

// GetArticle returns firstname field of chosen Article
func GetArticle(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	art := new(models.Art)
	art, err := models.GetArticle(id)

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, art)
}

// AddArticle compiles new article and calls model method
func AddArticle(c echo.Context) error {
	art := new(models.Art)
	if err := c.Bind(art); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err := art.AddArticle()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, art)
}

// DeleteArticle calls model method to delete Article from db
func DeleteArticle(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.DeleteArticle(id)
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}
	return c.String(http.StatusOK, "Article WAS SUCCESSFULLY DELETED")
}

// UpdateArticle calls db to update Article
func UpdateArticle(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	art := new(models.Art)

	if err := c.Bind(art); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	art.UpdateArticle(id)

	return c.JSON(http.StatusOK, art)
}
