package main

import (
	"golearn1/services/users/config"
	"golearn1/services/users/models"
	"golearn1/services/users/routes"
)

func main() {
	config.Init()
	e := routes.Init()
	db, err := models.Init()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	e.Logger.Fatal(e.Start(":1323"))
}
