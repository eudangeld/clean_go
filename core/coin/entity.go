package coin


type Risk int

type Coin struct {
	Id  int  `json:"id"`
	Name  string  `json:"name"`
	Value  int32  `json:"value"`
	Risk  Risk  `json:"risk"`
}

/*
CREATE TABLE coin (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name text NOT NULL,
	value integer NOT NULL,
	risk integer NOT NULL
);
*/



const (
	RiskMinor = iota +1
	RiskModerate
	RiskHight
)


func (risk Risk) String() string {
	switch risk {
	case RiskMinor:
		return "Baixo risco"
	case RiskModerate:
		return "Risco moderado"
	}
	return "Risco não calculado"

	}
