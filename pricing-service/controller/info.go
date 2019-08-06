package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tancejang/top-coin-price-list-challenge/pricing-service/accessor"
)

// GetPriceInfo is a controller that handle all request related to /info path
func GetPriceInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		result, err := accessor.GetPricingInfo()
		defer fmt.Fprintf(w, result.String())
		if err != nil {
			log.Fatal(err)
		}
	} else {
		http.Error(w, "404 page not found", http.StatusNotFound)
	}
}
