package handlers

import (
	"encoding/json"
	"net/http"
	"stripe-payment/services"

	"github.com/gorilla/mux"
)

func CreateRefundHandler(paymentService services.PaymentService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        id := vars["id"]

        refund, err := paymentService.CreateRefund(id)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(refund)
    }
}
