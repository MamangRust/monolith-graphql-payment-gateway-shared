package cardserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrFailedFindMonthlyTopupAmountByCard returns an error response when retrieving
// monthly top-up amount by card number fails. Its HTTP status code is 500.
var ErrFailedFindMonthlyTopupAmountByCard = response.NewErrorResponse("Failed to get monthly topup amount by card", http.StatusInternalServerError)

// ErrFailedFindYearlyTopupAmountByCard returns an error response when retrieving
// yearly top-up amount by card number fails. Its HTTP status code is 500.
var ErrFailedFindYearlyTopupAmountByCard = response.NewErrorResponse("Failed to get yearly topup amount by card", http.StatusInternalServerError)

// ErrFailedFindMonthlyWithdrawAmountByCard returns an error response when retrieving
// monthly withdraw amount by card number fails. Its HTTP status code is 500.
var ErrFailedFindMonthlyWithdrawAmountByCard = response.NewErrorResponse("Failed to get monthly withdraw amount by card", http.StatusInternalServerError)

// ErrFailedFindYearlyWithdrawAmountByCard returns an error response when retrieving
// yearly withdraw amount by card number fails. Its HTTP status code is 500.
var ErrFailedFindYearlyWithdrawAmountByCard = response.NewErrorResponse("Failed to get yearly withdraw amount by card", http.StatusInternalServerError)

// ErrFailedFindMonthlyTransactionAmountByCard returns an error response when retrieving
// monthly transaction amount by card number fails. Its HTTP status code is 500.
var ErrFailedFindMonthlyTransactionAmountByCard = response.NewErrorResponse("Failed to get monthly transaction amount by card", http.StatusInternalServerError)

// ErrFailedFindYearlyTransactionAmountByCard returns an error response when retrieving
// yearly transaction amount by card number fails. Its HTTP status code is 500.
var ErrFailedFindYearlyTransactionAmountByCard = response.NewErrorResponse("Failed to get yearly transaction amount by card", http.StatusInternalServerError)

// ErrFailedFindMonthlyTransferAmountBySender returns an error response when retrieving
// monthly transfer amount by sender card number fails. Its HTTP status code is 500.
var ErrFailedFindMonthlyTransferAmountBySender = response.NewErrorResponse("Failed to get monthly transfer amount by sender", http.StatusInternalServerError)

// ErrFailedFindYearlyTransferAmountBySender returns an error response when retrieving
// yearly transfer amount by sender card number fails. Its HTTP status code is 500.
var ErrFailedFindYearlyTransferAmountBySender = response.NewErrorResponse("Failed to get yearly transfer amount by sender", http.StatusInternalServerError)

// ErrFailedFindMonthlyTransferAmountByReceiver returns an error response when retrieving
// monthly transfer amount by receiver card number fails. Its HTTP status code is 500.
var ErrFailedFindMonthlyTransferAmountByReceiver = response.NewErrorResponse("Failed to get monthly transfer amount by receiver", http.StatusInternalServerError)

// ErrFailedFindYearlyTransferAmountByReceiver returns an error response when retrieving
// yearly transfer amount by receiver card number fails. Its HTTP status code is 500.
var ErrFailedFindYearlyTransferAmountByReceiver = response.NewErrorResponse("Failed to get yearly transfer amount by receiver", http.StatusInternalServerError)
