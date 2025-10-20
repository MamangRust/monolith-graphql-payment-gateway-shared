package topupapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiInvalidTopupID returns a 400 Bad Request error when the provided top-up ID is invalid.
var ErrApiInvalidTopupID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid topup ID", http.StatusBadRequest)
}

// ErrApiInvalidCardNumber returns a 400 Bad Request error when the provided card number is invalid.
var ErrApiInvalidCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid card number", http.StatusBadRequest)
}

// ErrApiInvalidMonth returns a 400 Bad Request error when the provided month value is invalid or missing.
var ErrApiInvalidMonth = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid month value", http.StatusBadRequest)
}

// ErrApiInvalidYear returns a 400 Bad Request error when the provided year value is invalid or missing.
var ErrApiInvalidYear = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid year value", http.StatusBadRequest)
}
