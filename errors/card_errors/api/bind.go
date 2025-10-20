package cardapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiBindCreateCard represents binding error for create card request.
var ErrApiBindCreateCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid create Card request", http.StatusBadRequest)
}

// ErrApiBindUpdateCard represents binding error for update card request.
var ErrApiBindUpdateCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid update Card request", http.StatusBadRequest)
}
