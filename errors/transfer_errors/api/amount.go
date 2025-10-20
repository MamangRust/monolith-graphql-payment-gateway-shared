package transactionapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiFailedFindMonthlyTransferAmounts returns an API error response
// when fetching monthly total transfer amounts fails.
var ErrApiFailedFindMonthlyTransferAmounts = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly transfer amounts", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTransferAmounts returns an API error response
// when fetching yearly total transfer amounts fails.
var ErrApiFailedFindYearlyTransferAmounts = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly transfer amounts", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTransferAmountsBySenderCardNumber returns an API error response
// when fetching monthly transfer amounts by sender card number fails.
var ErrApiFailedFindMonthlyTransferAmountsBySenderCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly transfer amounts by sender card number", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTransferAmountsByReceiverCardNumber returns an API error response
// when fetching monthly transfer amounts by receiver card number fails.
var ErrApiFailedFindMonthlyTransferAmountsByReceiverCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly transfer amounts by receiver card number", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTransferAmountsBySenderCardNumber returns an API error response
// when fetching yearly transfer amounts by sender card number fails.
var ErrApiFailedFindYearlyTransferAmountsBySenderCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly transfer amounts by sender card number", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTransferAmountsByReceiverCardNumber returns an API error response
// when fetching yearly transfer amounts by receiver card number fails.
var ErrApiFailedFindYearlyTransferAmountsByReceiverCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly transfer amounts by receiver card number", http.StatusInternalServerError)
}
