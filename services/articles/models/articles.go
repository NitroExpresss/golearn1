package models

import "golearn1/services/structs"

// Art struct for internal usage
type Art struct {
	tableName struct{} `pg:"articles"`
	structs.Article
}

// GetArticle returns the Article selected by ID
func GetArticle(id int) (*Art, error) {
	art := new(Art)
	err := db.Model(art).Where("id = ?", id).Select()
	return art, err
}

// AddArticle adds new Article to db
func (art *Art) AddArticle() error {
	_, err := db.Model(art).Insert()
	return err
}

// GetArticles returns all articles nested in db
func GetArticles() ([]Art, error) {
	var arts []Art
	err := db.Model(&arts).Order("id ASC").Select()
	return arts, err
}

// GetUserArticles returns all articles nested in db
func GetUserArticles(id int) ([]Art, error) {
	var arts []Art
	err := db.Model(&arts).Where("userid = ?", id).Order("id ASC").Select()
	return arts, err
}

// DeleteArticle deletes Article from db
func DeleteArticle(id int) error {
	art := new(Art)
	art.ID = id
	_, err := db.Model(&art).WherePK().Delete()
	return err
}

//UpdateArticle updates the Article in db
func (art *Art) UpdateArticle(id int) error {

	_, err := db.Model(art).Where("id = ?", id).Returning("*").UpdateNotZero()
	return err
}
