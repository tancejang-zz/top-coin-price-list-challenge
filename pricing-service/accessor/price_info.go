package accessor

import (
	"gopkg.in/resty.v0"
)

// GetPricingInfo is a method to get info of pricing from coinapi.io
func GetPricingInfo() (*resty.Response, error) {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-CoinAPI-Key", "A1269882-C37A-49A8-9413-270B3E2FDC5B").
		Get("https://rest.coinapi.io/v1/trades/latest")

	return resp, err
}
