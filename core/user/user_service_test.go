package user_test

import (
	"database/sql"
	"testing"

	"github.com/eudangeld/clean_go/core/user"

	_ "github.com/mattn/go-sqlite3"
)

func TestSignup(t *testing.T) {
	u:= &user.User{
		Email: "asd@asd.com.br",
		Role: 1,
		Pass: "asdasd",
	}
	
	
	db,err := sql.Open("sqlite3", "../../data/coin_test.db")
	defer db.Close()
	
	
	if err != nil {
		t.Fatalf("Error creating table: %s",err)
	}

	err = clearDb(db)

	if err != nil {
		t.Fatalf("Db clear failed: %s", err.Error())
	}


	if(err != nil){
		t.Fatalf("Db connect failed: %s", err.Error())
	}

		
	service := user.NewUserService(db)
	err = service.Signup(u.Email, u.Pass, u.Role)
	
	if(err != nil){
		t.Fatalf("Error saving user into database: %s", err.Error())
	}

}



func TestLogin(t *testing.T) {
	u:= &user.User{
		Id: 1,
		Email: "asd@asd.com.br",
		Role: 1,
		Pass: "asdasd",
	}
	db,err := sql.Open("sqlite3", "../../data/coin_test.db")
	defer db.Close()
	
	
	if err != nil {
		t.Fatalf("Error creating table: %s",err)
	}

	err = clearDb(db)

	service := user.NewUserService(db)
	err = service.Signup(u.Email, u.Pass, u.Role)

	loggedUser,err := service.Login(u.Email, u.Pass)

	
	if err != nil{
		t.Fatalf("Loggin error: %s", err.Error())
	}



	if loggedUser.Email != u.Email{
		t.Fatalf("Email not match: %s", loggedUser.Email)
	}

}

func clearDb (db *sql.DB) error{
	tx, err := db.Begin()

	if err != nil{
		return err
	}
	_, err = db.Exec("delete from userr")

	tx.Commit()
	return err
}
