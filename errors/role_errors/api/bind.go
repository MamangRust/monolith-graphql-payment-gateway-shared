package roleapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiBindCreateRole returns an API error response for failed Role creation request binding.
var ErrApiBindCreateRole = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid create Role request", http.StatusBadRequest)
}

// ErrApiBindUpdateRole returns an API error response for failed Role update request binding.
var ErrApiBindUpdateRole = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid update Role request", http.StatusBadRequest)
}
