package saldoapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiInvalidSaldoID indicates an invalid saldo ID in the request.
var ErrApiInvalidSaldoID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid saldo ID", http.StatusBadRequest)
}

// ErrApiInvalidMonth indicates an invalid month value in the request.
var ErrApiInvalidMonth = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid month value", http.StatusBadRequest)
}

// ErrApiInvalidYear indicates an invalid year value in the request.
var ErrApiInvalidYear = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid year value", http.StatusBadRequest)
}

// ErrApiInvalidCardNumber indicates an invalid card number value in the request.
var ErrApiInvalidCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid card-number value", http.StatusBadRequest)
}
