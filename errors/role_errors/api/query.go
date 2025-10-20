package roleapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiRoleNotFound returns an API error response when the requested Role is not found.
var ErrApiRoleNotFound = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Role not found", http.StatusNotFound)
}

// ErrApiFailedFindAll returns an API error response when fetching all Roles fails.
var ErrApiFailedFindAll = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to fetch Roles", http.StatusInternalServerError)
}

// ErrApiFailedFindActive returns an API error response when fetching active Roles fails.
var ErrApiFailedFindActive = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to fetch active Roles", http.StatusInternalServerError)
}

// ErrApiFailedFindTrashed returns an API error response when fetching trashed Roles fails.
var ErrApiFailedFindTrashed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to fetch trashed Roles", http.StatusInternalServerError)
}
