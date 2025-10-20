package cardapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiInvalidCardID returns a 400 Bad Request error when the card ID is invalid.
var ErrApiInvalidCardID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid card ID", http.StatusBadRequest)
}

// ErrApiInvalidUserID returns a 400 Bad Request error when the user ID is invalid.
var ErrApiInvalidUserID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid user ID", http.StatusBadRequest)
}

// ErrApiInvalidCardNumber returns a 400 Bad Request error when the card number is invalid.
var ErrApiInvalidCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid card number", http.StatusBadRequest)
}

// ErrApiInvalidMonth returns a 400 Bad Request error when the month value is invalid.
var ErrApiInvalidMonth = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid month value", http.StatusBadRequest)
}

// ErrApiInvalidYear returns a 400 Bad Request error when the year value is invalid.
var ErrApiInvalidYear = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid year value", http.StatusBadRequest)
}
