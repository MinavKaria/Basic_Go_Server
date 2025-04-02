package models

import (
	"go-server/database"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  int    `json:"user_id"`
}

func CreateUser(name, email string) (User, error) {
	user := User{Name: name, Email: email}

	err := database.DB.QueryRow(
		"INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id",
		user.Name, user.Email,
	).Scan(&user.ID)

	return user, err
}

func GetUser(id int) (User, error) {
	user := User{}

	err := database.DB.QueryRow(
		"SELECT id, name, email FROM users WHERE id = $1",
		id,
	).Scan(&user.ID, &user.Name, &user.Email)

	return user, err
}
