package withdrawapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiValidateCreateWithdraw is an error response for validation failed: invalid create withdraw request.
var ErrApiValidateCreateWithdraw = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid create withdraw request", http.StatusBadRequest)
}

// ErrApiValidateUpdateWithdraw is an error response for validation failed: invalid update withdraw request.
var ErrApiValidateUpdateWithdraw = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid update withdraw request", http.StatusBadRequest)
}
