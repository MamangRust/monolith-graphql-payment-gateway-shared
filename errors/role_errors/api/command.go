package roleapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiFailedCreateRole returns an API error response when Role creation fails.
var ErrApiFailedCreateRole = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to create Role", http.StatusInternalServerError)
}

// ErrApiFailedUpdateRole returns an API error response when Role update fails.
var ErrApiFailedUpdateRole = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to update Role", http.StatusInternalServerError)
}

// ErrApiFailedTrashedRole returns an API error response when soft-deleting (trashing) a Role fails.
var ErrApiFailedTrashedRole = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to move Role to trash", http.StatusInternalServerError)
}

// ErrApiFailedRestoreRole returns an API error response when restoring a trashed Role fails.
var ErrApiFailedRestoreRole = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to restore Role", http.StatusInternalServerError)
}

// ErrApiFailedDeletePermanent returns an API error response when permanently deleting a Role fails.
var ErrApiFailedDeletePermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to delete Role permanently", http.StatusInternalServerError)
}

// ErrApiFailedRestoreAll returns an API error response when restoring all trashed Roles fails.
var ErrApiFailedRestoreAll = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to restore all Roles", http.StatusInternalServerError)
}

// ErrApiFailedDeleteAll returns an API error response when permanently deleting all Roles fails.
var ErrApiFailedDeleteAll = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to delete all Roles permanently", http.StatusInternalServerError)
}
