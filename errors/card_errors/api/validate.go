package cardapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiValidateCreateCard represents validation error for create card request.
var ErrApiValidateCreateCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid create Card request", http.StatusBadRequest)
}

// ErrApiValidateUpdateCard represents validation error for update card request.
var ErrApiValidateUpdateCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid update Card request", http.StatusBadRequest)
}
