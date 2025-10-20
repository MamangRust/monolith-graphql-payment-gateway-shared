package transactionapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiValidateCreateTransaction returns an API error response
// when validation of the create transaction request fails.
var ErrApiValidateCreateTransaction = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid create transaction request", http.StatusBadRequest)
}

// ErrApiValidateUpdateTransaction returns an API error response
// when validation of the update transaction request fails.
var ErrApiValidateUpdateTransaction = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid update transaction request", http.StatusBadRequest)
}
