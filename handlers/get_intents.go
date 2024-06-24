package handlers

import (
    "encoding/json"
    "net/http"
    "stripe-payment/services"
)

func GetIntentsHandler(paymentService services.PaymentService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        intents, err := paymentService.GetIntents()
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(intents)
    }
}
