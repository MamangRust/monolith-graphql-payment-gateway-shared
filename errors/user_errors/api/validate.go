package userapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiValidateCreateUser is an error response for validation failure in creating a user.
var ErrApiValidateCreateUser = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid create User request", http.StatusBadRequest)
}

// ErrApiValidateUpdateUser is an error response for validation failure in updating a user.
var ErrApiValidateUpdateUser = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid update User request", http.StatusBadRequest)
}
