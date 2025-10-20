package transactionapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiValidateCreateTransfer returns an API error response
// when validation fails on the create transfer request payload.
var ErrApiValidateCreateTransfer = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid create transfer request", http.StatusBadRequest)
}

// ErrApiValidateUpdateTransfer returns an API error response
// when validation fails on the update transfer request payload.
var ErrApiValidateUpdateTransfer = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid update transfer request", http.StatusBadRequest)
}
