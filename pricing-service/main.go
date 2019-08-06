package main

import (
	"log"
	"net/http"

	"github.com/tancejang/top-coin-price-list-challenge/pricing-service/controller"
)

func main() {

	http.HandleFunc("/info", controller.GetPriceInfo)

	log.Println("Starting Server")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
