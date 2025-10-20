package withdrawapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiBindCreateWithdraw is an error response for bind failed: invalid create withdraw request.
var ErrApiBindCreateWithdraw = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid create withdraw request", http.StatusBadRequest)
}

// ErrApiBindUpdateWithdraw is an error response for bind failed: invalid update withdraw request.
var ErrApiBindUpdateWithdraw = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid update withdraw request", http.StatusBadRequest)
}
