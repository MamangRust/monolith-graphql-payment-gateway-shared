package topupserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrFailedFindMonthTopupStatusSuccess indicates failure in retrieving monthly successful top-up status.
var ErrFailedFindMonthTopupStatusSuccess = response.NewErrorResponse("Failed to get monthly topup success status", http.StatusInternalServerError)

// ErrFailedFindYearlyTopupStatusSuccess indicates failure in retrieving yearly successful top-up status.
var ErrFailedFindYearlyTopupStatusSuccess = response.NewErrorResponse("Failed to get yearly topup success status", http.StatusInternalServerError)

// ErrFailedFindMonthTopupStatusFailed indicates failure in retrieving monthly failed top-up status.
var ErrFailedFindMonthTopupStatusFailed = response.NewErrorResponse("Failed to get monthly topup failed status", http.StatusInternalServerError)

// ErrFailedFindYearlyTopupStatusFailed indicates failure in retrieving yearly failed top-up status.
var ErrFailedFindYearlyTopupStatusFailed = response.NewErrorResponse("Failed to get yearly topup failed status", http.StatusInternalServerError)

// ErrFailedFindMonthTopupStatusSuccessByCard indicates failure in retrieving monthly successful top-up status by card.
var ErrFailedFindMonthTopupStatusSuccessByCard = response.NewErrorResponse("Failed to get monthly topup success status by card", http.StatusInternalServerError)

// ErrFailedFindYearlyTopupStatusSuccessByCard indicates failure in retrieving yearly successful top-up status by card.
var ErrFailedFindYearlyTopupStatusSuccessByCard = response.NewErrorResponse("Failed to get yearly topup success status by card", http.StatusInternalServerError)

// ErrFailedFindMonthTopupStatusFailedByCard indicates failure in retrieving monthly failed top-up status by card.
var ErrFailedFindMonthTopupStatusFailedByCard = response.NewErrorResponse("Failed to get monthly topup failed status by card", http.StatusInternalServerError)

// ErrFailedFindYearlyTopupStatusFailedByCard indicates failure in retrieving yearly failed top-up status by card.
var ErrFailedFindYearlyTopupStatusFailedByCard = response.NewErrorResponse("Failed to get yearly topup failed status by card", http.StatusInternalServerError)
