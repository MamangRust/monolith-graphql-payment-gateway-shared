package transferserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrFailedFindMonthTransferStatusSuccess indicates a failure when retrieving monthly successful transfer statistics.
var ErrFailedFindMonthTransferStatusSuccess = response.NewErrorResponse("Failed to fetch monthly successful transfers", http.StatusInternalServerError)

// ErrFailedFindYearTransferStatusSuccess indicates a failure when retrieving yearly successful transfer statistics.
var ErrFailedFindYearTransferStatusSuccess = response.NewErrorResponse("Failed to fetch yearly successful transfers", http.StatusInternalServerError)

// ErrFailedFindMonthTransferStatusFailed indicates a failure when retrieving monthly failed transfer statistics.
var ErrFailedFindMonthTransferStatusFailed = response.NewErrorResponse("Failed to fetch monthly failed transfers", http.StatusInternalServerError)

// ErrFailedFindYearTransferStatusFailed indicates a failure when retrieving yearly failed transfer statistics.
var ErrFailedFindYearTransferStatusFailed = response.NewErrorResponse("Failed to fetch yearly failed transfers", http.StatusInternalServerError)

// ErrFailedFindMonthTransferStatusSuccessByCard indicates a failure when retrieving monthly successful transfers by card number.
var ErrFailedFindMonthTransferStatusSuccessByCard = response.NewErrorResponse("Failed to fetch monthly successful transfers by card", http.StatusInternalServerError)

// ErrFailedFindYearTransferStatusSuccessByCard indicates a failure when retrieving yearly successful transfers by card number.
var ErrFailedFindYearTransferStatusSuccessByCard = response.NewErrorResponse("Failed to fetch yearly successful transfers by card", http.StatusInternalServerError)

// ErrFailedFindMonthTransferStatusFailedByCard indicates a failure when retrieving monthly failed transfers by card number.
var ErrFailedFindMonthTransferStatusFailedByCard = response.NewErrorResponse("Failed to fetch monthly failed transfers by card", http.StatusInternalServerError)

// ErrFailedFindYearTransferStatusFailedByCard indicates a failure when retrieving yearly failed transfers by card number.
var ErrFailedFindYearTransferStatusFailedByCard = response.NewErrorResponse("Failed to fetch yearly failed transfers by card", http.StatusInternalServerError)
