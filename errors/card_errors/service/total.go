package cardserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrFailedFindTotalBalances is an error response when retrieving total balances
// fails. Its HTTP status code is 500.
var ErrFailedFindTotalBalances = response.NewErrorResponse("Failed to Find total balances", http.StatusInternalServerError)

// ErrFailedFindTotalTopAmount is an error response when retrieving total topup
// amount fails. Its HTTP status code is 500.
var ErrFailedFindTotalTopAmount = response.NewErrorResponse("Failed to Find total topup amount", http.StatusInternalServerError)

// ErrFailedFindTotalWithdrawAmount is an error response when retrieving total
// withdraw amount fails. Its HTTP status code is 500.
var ErrFailedFindTotalWithdrawAmount = response.NewErrorResponse("Failed to Find total withdraw amount", http.StatusInternalServerError)

// ErrFailedFindTotalTransactionAmount is an error response when retrieving total
// transaction amount fails. Its HTTP status code is 500.
var ErrFailedFindTotalTransactionAmount = response.NewErrorResponse("Failed to Find total transaction amount", http.StatusInternalServerError)

// ErrFailedFindTotalTransferAmount is an error response when retrieving total
// transfer amount fails. Its HTTP status code is 500.
var ErrFailedFindTotalTransferAmount = response.NewErrorResponse("Failed to Find total transfer amount", http.StatusInternalServerError)

// ErrFailedFindTotalBalanceByCard is an error response when retrieving total
// balance by card fails. Its HTTP status code is 500.
var ErrFailedFindTotalBalanceByCard = response.NewErrorResponse("Failed to Find total balance by card", http.StatusInternalServerError)

// ErrFailedFindTotalTopupAmountByCard is an error response when retrieving total
// topup amount by card fails. Its HTTP status code is 500.
var ErrFailedFindTotalTopupAmountByCard = response.NewErrorResponse("Failed to Find total topup amount by card", http.StatusInternalServerError)

// ErrFailedFindTotalWithdrawAmountByCard is an error response when retrieving
// total withdraw amount by card fails. Its HTTP status code is 500.
var ErrFailedFindTotalWithdrawAmountByCard = response.NewErrorResponse("Failed to Find total withdraw amount by card", http.StatusInternalServerError)

// ErrFailedFindTotalTransactionAmountByCard is an error response when
// retrieving total transaction amount by card fails. Its HTTP status code is
// 500.
var ErrFailedFindTotalTransactionAmountByCard = response.NewErrorResponse("Failed to Find total transaction amount by card", http.StatusInternalServerError)

// ErrFailedFindTotalTransferAmountBySender is an error response when
// retrieving total transfer amount by sender fails. Its HTTP status code is
// 500.
var ErrFailedFindTotalTransferAmountBySender = response.NewErrorResponse("Failed to Find total transfer amount by sender", http.StatusInternalServerError)

// ErrFailedFindTotalTransferAmountByReceiver is an error response when
// retrieving total transfer amount by receiver fails. Its HTTP status code is
// 500.
var ErrFailedFindTotalTransferAmountByReceiver = response.NewErrorResponse("Failed to Find total transfer amount by receiver", http.StatusInternalServerError)
