package cardrepositoryerrors

import "errors"

// ErrGetMonthlyTopupAmountByCardFailed is returned when fetching monthly top-up amounts by card fails.
var ErrGetMonthlyTopupAmountByCardFailed = errors.New("failed to get monthly topup amount by card")

// ErrGetYearlyTopupAmountByCardFailed is returned when fetching yearly top-up amounts by card fails.
var ErrGetYearlyTopupAmountByCardFailed = errors.New("failed to get yearly topup amount by card")

// ErrGetMonthlyWithdrawAmountByCardFailed is returned when fetching monthly withdrawal amounts by card fails.
var ErrGetMonthlyWithdrawAmountByCardFailed = errors.New("failed to get monthly withdraw amount by card")

// ErrGetYearlyWithdrawAmountByCardFailed is returned when fetching yearly withdrawal amounts by card fails.
var ErrGetYearlyWithdrawAmountByCardFailed = errors.New("failed to get yearly withdraw amount by card")

// ErrGetMonthlyTransactionAmountByCardFailed is returned when fetching monthly transaction amounts by card fails.
var ErrGetMonthlyTransactionAmountByCardFailed = errors.New("failed to get monthly transaction amount by card")

// ErrGetYearlyTransactionAmountByCardFailed is returned when fetching yearly transaction amounts by card fails.
var ErrGetYearlyTransactionAmountByCardFailed = errors.New("failed to get yearly transaction amount by card")

// ErrGetMonthlyTransferAmountBySenderFailed is returned when fetching monthly transfer amount by sender fails.
var ErrGetMonthlyTransferAmountBySenderFailed = errors.New("failed to get monthly transfer amount by sender")

// ErrGetYearlyTransferAmountBySenderFailed is returned when fetching yearly transfer amount by sender fails.
var ErrGetYearlyTransferAmountBySenderFailed = errors.New("failed to get yearly transfer amount by sender")

// ErrGetMonthlyTransferAmountByReceiverFailed is returned when fetching monthly transfer amount by receiver fails.
var ErrGetMonthlyTransferAmountByReceiverFailed = errors.New("failed to get monthly transfer amount by receiver")

// ErrGetYearlyTransferAmountByReceiverFailed is returned when fetching yearly transfer amount by receiver fails.
var ErrGetYearlyTransferAmountByReceiverFailed = errors.New("failed to get yearly transfer amount by receiver")
