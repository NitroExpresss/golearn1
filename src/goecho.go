package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

// User from db
type User struct {
	ID        int
	Age       int
	FirstName string
	LastName  string
	Email     string
}

const (
	dbhost     = "localhost"
	dbport     = 5432
	dbuser     = "postgres"
	dbpassword = "6655"
	dbname     = "gotesting"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	dbhost, dbport, dbuser, dbpassword, dbname)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.GET("/users", getUsers)
	e.GET("/users/add", addUser)
	e.GET("/users/:id", getUser)
	e.GET("/users/upd/:id", updateUser)
	e.GET("/users/del/:id", deleteUser)
	e.GET("/", greeting)

	e.Logger.Fatal(e.Start(":1323"))

}

func getUsers(c echo.Context) error {

	users := []User{}
	rows, _ := db.Query("SELECT * FROM users ORDER BY id ASC")
	defer rows.Close()
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.ID, &user.Age, &user.FirstName, &user.LastName, &user.Email)
		if err != nil {
			// handle this error
			panic(err)
		}
		users = append(users, user)
	}
	// get any error encountered during iteration
	err := rows.Err()
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, users)
}
func greeting(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, this is crooked representation of CRUD on Golang! Enjoy")
}

func getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var name string
	sqlStatement := `SELECT first_name FROM users WHERE id=$1;`
	row := db.QueryRow(sqlStatement, id)
	switch err := row.Scan(&name); err {
	case sql.ErrNoRows:
		return c.String(http.StatusOK, "No rows was returned from db")
	case nil:
		return c.String(http.StatusOK, "The name you requested is: "+name)
	default:
		return c.String(http.StatusOK, err.Error())
	}
}

func addUser(c echo.Context) error {
	user := User{}
	user.Age, _ = strconv.Atoi(c.QueryParam("age"))
	user.Email = c.QueryParam("email")
	user.FirstName = c.QueryParam("first_name")
	user.LastName = c.QueryParam("last_name")
	sqlStatement := `
	INSERT INTO users (age, email, first_name, last_name)
	VALUES ($1, $2, $3, $4)`
	db.Exec(sqlStatement, user.Age, user.Email, user.FirstName, user.LastName)
	return c.String(http.StatusOK, "USER WAS SUCCESSFULLY ADDED")
}

func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	sqlStatement := `
		DELETE FROM users
		WHERE id = $1;`
	db.Exec(sqlStatement, id)
	return c.String(http.StatusOK, "USER WAS SUCCESSFULLY DELETED")
}
func updateUser(c echo.Context) error {
	name := c.QueryParam("first_name")
	lastName := c.QueryParam("last_name")
	id, _ := strconv.Atoi(c.Param("id"))
	sqlStatement := `
		UPDATE users
		SET first_name = $2, last_name = $3
		WHERE id = $1;`
	db.Exec(sqlStatement, id, name, lastName)

	return c.String(http.StatusOK, "USER WAS SUCCESSFULY UPDATED")

}
