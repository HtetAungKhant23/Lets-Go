package main

import (
	"fmt"

	"github.com/htetaungkhant/go/crypto/api"
)

func main() {
	res, err := api.GetRate("btc")
	if err == nil {
		fmt.Printf("This rate for %v is %.2f \n", res.Currency, res.Price)
	}
}
