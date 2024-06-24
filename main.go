package main

import (
	"log"
	"net/http"
	"stripe-payment/config"
	router "stripe-payment/routes"

	"stripe-payment/utils"
)

func main() {
    config.LoadConfig()
    utils.SetupLogger()

    router := router.SetupRouter()
    log.Fatal(http.ListenAndServe(":8080",router))
}
