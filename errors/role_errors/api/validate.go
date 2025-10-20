package roleapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiValidateCreateRole returns an API error response for invalid Role creation request validation.
var ErrApiValidateCreateRole = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid create Role request", http.StatusBadRequest)
}

// ErrApiValidateUpdateRole returns an API error response for invalid Role update request validation.
var ErrApiValidateUpdateRole = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid update Role request", http.StatusBadRequest)
}
