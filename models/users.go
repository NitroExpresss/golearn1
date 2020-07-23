package models

//User struct for internal processes
type User struct {
	ID        int
	Age       int    `json:"age"`
	FirstName string `json:"name"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

// GetUser Returns the user selected by ID
func GetUser(id int) (*User, error) {
	user := new(User)
	err := db.Model(user).Where("id = ?", id).Select()
	return user, err
}

// AddUser adds new user to db
func (user *User) AddUser() error {
	_, err := db.Model(user).Insert()
	return err
}

// GetUsers returns all users nested in db
func GetUsers() ([]User, error) {
	var users []User
	err := db.Model(&users).Order("id ASC").Select()
	return users, err
}

// DeleteUser deletes user from db
func DeleteUser(id int) error {
	user := User{ID: id}
	_, err := db.Model(&user).WherePK().Delete()
	return err
}

//UpdateUser updates the user in db
func (user *User) UpdateUser(id int) error {

	_, err := db.Model(user).Where("id = ?", id).Returning("*").UpdateNotZero()
	return err
}
