package topuprepositoryerrors

import "errors"

// ErrFindAllTopupsFailed indicates failure in retrieving all top-up records from the database.
var ErrFindAllTopupsFailed = errors.New("failed to find all topups")

// ErrFindTopupsByActiveFailed indicates failure in retrieving only the active (non-deleted) top-ups.
var ErrFindTopupsByActiveFailed = errors.New("failed to find active topups")

// ErrFindTopupsByTrashedFailed indicates failure in retrieving trashed (soft-deleted) top-ups.
var ErrFindTopupsByTrashedFailed = errors.New("failed to find trashed topups")

// ErrFindTopupsByCardNumberFailed indicates failure in finding top-ups by a specific card number.
var ErrFindTopupsByCardNumberFailed = errors.New("failed to find topups by card number")

// ErrFindTopupByIdFailed indicates failure in finding a top-up using its unique ID.
var ErrFindTopupByIdFailed = errors.New("failed to find topup by ID")
