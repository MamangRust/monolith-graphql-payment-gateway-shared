package withdrawserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrFailedFindMonthWithdrawStatusSuccess is used when failed to fetch monthly successful withdraws
var ErrFailedFindMonthWithdrawStatusSuccess = response.NewErrorResponse("Failed to fetch monthly successful withdraws", http.StatusInternalServerError)

// ErrFailedFindYearWithdrawStatusSuccess is used when failed to fetch yearly successful withdraws
var ErrFailedFindYearWithdrawStatusSuccess = response.NewErrorResponse("Failed to fetch yearly successful withdraws", http.StatusInternalServerError)

// ErrFailedFindMonthWithdrawStatusFailed is used when failed to fetch monthly failed withdraws
var ErrFailedFindMonthWithdrawStatusFailed = response.NewErrorResponse("Failed to fetch monthly failed withdraws", http.StatusInternalServerError)

// ErrFailedFindYearWithdrawStatusFailed is used when failed to fetch yearly failed withdraws
var ErrFailedFindYearWithdrawStatusFailed = response.NewErrorResponse("Failed to fetch yearly failed withdraws", http.StatusInternalServerError)

// ErrFailedFindMonthWithdrawStatusSuccessByCard is used when failed to fetch monthly successful withdraws by card
var ErrFailedFindMonthWithdrawStatusSuccessByCard = response.NewErrorResponse("Failed to fetch monthly successful withdraws by card", http.StatusInternalServerError)

// ErrFailedFindYearWithdrawStatusSuccessByCard is used when failed to fetch yearly successful withdraws by card
var ErrFailedFindYearWithdrawStatusSuccessByCard = response.NewErrorResponse("Failed to fetch yearly successful withdraws by card", http.StatusInternalServerError)

// ErrFailedFindMonthWithdrawStatusFailedByCard is used when failed to fetch monthly failed withdraws by card
var ErrFailedFindMonthWithdrawStatusFailedByCard = response.NewErrorResponse("Failed to fetch monthly failed withdraws by card", http.StatusInternalServerError)

// ErrFailedFindYearWithdrawStatusFailedByCard is used when failed to fetch yearly failed withdraws by card
var ErrFailedFindYearWithdrawStatusFailedByCard = response.NewErrorResponse("Failed to fetch yearly failed withdraws by card", http.StatusInternalServerError)
