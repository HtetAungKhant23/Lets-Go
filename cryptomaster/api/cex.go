package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	cryptorate "github.com/htetaungkhant/go/crypto/cryptoRate"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*cryptorate.Rate, error) { // pointer can be nil
	upCurrency := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiUrl, upCurrency))
	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		var cryptoRate cryptorate.Rate
		err = json.Unmarshal(bodyBytes, &cryptoRate) // json parse
		if err != nil {
			return nil, err
		}
	}
	crypto := cryptorate.Rate{Currency: currency, Price: 20}
	return &crypto, nil
}
