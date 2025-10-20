package merchantserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrFailedCreateMerchant indicates failure when creating a new merchant.
var ErrFailedCreateMerchant = response.NewErrorResponse("Failed to create Merchant", http.StatusInternalServerError)

// ErrFailedUpdateMerchant indicates failure when updating a merchant.
var ErrFailedUpdateMerchant = response.NewErrorResponse("Failed to update Merchant", http.StatusInternalServerError)

// ErrFailedTrashMerchant indicates failure when soft-deleting (trashing) a merchant.
var ErrFailedTrashMerchant = response.NewErrorResponse("Failed to trash Merchant", http.StatusInternalServerError)

// ErrFailedRestoreMerchant indicates failure when restoring a trashed merchant.
var ErrFailedRestoreMerchant = response.NewErrorResponse("Failed to restore Merchant", http.StatusInternalServerError)

// ErrFailedDeleteMerchant indicates failure when permanently deleting a merchant.
var ErrFailedDeleteMerchant = response.NewErrorResponse("Failed to delete Merchant permanently", http.StatusInternalServerError)

// ErrFailedRestoreAllMerchants indicates failure when restoring all trashed merchants.
var ErrFailedRestoreAllMerchants = response.NewErrorResponse("Failed to restore all Merchants", http.StatusInternalServerError)

// ErrFailedDeleteAllMerchants indicates failure when permanently deleting all trashed merchants.
var ErrFailedDeleteAllMerchants = response.NewErrorResponse("Failed to delete all Merchants permanently", http.StatusInternalServerError)
