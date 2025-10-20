package roleapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrInvalidRoleId returns an API error response for an invalid Role ID format.
var ErrInvalidRoleId = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "invalid Role id", http.StatusBadRequest)
}

// ErrApiRoleInvalidId returns an API error response for an invalid Role ID.
var ErrApiRoleInvalidId = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "invalid Role id", http.StatusNotFound)
}
