package config

import (
    "os"

    "github.com/joho/godotenv"
)

var (
    StripeSecretKey string
)

func LoadConfig() {
    if err := godotenv.Load(); err != nil {
        panic("Error loading .env file")
    }

    StripeSecretKey = os.Getenv("STRIPE_SECRET_KEY")
}
