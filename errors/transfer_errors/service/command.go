package transferserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrFailedCreateTransfer indicates a failure when attempting to create a new transfer record.
var ErrFailedCreateTransfer = response.NewErrorResponse("Failed to create transfer", http.StatusInternalServerError)

// ErrFailedUpdateTransfer indicates a failure when attempting to update an existing transfer record.
var ErrFailedUpdateTransfer = response.NewErrorResponse("Failed to update transfer", http.StatusInternalServerError)

// ErrFailedTrashedTransfer indicates a failure when attempting to soft-delete (trash) a transfer.
var ErrFailedTrashedTransfer = response.NewErrorResponse("Failed to trash transfer", http.StatusInternalServerError)

// ErrFailedRestoreTransfer indicates a failure when attempting to restore a previously trashed transfer.
var ErrFailedRestoreTransfer = response.NewErrorResponse("Failed to restore transfer", http.StatusInternalServerError)

// ErrFailedDeleteTransferPermanent indicates a failure when attempting to permanently delete a transfer.
var ErrFailedDeleteTransferPermanent = response.NewErrorResponse("Failed to permanently delete transfer", http.StatusInternalServerError)

// ErrFailedRestoreAllTransfers indicates a failure when attempting to restore all trashed transfers.
var ErrFailedRestoreAllTransfers = response.NewErrorResponse("Failed to restore all transfers", http.StatusInternalServerError)

// ErrFailedDeleteAllTransfersPermanent indicates a failure when attempting to permanently delete all transfer records.
var ErrFailedDeleteAllTransfersPermanent = response.NewErrorResponse("Failed to permanently delete all transfers", http.StatusInternalServerError)
