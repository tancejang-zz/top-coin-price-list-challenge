package main

import (
	"log"
	"net/http"

	"github.com/tancejang/go_challenge_pricing_service/pricing-service/controller"
)

func main() {

	http.HandleFunc("/info", controller.GetPriceInfo)

	log.Println("Starting Server")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
