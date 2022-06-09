package user

import (
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UseCase interface {
	GetUser(ID int)(*User, error)	
	Login(Email string, Pass string)(*User, error)	
	Signup(Email string, Pass string, Role int)(error)	
}


type UserService struct {
	DB *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		DB:db,
	}
}

func (s *UserService)GetUser(ID int)(*User, error) {
	var u User
	stmt, err := s.DB.Prepare("SELECT id, email, role from userr where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(ID).Scan(&u.Id, &u.Email, &u.Role)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (s *UserService)Signup(Email string, Pass string, Role int)(error) {
	var u User
	bytePass := []byte(Pass)
	saltedPass := hashPassword(bytePass)

	tx,err := s.DB.Begin()

	stmt, err := s.DB.Prepare("INSERT INTO userr (email, pass, role) values (?, ?, ?)")
	if err != nil {
		return err
	}
	

	defer stmt.Close()	

	_,err = stmt.Exec(u.Email, saltedPass, Role)
	if err != nil {
		return err
	}
	tx.Commit()

	return nil
}

func (s *UserService)Login(Email string, Pass string)(*User, error)	{
	var u User
	stmt, err := s.DB.Prepare("SELECT id, email, role, pass from userr where email = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(Email).Scan(&u.Id, &u.Email, &u.Role, &u.Pass)
	if err != nil {
		return nil, err
	}
	if comparePasswords(u.Pass, []byte(Pass)) {
		return &u, nil
	}
	return nil, nil
}



func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func hashPassword(pwd []byte ) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}