# Stripe Payment Gateway Integration

Welcome to the Stripe Payment Gateway Integration project! This project involves designing and implementing backend APIs for integrating Stripe Payment Gateway using Go. The goal is to create a robust and scalable payment processing system.

## Project Features

1. **Create Payment Intent**: Initiates a payment process.
   - **Endpoint**: `POST /api/v1/create_intent`
2. **Capture Payment Intent**: Completes the payment process.
   - **Endpoint**: `POST /api/v1/capture_intent/:id`
3. **Create Refund**: Processes a refund for a payment.
   - **Endpoint**: `POST /api/v1/create_refund/:id`
4. **Retrieve List of Payment Intents**: Fetches all payment intents.
   - **Endpoint**: `GET /api/v1/get_intents`

## Requirements

- Go 1.20
- Gin or Gorilla-MUX
- Docker (optional for containerization)
- Stripe sandbox account with API keys

## Setup Instructions

### Clone the Repository

```sh
git clone https://github.com/yourusername/stripe-payment-gateway.git
cd stripe-payment-gateway
```

### Set Up Environment Variables

Create a `.env` file in the root directory and add your Stripe API keys:

```env
STRIPE_SECRET_KEY=your_secret_key
STRIPE_PUBLIC_KEY=your_public_key
```

### Install Dependencies

```sh
go mod tidy
```

### Run the Application

```sh
go run main.go
```

## API Endpoints

### Create Payment Intent

- **URL**: `/api/v1/create_intent`
- **Method**: `POST`
- **Description**: Creates a new payment intent.

### Capture Payment Intent

- **URL**: `/api/v1/capture_intent/:id`
- **Method**: `POST`
- **Description**: Captures a previously created payment intent.

### Create Refund

- **URL**: `/api/v1/create_refund/:id`
- **Method**: `POST`
- **Description**: Creates a refund for a captured payment intent.

### Retrieve List of Payment Intents

- **URL**: `/api/v1/get_intents`
- **Method**: `GET`
- **Description**: Retrieves a list of all payment intents.

## Testing

### Postman Collection

A Postman collection is provided to test the APIs. Import the collection from the following link: [Postman Collection Link](#)

### API URL for Testing

You can test the APIs using the following base URL in Postman:

```
http://localhost:8080
```

For example:
- To create a payment intent, use `POST http://localhost:8080/api/v1/create_intent`
- To capture a payment intent, use `POST http://localhost:8080/api/v1/capture_intent/:id`
- To create a refund, use `POST http://localhost:8080/api/v1/create_refund/:id`
- To get a list of payment intents, use `GET http://localhost:8080/api/v1/get_intents`

## Containerization

### Docker Setup

Build and run the Docker container:

```sh
docker build -t stripe-payment-gateway .
docker run -p 8080:8080 stripe-payment-gateway
```
