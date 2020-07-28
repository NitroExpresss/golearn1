package routes

import (
	"golearn1/services/users/handlers"

	"github.com/labstack/echo"
)

// Init echo instance
func Init() *echo.Echo {
	e := echo.New()
	e.GET("/users", handlers.GetUsers)
	e.GET("/users/add", handlers.AddUser)
	e.GET("/users/:id", handlers.GetUser)
	e.GET("/users/upd/:id", handlers.UpdateUser)
	e.GET("/users/del/:id", handlers.DeleteUser)
	e.GET("/", handlers.Greeting)
	e.GET("/users/articles/:id", handlers.GetUserArticles)
	return e
}
