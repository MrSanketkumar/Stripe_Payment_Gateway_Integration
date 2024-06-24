package services

import (
    "stripe-payment/models"
    "stripe-payment/repositories"
)

type PaymentService interface {
    CreateIntent(req models.CreateIntentRequest) (*models.PaymentIntent, error)
    CaptureIntent(id string) (*models.PaymentIntent, error)
    CreateRefund(id string) (*models.Refund, error)
    GetIntents() ([]models.PaymentIntent, error)
}

type paymentService struct {
    stripeRepo repositories.StripeRepository
}

func NewPaymentService(repo repositories.StripeRepository) PaymentService {
    return &paymentService{stripeRepo: repo}
}

func (s *paymentService) CreateIntent(req models.CreateIntentRequest) (*models.PaymentIntent, error) {
    return s.stripeRepo.CreateIntent(req)
}

func (s *paymentService) CaptureIntent(id string) (*models.PaymentIntent, error) {
    return s.stripeRepo.CaptureIntent(id)
}

func (s *paymentService) CreateRefund(id string) (*models.Refund, error) {
    return s.stripeRepo.CreateRefund(id)
}

func (s *paymentService) GetIntents() ([]models.PaymentIntent, error) {
    return s.stripeRepo.GetIntents()
}
