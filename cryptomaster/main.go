package main

import (
	"fmt"
	"sync"

	"github.com/htetaungkhant/go/crypto/api"
)

func main() {
	currencies := []string{"btc", "eth", "sol", "bch"}
	var wg sync.WaitGroup // using wait group for async await feature
	for _, currency := range currencies {
		wg.Add(1)
		go func(currency string) { // closure function
			getCurrencyRate(currency)
			wg.Done()
		}(currency)
	}
	wg.Wait()
}

func getCurrencyRate(currency string) {
	res, err := api.GetRate(currency)
	if err == nil {
		fmt.Printf("This rate for %v is %.2f \n", res.Currency, res.Price)
	}
}
