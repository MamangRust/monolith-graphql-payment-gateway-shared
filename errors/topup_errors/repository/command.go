package topuprepositoryerrors

import "errors"

// ErrCreateTopupFailed indicates failure in creating a new top-up record.
var ErrCreateTopupFailed = errors.New("failed to create topup")

// ErrUpdateTopupFailed indicates failure in updating an existing top-up record.
var ErrUpdateTopupFailed = errors.New("failed to update topup")

// ErrUpdateTopupAmountFailed indicates failure in updating only the top-up amount.
var ErrUpdateTopupAmountFailed = errors.New("failed to update topup amount")

// ErrUpdateTopupStatusFailed indicates failure in updating the top-up status (e.g., success/failed).
var ErrUpdateTopupStatusFailed = errors.New("failed to update topup status")

// ErrTrashedTopupFailed indicates failure in soft-deleting (trashing) a top-up.
var ErrTrashedTopupFailed = errors.New("failed to soft-delete (trash) topup")

// ErrRestoreTopupFailed indicates failure in restoring a previously trashed top-up.
var ErrRestoreTopupFailed = errors.New("failed to restore topup")

// ErrDeleteTopupPermanentFailed indicates failure in permanently deleting a top-up.
var ErrDeleteTopupPermanentFailed = errors.New("failed to permanently delete topup")

// ErrRestoreAllTopupFailed indicates failure in restoring all trashed top-ups.
var ErrRestoreAllTopupFailed = errors.New("failed to restore all topups")

// ErrDeleteAllTopupPermanentFailed indicates failure in permanently deleting all trashed top-ups.
var ErrDeleteAllTopupPermanentFailed = errors.New("failed to permanently delete all topups")
