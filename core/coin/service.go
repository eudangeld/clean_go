package coin

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Usecase interface {
	GetAll()([]*Coin, error)
	Get(ID int64)(*Coin, error)
	Store(c *Coin) error
}

type Service struct {
	DB *sql.DB
}

func NewService (db *sql.DB)*Service {
	return &Service{
		DB:db,
	}
}


func (s *Service) GetAll()([]*Coin, error) {
	var result []*Coin

	rows,err := s.DB.Query("SELECT id, name, value, risk from coin")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var c Coin
		err = rows.Scan(&c.Id, &c.Name, &c.Value, &c.Risk)
		if err != nil {
			return nil, err
		}
		result = append(result, &c)
	}

	return result, nil
}

func (s *Service) Get(ID int)(*Coin, error){
	var c Coin

	stmt, err := s.DB.Prepare("SELECT id, name, value, risk from coin where id = ?")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(ID).Scan(&c.Id, &c.Name, &c.Value, &c.Risk)
	if err != nil {
		return nil, err
	}

	return &c,nil
}

func (s *Service) Store(c *Coin)(error){

	tx,err := s.DB.Begin()

	if err != nil {
		return err
	}

	stmt ,err := tx.Prepare("INSERT INTO coin (name, value, risk) VALUES (?,?,?)")

	if err != nil {
		return err
	}

	defer stmt.Close()	

	_,err = stmt.Exec(c.Name, c.Value, c.Risk)
	if err != nil {
		return err
	}
	tx.Commit()

	return nil
}


