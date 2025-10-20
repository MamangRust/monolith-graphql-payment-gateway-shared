package transactionapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiBindCreateTransaction returns an API error response
// when binding the create transaction request body to struct fails.
var ErrApiBindCreateTransaction = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid create transaction request", http.StatusBadRequest)
}

// ErrApiBindUpdateTransaction returns an API error response
// when binding the update transaction request body to struct fails.
var ErrApiBindUpdateTransaction = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid update transaction request", http.StatusBadRequest)
}
