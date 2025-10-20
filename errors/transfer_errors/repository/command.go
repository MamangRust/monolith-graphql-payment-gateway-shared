package transferrepositoryerrors

import "errors"

// ErrCreateTransferFailed indicates a failure when creating a new transfer.
var ErrCreateTransferFailed = errors.New("failed to create transfer")

// ErrUpdateTransferFailed indicates a failure when updating an existing transfer.
var ErrUpdateTransferFailed = errors.New("failed to update transfer")

// ErrUpdateTransferAmountFailed indicates a failure when updating the amount of a transfer.
var ErrUpdateTransferAmountFailed = errors.New("failed to update transfer amount")

// ErrUpdateTransferStatusFailed indicates a failure when updating the status of a transfer.
var ErrUpdateTransferStatusFailed = errors.New("failed to update transfer status")

// ErrTrashedTransferFailed indicates a failure when soft-deleting (trashing) a transfer.
var ErrTrashedTransferFailed = errors.New("failed to soft-delete (trash) transfer")

// ErrRestoreTransferFailed indicates a failure when restoring a previously trashed transfer.
var ErrRestoreTransferFailed = errors.New("failed to restore transfer")

// ErrDeleteTransferPermanentFailed indicates a failure when permanently deleting a transfer.
var ErrDeleteTransferPermanentFailed = errors.New("failed to permanently delete transfer")

// ErrRestoreAllTransfersFailed indicates a failure when restoring all trashed transfers.
var ErrRestoreAllTransfersFailed = errors.New("failed to restore all transfers")

// ErrDeleteAllTransfersPermanentFailed indicates a failure when permanently deleting all transfers.
var ErrDeleteAllTransfersPermanentFailed = errors.New("failed to permanently delete all transfers")
