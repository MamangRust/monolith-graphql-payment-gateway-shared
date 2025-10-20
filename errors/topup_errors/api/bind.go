package topupapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiBindCreateTopup returns a 400 error when binding the create topup request fails.
var ErrApiBindCreateTopup = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid create topup request", http.StatusBadRequest)
}

// ErrApiBindUpdateTopup returns a 400 error when binding the update topup request fails.
var ErrApiBindUpdateTopup = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid update topup request", http.StatusBadRequest)
}
