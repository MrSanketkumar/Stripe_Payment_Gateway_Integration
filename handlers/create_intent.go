package handlers

import (
    "encoding/json"
    "net/http"
    "stripe-payment/models"
    "stripe-payment/services"
)

func CreateIntentHandler(paymentService services.PaymentService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req models.CreateIntentRequest
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        intent, err := paymentService.CreateIntent(req)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(intent)
    }
}
