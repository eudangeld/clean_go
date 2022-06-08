package coin_test

import (
	"database/sql"
	"testing"

	"github.com/eudangeld/clean_go/core/coin"

	_ "github.com/mattn/go-sqlite3"
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


	
	
	service := coin.NewService(db)
	err = service.Store(c)
	
	if(err != nil){
		t.Fatalf("Error saving data into database: %s", err.Error())
	}
	
	saved,err :=service.Get(1);

	

	
	if err != nil{
		t.Fatalf("Error getting data from database: %s", err.Error())
	}
	
	if saved.Id != 1{
		t.Fatalf("Invalid data: %s", err.Error())
	}
	
}

func TestGetAll(t *testing.T) {


	mockedData := [2]*coin.Coin{
		{
			Id: 1,
			Name: "BTC",
			Value: 100,
			Risk: coin.RiskMinor,
		},
		{
			Id: 2,
			Name: "ETH",
			Value: 110,
			Risk: coin.RiskModerate,
		},
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

	
	err = clearDb(db)

	if err != nil {
		t.Fatalf("Db cleaning fail %s", err.Error())
	}


	service := coin.NewService(db)
	
	if(err != nil){
		t.Fatalf("Error saving data into database: %s", err.Error())
	}
	
	
	for _, c := range mockedData {
		err = service.Store(c)
		if err != nil {
			t.Fatalf("Error saving data into database: %s", err.Error())
		}
	}

	resultCoins,err := service.GetAll()
	if(resultCoins[0].Id != mockedData[0].Id){
		t.Fatal("Database getAll result is different from input data", )
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
