package routes

import (
	"golearn1/services/articles/handlers"

	"github.com/labstack/echo"
)

// Init echo instance
func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", handlers.Greeting)
	e.GET("/arts", handlers.GetArticles)
	e.GET("/arts/usr/:id", handlers.GetUserArticles)
	e.GET("/arts/add", handlers.AddArticle)
	e.GET("/arts/:id", handlers.GetArticle)
	e.GET("/arts/del/:id", handlers.DeleteArticle)
	e.GET("/arts/upd/:id", handlers.UpdateArticle)
	return e
}
