package main

import (
	"golearn1/services/articles/config"
	"golearn1/services/articles/models"
	"golearn1/services/articles/routes"
)

func main() {
	config.Init()
	e := routes.Init()
	db, err := models.Init()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	e.Logger.Fatal(e.Start(":1324"))
}
