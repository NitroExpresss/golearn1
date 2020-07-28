package structs

//User struct for internal processes
type User struct {
	ID       int    `json:"id"`
	Username string `json:"name"`
}

//Article struct for internal processes
type Article struct {
	ID     int
	Title  string `json:"title"`
	Body   string `json:"body"`
	Userid int    `json:"userid"`
}

//UserArticles it is
type UserArticles struct {
	User     User      `json:"user"`
	Articles []Article `json:"articles"`
}
