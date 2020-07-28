package models

import (
	"encoding/json"
	"fmt"
	"golearn1/services/structs"
	"net/http"
)

// Usr struct for internal use
type Usr struct {
	structs.User
}

// UsrArts struct for internal use
type UsrArts struct {
	structs.UserArticles
}

//GetUserArticles returns UserArticles struct
func (ua *UsrArts) GetUserArticles(id string) error {
	err := db.Model(&ua.User).Where("id = ?", id).Select()
	url := "http://localhost:1324/arts/usr/"
	resp, err := http.Get(url + id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&ua.Articles)

	return err
}

// GetUser Returns the user selected by ID
func GetUser(id int) (*Usr, error) {
	user := new(Usr)
	err := db.Model(user).Where("id = ?", id).Select()
	return user, err
}

// AddUser adds new user to db
func (user *Usr) AddUser() error {
	_, err := db.Model(user).Insert()
	return err
}

// GetUsers returns all users nested in db
func GetUsers() ([]Usr, error) {
	var users []Usr
	err := db.Model(&users).Order("id ASC").Select()
	return users, err
}

// DeleteUser deletes user from db
func DeleteUser(id int) error {
	user := new(Usr)
	user.ID = id
	_, err := db.Model(&user).WherePK().Delete()
	return err
}

//UpdateUser updates the user in db
func (user *Usr) UpdateUser(id int) error {

	_, err := db.Model(user).Where("id = ?", id).Returning("*").UpdateNotZero()
	return err
}
