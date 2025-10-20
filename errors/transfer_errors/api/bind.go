package transactionapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiBindCreateTransfer returns an API error response
// when binding the request for creating a transfer fails due to invalid input.
var ErrApiBindCreateTransfer = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid create transfer request", http.StatusBadRequest)
}

// ErrApiBindUpdateTransfer returns an API error response
// when binding the request for updating a transfer fails due to invalid input.
var ErrApiBindUpdateTransfer = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid update transfer request", http.StatusBadRequest)
}
