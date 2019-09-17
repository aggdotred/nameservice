package types

import (
	"fmt"
	"strings"
)

// Whois is a struct that contains all of the metadata of a name
type Whois struct {
	Value string         `json:"value"`
	Owner sdk.AccAddress `json:"owner"`
	Price sdk.Coins      `json:"price"`
}

// MinNamePrice is the initial starting price for a name that was never previously owned
var MinNamePrice = sdk.Coins{sdk.NewInt64Coin("nametoken", 1)}

// NewWhois returns a new Whois with the minprice as the price
func NewWhois() Whois {
	return Whois{
		Price: MinNamePrice,
	}
}

func (w Whois) String() string {
	return strings.TrimSpace(fmt.Sprint(
		`Owner: %s
		 Value: %s
		 Price: %s`, w.Owner, w.Value, w.Price))
}
