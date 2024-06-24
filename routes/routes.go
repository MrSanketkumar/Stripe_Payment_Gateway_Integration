package router

import (
    "net/http"
    "stripe-payment/handlers"
    "stripe-payment/repositories"
    "stripe-payment/services"

    "github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
    r := mux.NewRouter()

    stripeRepo := repositories.NewStripeRepository()
    paymentService := services.NewPaymentService(stripeRepo)

    r.HandleFunc("/api/v1/create_intent", handlers.CreateIntentHandler(paymentService)).Methods(http.MethodPost)
    r.HandleFunc("/api/v1/capture_intent/{id}", handlers.CaptureIntentHandler(paymentService)).Methods(http.MethodPost)
    r.HandleFunc("/api/v1/create_refund/{id}", handlers.CreateRefundHandler(paymentService)).Methods(http.MethodPost)
    r.HandleFunc("/api/v1/get_intents", handlers.GetIntentsHandler(paymentService)).Methods(http.MethodGet)

    return r
}
