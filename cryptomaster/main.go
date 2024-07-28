package main

import (
	"fmt"

	"github.com/htetaungkhant/go/crypto/api"
)

func main() {
	res, err := api.GetRate("btc")
	fmt.Println("result ", *res)
	fmt.Println("err ", err)
}
