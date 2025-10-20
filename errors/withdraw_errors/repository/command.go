package withdrawrepositoryerrors

import "errors"

// ErrCreateWithdrawFailed indicates a failure when creating a new withdraw record.
var ErrCreateWithdrawFailed = errors.New("failed to create withdraw")

// ErrUpdateWithdrawFailed indicates a failure when updating a withdraw record.
var ErrUpdateWithdrawFailed = errors.New("failed to update withdraw")

// ErrUpdateWithdrawStatusFailed indicates a failure when updating the status of a withdraw record.
var ErrUpdateWithdrawStatusFailed = errors.New("failed to update withdraw status")

// ErrTrashedWithdrawFailed indicates a failure when soft-deleting (trashing) a withdraw record.
var ErrTrashedWithdrawFailed = errors.New("failed to soft-delete (trash) withdraw")

// ErrRestoreWithdrawFailed indicates a failure when restoring a previously trashed withdraw record.
var ErrRestoreWithdrawFailed = errors.New("failed to restore withdraw")

// ErrDeleteWithdrawPermanentFailed indicates a failure when permanently deleting a withdraw record.
var ErrDeleteWithdrawPermanentFailed = errors.New("failed to permanently delete withdraw")

// ErrRestoreAllWithdrawsFailed indicates a failure when restoring all trashed withdraw records.
var ErrRestoreAllWithdrawsFailed = errors.New("failed to restore all withdraws")

// ErrDeleteAllWithdrawsPermanentFailed indicates a failure when permanently deleting all withdraw records.
var ErrDeleteAllWithdrawsPermanentFailed = errors.New("failed to permanently delete all withdraws")
