package userapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiBindCreateUser is an error response for bind failure: invalid create user request.
var ErrApiBindCreateUser = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid create User request", http.StatusBadRequest)
}

// ErrApiBindUpdateUser is an error response for bind failure: invalid update user request.
var ErrApiBindUpdateUser = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid update User request", http.StatusBadRequest)
}
