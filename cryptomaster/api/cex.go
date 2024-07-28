package api

import (
	"fmt"
	"net/http"
	"strings"

	cryptorate "github.com/htetaungkhant/go/crypto/cryptoRate"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (cryptorate.Rate, error) {
	upCurrency := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiUrl, upCurrency))
}
