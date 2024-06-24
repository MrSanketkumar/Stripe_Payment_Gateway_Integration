package handlers

import (
    "encoding/json"
    "net/http"
    "stripe-payment/utils"
    "stripe-payment/services"

    "github.com/gorilla/mux"
)

func CaptureIntentHandler(paymentService services.PaymentService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        id := vars["id"]

        intent, err := paymentService.CaptureIntent(id)
        if err != nil {
            utils.ErrorLogger.Printf("Failed to capture payment intent: %v", err)
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(intent)
    }
}
