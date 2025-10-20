package withdrawrepositoryerrors

import "errors"

// ErrFindAllWithdrawsFailed is used when the system fails to find all withdraws
var ErrFindAllWithdrawsFailed = errors.New("failed to find all withdraws")

// ErrFindActiveWithdrawsFailed is used when the system fails to find active withdraws
var ErrFindActiveWithdrawsFailed = errors.New("failed to find active withdraws")

// ErrFindTrashedWithdrawsFailed is used when the system fails to find trashed withdraws
var ErrFindTrashedWithdrawsFailed = errors.New("failed to find trashed withdraws")

// ErrFindWithdrawsByCardNumberFailed is used when the system fails to find withdraws by card number
var ErrFindWithdrawsByCardNumberFailed = errors.New("failed to find withdraws by card number")

// ErrFindWithdrawByIdFailed is used when the system fails to find a withdraw by ID
var ErrFindWithdrawByIdFailed = errors.New("failed to find withdraw by ID")
