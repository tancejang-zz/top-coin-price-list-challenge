package accessor

import (
	"os"

	"gopkg.in/resty.v0"
)

// GetPricingInfo is a method to get info of pricing from coinapi.io
func GetPricingInfo() (*resty.Response, error) {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-CoinAPI-Key", os.Getenv("COINAPI_API_KEY")).
		Get("https://rest.coinapi.io/v1/trades/latest")

	return resp, err
}
