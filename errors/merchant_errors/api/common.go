package merchantapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiInvalidCardID is returned when the card ID provided is invalid.
var ErrApiInvalidCardID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid card ID", http.StatusBadRequest)
}

// ErrApiInvalidMerchantID is returned when the merchant ID provided is invalid.
var ErrApiInvalidMerchantID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid merchant ID", http.StatusBadRequest)
}

// ErrApiInvalidMonth is returned when the month parameter is invalid or malformed.
var ErrApiInvalidMonth = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid month value", http.StatusBadRequest)
}

// ErrApiInvalidYear is returned when the year parameter is invalid or malformed.
var ErrApiInvalidYear = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid year value", http.StatusBadRequest)
}

// ErrApiInvalidUserID is returned when the user ID provided is invalid.
var ErrApiInvalidUserID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid user ID", http.StatusBadRequest)
}

// ErrApiInvalidApiKey is returned when the API key provided is invalid or unauthorized.
var ErrApiInvalidApiKey = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid API key", http.StatusUnauthorized)
}
