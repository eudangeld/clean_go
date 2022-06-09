package user

/*
CREATE TABLE userr (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email text NOT NULL,
	role integer NOT NULL,
	pass text NOT NULL
);
*/

type User struct {
	Id int `json:"id"`
	Email string `json:"email"`
	Role int `json:"role"`
	Pass string `json:"-"` //ignore this field on encode :)
}

type Role int

const (
	Basic = iota +1
	General
	Admin  
)