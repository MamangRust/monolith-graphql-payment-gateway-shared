package withdrawserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrFailedCreateWithdraw is used when failed to create withdraw
var ErrFailedCreateWithdraw = response.NewErrorResponse("Failed to create withdraw", http.StatusInternalServerError)

// ErrFailedUpdateWithdraw is used when failed to update withdraw
var ErrFailedUpdateWithdraw = response.NewErrorResponse("Failed to update withdraw", http.StatusInternalServerError)

// ErrFailedTrashedWithdraw is used when failed to trash withdraw
var ErrFailedTrashedWithdraw = response.NewErrorResponse("Failed to trash withdraw", http.StatusInternalServerError)

// ErrFailedRestoreWithdraw is used when failed to restore withdraw
var ErrFailedRestoreWithdraw = response.NewErrorResponse("Failed to restore withdraw", http.StatusInternalServerError)

// ErrFailedDeleteWithdrawPermanent is used when failed to permanently delete withdraw
var ErrFailedDeleteWithdrawPermanent = response.NewErrorResponse("Failed to permanently delete withdraw", http.StatusInternalServerError)

// ErrFailedRestoreAllWithdraw is used when failed to restore all withdraws
var ErrFailedRestoreAllWithdraw = response.NewErrorResponse("Failed to restore all withdraws", http.StatusInternalServerError)

// ErrFailedDeleteAllWithdrawPermanent is used when failed to permanently delete all withdraws
var ErrFailedDeleteAllWithdrawPermanent = response.NewErrorResponse("Failed to permanently delete all withdraws", http.StatusInternalServerError)
