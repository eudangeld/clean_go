package coin_test

import (
	"testing"

	"github.com/eudangeld/clean_go/core/coin"
)


func TestStore(t *testing.T) {
	c:= &coin.Coin{
		Id:1,
		Name:"Eu",
		Value:1,
		Risk:1,
	}
}