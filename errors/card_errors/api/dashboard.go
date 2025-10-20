package cardapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiFailedDashboardCard returns a 500 Internal Server Error when fetching the card dashboard fails.
var ErrApiFailedDashboardCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch card dashboard", http.StatusInternalServerError)
}

// ErrApiFailedDashboardCardByCardNumber returns a 500 Internal Server Error when fetching the dashboard by card number fails.
var ErrApiFailedDashboardCardByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch dashboard by card number", http.StatusInternalServerError)
}
