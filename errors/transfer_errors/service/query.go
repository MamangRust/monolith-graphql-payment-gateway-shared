package transferserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrFailedFindAllTransfers indicates a failure when retrieving all transfer records.
var ErrFailedFindAllTransfers = response.NewErrorResponse("Failed to fetch all transfers", http.StatusInternalServerError)

// ErrTransferNotFound indicates that a specific transfer record was not found.
var ErrTransferNotFound = response.NewErrorResponse("Transfer not found", http.StatusNotFound)

// ErrFailedFindActiveTransfers indicates a failure when retrieving active transfer records.
var ErrFailedFindActiveTransfers = response.NewErrorResponse("Failed to fetch active transfers", http.StatusInternalServerError)

// ErrFailedFindTrashedTransfers indicates a failure when retrieving trashed (soft-deleted) transfer records.
var ErrFailedFindTrashedTransfers = response.NewErrorResponse("Failed to fetch trashed transfers", http.StatusInternalServerError)

// ErrFailedFindTransfersBySender indicates a failure when retrieving transfers filtered by sender card.
var ErrFailedFindTransfersBySender = response.NewErrorResponse("Failed to fetch transfers by sender", http.StatusInternalServerError)

// ErrFailedFindTransfersByReceiver indicates a failure when retrieving transfers filtered by receiver card.
var ErrFailedFindTransfersByReceiver = response.NewErrorResponse("Failed to fetch transfers by receiver", http.StatusInternalServerError)
