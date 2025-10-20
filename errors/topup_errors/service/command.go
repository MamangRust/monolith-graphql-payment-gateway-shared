package topupserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrFailedCreateTopup indicates failure in creating a new top-up record.
var ErrFailedCreateTopup = response.NewErrorResponse("Failed to create Topup", http.StatusInternalServerError)

// ErrFailedUpdateTopup indicates failure in updating an existing top-up record.
var ErrFailedUpdateTopup = response.NewErrorResponse("Failed to update Topup", http.StatusInternalServerError)

// ErrFailedTrashTopup indicates failure in soft-deleting (trashing) a top-up.
var ErrFailedTrashTopup = response.NewErrorResponse("Failed to trash Topup", http.StatusInternalServerError)

// ErrFailedRestoreTopup indicates failure in restoring a previously trashed top-up.
var ErrFailedRestoreTopup = response.NewErrorResponse("Failed to restore Topup", http.StatusInternalServerError)

// ErrFailedDeleteTopup indicates failure in permanently deleting a top-up.
var ErrFailedDeleteTopup = response.NewErrorResponse("Failed to delete Topup permanently", http.StatusInternalServerError)

// ErrFailedRestoreAllTopups indicates failure in restoring all trashed top-up records.
var ErrFailedRestoreAllTopups = response.NewErrorResponse("Failed to restore all Topups", http.StatusInternalServerError)

// ErrFailedDeleteAllTopups indicates failure in permanently deleting all trashed top-up records.
var ErrFailedDeleteAllTopups = response.NewErrorResponse("Failed to delete all Topups permanently", http.StatusInternalServerError)
