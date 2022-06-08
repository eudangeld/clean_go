package coin_test

import (
	"database/sql"
	"testing"

	"github.com/eudangeld/clean_go/core/coin"
)


func TestStore(t *testing.T) {
	c:= &coin.Coin{
		Id:1,
		Name:"BTN",
		Value:156,
		Risk:coin.RiskHight,
	}
	

	db,err := sql.Open("sqlite3", "../../data/coin_test.db")
	defer db.Close()


	if(err != nil){
		t.Fatal("Db connect failed: ", err.Error())
	}


	err = clearDb(db)

	if err != nil {
		t.Fatal("Db clear failed: ", err.Error())
	}


	service := coin.NewService(db)
	err = service.Store(c)

	if(err != nil){
		t.Fatal("Error saving data into database: ", err.Error())
	}

	saved,err :=service.Get(1);

	if err != nil{
		t.Fatal("Error getting data from database: ", err.Error())
	}

	if saved.Id != 1{
		t.Fatal("Invalid data: ", err.Error())
	}
}


func clearDb (db *sql.DB) error{
	tx, err := db.Begin()

	if err != nil{
		return err
	}
	_, err = db.Exec("delete from coin")

	tx.Commit()
	return err
}