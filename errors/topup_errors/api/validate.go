package topupapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiValidateCreateTopup returns a 400 error when the create topup request fails validation.
var ErrApiValidateCreateTopup = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid create topup request", http.StatusBadRequest)
}

// ErrApiValidateUpdateTopup returns a 400 error when the update topup request fails validation.
var ErrApiValidateUpdateTopup = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid update topup request", http.StatusBadRequest)
}
