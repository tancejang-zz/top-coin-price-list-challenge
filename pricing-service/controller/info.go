package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tancejang/go_challenge_pricing_service/pricing-service/accessor"
)

// GetPriceInfo is a controller that handle all request related to /info path
func GetPriceInfo(w http.ResponseWriter, r *http.Request) {
	result, err := accessor.GetPricingInfo()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, result.String())
}
