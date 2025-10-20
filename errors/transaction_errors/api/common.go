package transactionapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiInvalidTransactionApiKey returns an API error response
// when the transaction API key is missing or invalid.
var ErrApiInvalidTransactionApiKey = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid transaction api-key", http.StatusBadRequest)
}

// ErrApiInvalidTransactionID returns an API error response
// when the transaction ID is invalid or improperly formatted.
var ErrApiInvalidTransactionID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid transaction ID", http.StatusBadRequest)
}

// ErrApiInvalidTransactionCardNumber returns an API error response
// when the provided transaction card number is invalid.
var ErrApiInvalidTransactionCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid transaction Card Number", http.StatusBadRequest)
}

// ErrApiInvalidTransactionMerchantID returns an API error response
// when the provided merchant ID for the transaction is invalid.
var ErrApiInvalidTransactionMerchantID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid transaction merchant ID", http.StatusBadRequest)
}

// ErrApiInvalidMonth returns an API error response
// when the month value is missing, invalid, or out of range.
var ErrApiInvalidMonth = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid month value", http.StatusBadRequest)
}

// ErrApiInvalidYear returns an API error response
// when the year value is missing, invalid, or out of range.
var ErrApiInvalidYear = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid year value", http.StatusBadRequest)
}
