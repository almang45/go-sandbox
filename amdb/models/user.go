package models

type User struct {
	Id       int    `db:"id" json:"id"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
	Name     string `db:"name" json:"name"`
}
