package cardserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrFailedFindMonthlyTopupAmount indicates a failure in retrieving the monthly top-up amount.
var ErrFailedFindMonthlyTopupAmount = response.NewErrorResponse("Failed to get monthly topup amount", http.StatusInternalServerError)

// ErrFailedFindYearlyTopupAmount indicates a failure in retrieving the yearly top-up amount.
var ErrFailedFindYearlyTopupAmount = response.NewErrorResponse("Failed to get yearly topup amount", http.StatusInternalServerError)

// ErrFailedFindMonthlyWithdrawAmount indicates a failure in retrieving the monthly withdraw amount.
var ErrFailedFindMonthlyWithdrawAmount = response.NewErrorResponse("Failed to get monthly withdraw amount", http.StatusInternalServerError)

// ErrFailedFindYearlyWithdrawAmount indicates a failure in retrieving the yearly withdraw amount.
var ErrFailedFindYearlyWithdrawAmount = response.NewErrorResponse("Failed to get yearly withdraw amount", http.StatusInternalServerError)

// ErrFailedFindMonthlyTransactionAmount indicates a failure in retrieving the monthly transaction amount.
var ErrFailedFindMonthlyTransactionAmount = response.NewErrorResponse("Failed to get monthly transaction amount", http.StatusInternalServerError)

// ErrFailedFindYearlyTransactionAmount indicates a failure in retrieving the yearly transaction amount.
var ErrFailedFindYearlyTransactionAmount = response.NewErrorResponse("Failed to get yearly transaction amount", http.StatusInternalServerError)

// ErrFailedFindMonthlyTransferAmountSender indicates a failure in retrieving the monthly transfer amount by sender.
var ErrFailedFindMonthlyTransferAmountSender = response.NewErrorResponse("Failed to get monthly transfer amount by sender", http.StatusInternalServerError)

// ErrFailedFindYearlyTransferAmountSender indicates a failure in retrieving the yearly transfer amount by sender.
var ErrFailedFindYearlyTransferAmountSender = response.NewErrorResponse("Failed to get yearly transfer amount by sender", http.StatusInternalServerError)

// ErrFailedFindMonthlyTransferAmountReceiver indicates a failure in retrieving the monthly transfer amount by receiver.
var ErrFailedFindMonthlyTransferAmountReceiver = response.NewErrorResponse("Failed to get monthly transfer amount by receiver", http.StatusInternalServerError)

// ErrFailedFindYearlyTransferAmountReceiver indicates a failure in retrieving the yearly transfer amount by receiver.
var ErrFailedFindYearlyTransferAmountReceiver = response.NewErrorResponse("Failed to get yearly transfer amount by receiver", http.StatusInternalServerError)

//
