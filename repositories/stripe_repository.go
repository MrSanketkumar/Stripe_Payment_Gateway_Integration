package repositories

import (
    "stripe-payment/config"
    "stripe-payment/models"

    "github.com/stripe/stripe-go/v72"
    "github.com/stripe/stripe-go/v72/paymentintent"
    "github.com/stripe/stripe-go/v72/refund"
)

type StripeRepository interface {
    CreateIntent(req models.CreateIntentRequest) (*models.PaymentIntent, error)
    CaptureIntent(id string) (*models.PaymentIntent, error)
    CreateRefund(id string) (*models.Refund, error)
    GetIntents() ([]models.PaymentIntent, error)
}

type stripeRepository struct{}

func NewStripeRepository() StripeRepository {
    stripe.Key = config.StripeSecretKey
    return &stripeRepository{}
}

func (s *stripeRepository) CreateIntent(req models.CreateIntentRequest) (*models.PaymentIntent, error) {
    params := &stripe.PaymentIntentParams{
        Amount:   stripe.Int64(req.Amount),
        Currency: stripe.String(req.Currency),
    }

    intent, err := paymentintent.New(params)
    if err != nil {
        return nil, err
    }

    return &models.PaymentIntent{ID: intent.ID, Amount: intent.Amount, Currency: intent.Currency}, nil
}

func (s *stripeRepository) CaptureIntent(id string) (*models.PaymentIntent, error) {
    intent, err := paymentintent.Capture(id, nil)
    if err != nil {
        return nil, err
    }

    return &models.PaymentIntent{ID: intent.ID, Amount: intent.Amount, Currency: intent.Currency}, nil
}

func (s *stripeRepository) CreateRefund(id string) (*models.Refund, error) {
    params := &stripe.RefundParams{
        PaymentIntent: stripe.String(id),
    }

    refund, err := refund.New(params)
    if err != nil {
        return nil, err
    }

    return &models.Refund{ID: refund.ID, Amount: refund.Amount}, nil
}

func (s *stripeRepository) GetIntents() ([]models.PaymentIntent, error) {
    params := &stripe.PaymentIntentListParams{}
    params.Filters.AddFilter("limit", "", "10")

    i := paymentintent.List(params)
    var intents []models.PaymentIntent
    for i.Next() {
        pi := i.PaymentIntent()
        intents = append(intents, models.PaymentIntent{ID: pi.ID, Amount: pi.Amount, Currency: pi.Currency})
    }
    if err := i.Err(); err != nil {
        return nil, err
    }

    return intents, nil
}
