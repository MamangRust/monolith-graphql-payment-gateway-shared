package userapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiUserInvalidId is an error response for an invalid user ID.
var ErrApiUserInvalidId = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "invalid User id", http.StatusNotFound)
}
