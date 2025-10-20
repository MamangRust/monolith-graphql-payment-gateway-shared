package userapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiUserNotFound is an error response for when the user is not found.
var ErrApiUserNotFound = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "User not found", http.StatusNotFound)
}

// ErrApiFailedFindAll is an error response for failing to fetch all users.
var ErrApiFailedFindAll = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to fetch Users", http.StatusInternalServerError)
}

// ErrApiFailedFindActive is an error response for failing to fetch active users.
var ErrApiFailedFindActive = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to fetch active Users", http.StatusInternalServerError)
}

// ErrApiFailedFindTrashed is an error response for failing to fetch trashed users.
var ErrApiFailedFindTrashed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to fetch trashed Users", http.StatusInternalServerError)
}
