package main

import (
	"golearn/config"
	"golearn/models"
	"golearn/routes"
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
