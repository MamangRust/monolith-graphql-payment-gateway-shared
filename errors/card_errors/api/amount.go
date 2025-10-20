package cardapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiFailedFindMonthlyTopupAmount returns an API error response when fetching the monthly top-up amount fails.
var ErrApiFailedFindMonthlyTopupAmount = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly topup amount", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTopupAmount returns an API error response when fetching the yearly top-up amount fails.
var ErrApiFailedFindYearlyTopupAmount = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly topup amount", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyWithdrawAmount returns an API error response when fetching the monthly withdrawal amount fails.
var ErrApiFailedFindMonthlyWithdrawAmount = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly withdraw amount", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyWithdrawAmount returns an API error response when fetching the yearly withdrawal amount fails.
var ErrApiFailedFindYearlyWithdrawAmount = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly withdraw amount", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTransactionAmount returns an API error response when fetching the monthly transaction amount fails.
var ErrApiFailedFindMonthlyTransactionAmount = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly transaction amount", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTransactionAmount returns an API error response when fetching the yearly transaction amount fails.
var ErrApiFailedFindYearlyTransactionAmount = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly transaction amount", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTransferSenderAmount returns an API error response when fetching the monthly transfer sender amount fails.
var ErrApiFailedFindMonthlyTransferSenderAmount = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly transfer sender amount", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTransferSenderAmount returns an API error response when fetching the yearly transfer sender amount fails.
var ErrApiFailedFindYearlyTransferSenderAmount = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly transfer sender amount", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTransferReceiverAmount returns an API error response when fetching the monthly transfer receiver amount fails.
var ErrApiFailedFindMonthlyTransferReceiverAmount = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly transfer receiver amount", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTransferReceiverAmount returns an API error response when fetching the yearly transfer receiver amount fails.
var ErrApiFailedFindYearlyTransferReceiverAmount = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly transfer receiver amount", http.StatusInternalServerError)
}

//
