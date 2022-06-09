package user

/*
CREATE TABLE userr (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email text NOT NULL,
	role integer NOT NULL
);
*/

type User struct {
	Id int `json:"id"`
	Email string `json:"email"`
	Role int `json:"role"`
}

type Role int

const (
	Basic = iota +1
	General
	Admin  
)