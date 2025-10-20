package transactionapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiFailedFindAllTransactions returns an API error response
// when fetching all transactions from the database fails.
var ErrApiFailedFindAllTransactions = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all transactions", http.StatusInternalServerError)
}

// ErrApiFailedFindAllTransactionsActive returns an API error response
// when fetching active (non-deleted) transactions fails.
var ErrApiFailedFindAllTransactionsActive = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all transactions active", http.StatusInternalServerError)
}

// ErrApiFailedFindAllTransactionsTrashed returns an API error response
// when fetching trashed (soft-deleted) transactions fails.
var ErrApiFailedFindAllTransactionsTrashed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all transactions active", http.StatusInternalServerError)
}

// ErrApiFailedFindByCardNumber returns an API error response
// when fetching transactions filtered by card number fails.
var ErrApiFailedFindByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch transactions by card number", http.StatusInternalServerError)
}

// ErrApiFailedFindById returns an API error response
// when fetching a transaction by its ID fails.
var ErrApiFailedFindById = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch transaction by ID", http.StatusInternalServerError)
}

// ErrApiFailedFindByMerchantID returns an API error response
// when retrieving transactions by merchant ID fails.
var ErrApiFailedFindByMerchantID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch transactions by merchant ID", http.StatusInternalServerError)
}

// ErrApiFailedFindActive returns an API error response
// when retrieving active (non-deleted) transactions fails.
var ErrApiFailedFindActive = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch active transactions", http.StatusInternalServerError)
}

// ErrApiFailedFindTrashed returns an API error response
// when retrieving trashed (soft-deleted) transactions fails.
var ErrApiFailedFindTrashed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch trashed transactions", http.StatusInternalServerError)
}
