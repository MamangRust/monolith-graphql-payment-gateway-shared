package userapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiFailedCreateUser is an error response for failing to create a user.
var ErrApiFailedCreateUser = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to create User", http.StatusInternalServerError)
}

// ErrApiFailedUpdateUser is an error response for failing to update a user.
var ErrApiFailedUpdateUser = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to update User", http.StatusInternalServerError)
}

// ErrApiFailedTrashedUser is an error response for failed to move user to trash.
var ErrApiFailedTrashedUser = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to move User to trash", http.StatusInternalServerError)
}

// ErrApiFailedRestoreUser is an error response for failed to restore user.
var ErrApiFailedRestoreUser = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to restore User", http.StatusInternalServerError)
}

// ErrApiFailedDeletePermanent is an error response for failed to delete user permanently.
var ErrApiFailedDeletePermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to delete User permanently", http.StatusInternalServerError)
}

// ErrApiFailedRestoreAll is an error response for failed to restore all users.
var ErrApiFailedRestoreAll = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to restore all Users", http.StatusInternalServerError)
}

// ErrApiFailedDeleteAll is an error response for failed to permanently delete all users.
var ErrApiFailedDeleteAll = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to delete all Users permanently", http.StatusInternalServerError)
}
