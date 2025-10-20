package saldoserviceerror

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrFailedCreateSaldo returns a 500 error when creating a saldo record fails.
var ErrFailedCreateSaldo = response.NewErrorResponse("Failed to create saldo", http.StatusInternalServerError)

// ErrFailedUpdateSaldo returns a 500 error when updating a saldo record fails.
var ErrFailedUpdateSaldo = response.NewErrorResponse("Failed to update saldo", http.StatusInternalServerError)

// ErrFailedTrashSaldo returns a 500 error when moving a saldo record to trash fails.
var ErrFailedTrashSaldo = response.NewErrorResponse("Failed to trash saldo", http.StatusInternalServerError)

// ErrFailedRestoreSaldo returns a 500 error when restoring a trashed saldo record fails.
var ErrFailedRestoreSaldo = response.NewErrorResponse("Failed to restore saldo", http.StatusInternalServerError)

// ErrFailedDeleteSaldoPermanent returns a 500 error when permanently deleting a saldo record fails.
var ErrFailedDeleteSaldoPermanent = response.NewErrorResponse("Failed to permanently delete saldo", http.StatusInternalServerError)

// ErrFailedRestoreAllSaldo returns a 500 error when restoring all trashed saldo records fails.
var ErrFailedRestoreAllSaldo = response.NewErrorResponse("Failed to restore all saldos", http.StatusInternalServerError)

// ErrFailedDeleteAllSaldoPermanent returns a 500 error when permanently deleting all trashed saldo records fails.
var ErrFailedDeleteAllSaldoPermanent = response.NewErrorResponse("Failed to permanently delete all saldos", http.StatusInternalServerError)
