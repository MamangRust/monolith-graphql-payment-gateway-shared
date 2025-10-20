package transactonserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrFailedFindMonthTransactionSuccess indicates a failure when retrieving monthly successful transactions.
var ErrFailedFindMonthTransactionSuccess = response.NewErrorResponse("Failed to fetch monthly successful transactions", http.StatusInternalServerError)

// ErrFailedFindYearTransactionSuccess indicates a failure when retrieving yearly successful transactions.
var ErrFailedFindYearTransactionSuccess = response.NewErrorResponse("Failed to fetch yearly successful transactions", http.StatusInternalServerError)

// ErrFailedFindMonthTransactionFailed indicates a failure when retrieving monthly failed transactions.
var ErrFailedFindMonthTransactionFailed = response.NewErrorResponse("Failed to fetch monthly failed transactions", http.StatusInternalServerError)

// ErrFailedFindYearTransactionFailed indicates a failure when retrieving yearly failed transactions.
var ErrFailedFindYearTransactionFailed = response.NewErrorResponse("Failed to fetch yearly failed transactions", http.StatusInternalServerError)

// ErrFailedFindMonthTransactionSuccessByCard indicates a failure when retrieving monthly successful transactions filtered by card number.
var ErrFailedFindMonthTransactionSuccessByCard = response.NewErrorResponse("Failed to fetch monthly successful transactions by card", http.StatusInternalServerError)

// ErrFailedFindYearTransactionSuccessByCard indicates a failure when retrieving yearly successful transactions filtered by card number.
var ErrFailedFindYearTransactionSuccessByCard = response.NewErrorResponse("Failed to fetch yearly successful transactions by card", http.StatusInternalServerError)

// ErrFailedFindMonthTransactionFailedByCard indicates a failure when retrieving monthly failed transactions filtered by card number.
var ErrFailedFindMonthTransactionFailedByCard = response.NewErrorResponse("Failed to fetch monthly failed transactions by card", http.StatusInternalServerError)

// ErrFailedFindYearTransactionFailedByCard indicates a failure when retrieving yearly failed transactions filtered by card number.
var ErrFailedFindYearTransactionFailedByCard = response.NewErrorResponse("Failed to fetch yearly failed transactions by card", http.StatusInternalServerError)
