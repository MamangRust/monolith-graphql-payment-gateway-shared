package transactonserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrFailedCreateTransaction indicates a failure when creating a new transaction record.
var ErrFailedCreateTransaction = response.NewErrorResponse("Failed to create transaction", http.StatusInternalServerError)

// ErrFailedUpdateTransaction indicates a failure when updating an existing transaction record.
var ErrFailedUpdateTransaction = response.NewErrorResponse("Failed to update transaction", http.StatusInternalServerError)

// ErrFailedTrashedTransaction indicates a failure when soft-deleting (trashing) a transaction.
var ErrFailedTrashedTransaction = response.NewErrorResponse("Failed to trash transaction", http.StatusInternalServerError)

// ErrFailedRestoreTransaction indicates a failure when restoring a previously trashed transaction.
var ErrFailedRestoreTransaction = response.NewErrorResponse("Failed to restore transaction", http.StatusInternalServerError)

// ErrFailedDeleteTransactionPermanent indicates a failure when permanently deleting a transaction.
var ErrFailedDeleteTransactionPermanent = response.NewErrorResponse("Failed to permanently delete transaction", http.StatusInternalServerError)

// ErrFailedRestoreAllTransactions indicates a failure when restoring all trashed transactions.
var ErrFailedRestoreAllTransactions = response.NewErrorResponse("Failed to restore all transactions", http.StatusInternalServerError)

// ErrFailedDeleteAllTransactionsPermanent indicates a failure when permanently deleting all transactions.
var ErrFailedDeleteAllTransactionsPermanent = response.NewErrorResponse("Failed to permanently delete all transactions", http.StatusInternalServerError)
