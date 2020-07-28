package models

import (
	"context"
	"golearn1/services/users/config"

	"github.com/go-pg/pg/v10"
)

// DB Exportable database instance
var db *pg.DB

// Init DB
func Init() (*pg.DB, error) {
	db = pg.Connect(&pg.Options{
		Addr:     config.S.Host + ":" + config.S.Port,
		User:     config.S.Dbuser,
		Password: config.S.Dbpassword,
		Database: config.S.Name,
	})
	// Check if connection credentials are valid and PostgreSQL is up and running.
	err := db.Ping(context.Background())
	return db, err
}
