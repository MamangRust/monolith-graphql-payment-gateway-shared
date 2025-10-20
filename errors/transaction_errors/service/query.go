package transactonserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrFailedFindAllTransactions indicates a failure when retrieving all transaction records.
var ErrFailedFindAllTransactions = response.NewErrorResponse("Failed to fetch all transactions", http.StatusInternalServerError)

// ErrFailedFindAllByCardNumber indicates a failure when retrieving transactions filtered by card number.
var ErrFailedFindAllByCardNumber = response.NewErrorResponse("Failed to fetch transactions by card number", http.StatusInternalServerError)

// ErrTransactionNotFound indicates that the requested transaction could not be found.
var ErrTransactionNotFound = response.NewErrorResponse("Transaction not found", http.StatusNotFound)

// ErrFailedFindByActiveTransactions indicates a failure when retrieving active (non-deleted) transactions.
var ErrFailedFindByActiveTransactions = response.NewErrorResponse("Failed to fetch active transactions", http.StatusInternalServerError)

// ErrFailedFindByTrashedTransactions indicates a failure when retrieving trashed (soft-deleted) transactions.
var ErrFailedFindByTrashedTransactions = response.NewErrorResponse("Failed to fetch trashed transactions", http.StatusInternalServerError)

// ErrFailedFindByMerchantID indicates a failure when retrieving transactions filtered by merchant ID.
var ErrFailedFindByMerchantID = response.NewErrorResponse("Failed to fetch transactions by merchant ID", http.StatusInternalServerError)
