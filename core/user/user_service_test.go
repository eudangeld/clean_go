package user_test

import (
	"database/sql"
	"testing"

	"github.com/eudangeld/clean_go/core/user"

	_ "github.com/mattn/go-sqlite3"
)

func TestStore(t *testing.T) {
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
	
	saved,err :=service.GetUser(u.Id);

	

	
	if err != nil{
		t.Fatalf("Error getting data from database: %s", err.Error())
	}
	
	if saved.Id != 1{
		t.Fatalf("Invalid data: %s", err.Error())
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
