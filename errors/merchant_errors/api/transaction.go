package merchantapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiFailedFindAllTransactions is returned when failing to retrieve all merchant transactions.
var ErrApiFailedFindAllTransactions = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all transactions", http.StatusInternalServerError)
}

// ErrApiFailedFindAllTransactionByMerchant is returned when failing to retrieve transactions by merchant.
var ErrApiFailedFindAllTransactionByMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch transactions by merchant", http.StatusInternalServerError)
}

// ErrApiFailedFindAllTransactionByApikey is returned when failing to retrieve transactions by API key.
var ErrApiFailedFindAllTransactionByApikey = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch transactions by API key", http.StatusInternalServerError)
}
