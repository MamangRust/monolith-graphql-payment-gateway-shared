package transactionrepositoryerrors

import "errors"

// ErrCreateTransactionFailed indicates a failure when creating a new transaction.
var ErrCreateTransactionFailed = errors.New("failed to create transaction")

// ErrUpdateTransactionFailed indicates a failure when updating a transaction.
var ErrUpdateTransactionFailed = errors.New("failed to update transaction")

// ErrUpdateTransactionStatusFailed indicates a failure when updating the status of a transaction.
var ErrUpdateTransactionStatusFailed = errors.New("failed to update transaction status")

// ErrTrashedTransactionFailed indicates a failure when soft-deleting (trashing) a transaction.
var ErrTrashedTransactionFailed = errors.New("failed to soft-delete (trash) transaction")

// ErrRestoreTransactionFailed indicates a failure when restoring a trashed transaction.
var ErrRestoreTransactionFailed = errors.New("failed to restore transaction")

// ErrDeleteTransactionPermanentFailed indicates a failure when permanently deleting a transaction.
var ErrDeleteTransactionPermanentFailed = errors.New("failed to permanently delete transaction")

// ErrRestoreAllTransactionsFailed indicates a failure when restoring all trashed transactions.
var ErrRestoreAllTransactionsFailed = errors.New("failed to restore all transactions")

// ErrDeleteAllTransactionsPermanentFailed indicates a failure when permanently deleting all transactions.
var ErrDeleteAllTransactionsPermanentFailed = errors.New("failed to permanently delete all transactions")
