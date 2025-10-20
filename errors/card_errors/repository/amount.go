package cardrepositoryerrors

import "errors"

// ErrGetMonthlyTopupAmountFailed is returned when fetching monthly top-up amounts fails.
var ErrGetMonthlyTopupAmountFailed = errors.New("failed to get monthly topup amount")

// ErrGetYearlyTopupAmountFailed is returned when fetching yearly top-up amounts fails.
var ErrGetYearlyTopupAmountFailed = errors.New("failed to get yearly topup amount")

// ErrGetMonthlyWithdrawAmountFailed is returned when fetching monthly withdrawal amounts fails.
var ErrGetMonthlyWithdrawAmountFailed = errors.New("failed to get monthly withdraw amount")

// ErrGetYearlyWithdrawAmountFailed is returned when fetching yearly withdrawal amounts fails.
var ErrGetYearlyWithdrawAmountFailed = errors.New("failed to get yearly withdraw amount")

// ErrGetMonthlyTransactionAmountFailed is returned when fetching monthly transaction amounts fails.
var ErrGetMonthlyTransactionAmountFailed = errors.New("failed to get monthly transaction amount")

// ErrGetYearlyTransactionAmountFailed is returned when fetching yearly transaction amounts fails.
var ErrGetYearlyTransactionAmountFailed = errors.New("failed to get yearly transaction amount")

// ErrGetMonthlyTransferAmountSenderFailed is returned when fetching monthly transfer amounts by sender fails.
var ErrGetMonthlyTransferAmountSenderFailed = errors.New("failed to get monthly transfer amount by sender")

// ErrGetYearlyTransferAmountSenderFailed is returned when fetching yearly transfer amounts by sender fails.
var ErrGetYearlyTransferAmountSenderFailed = errors.New("failed to get yearly transfer amount by sender")

// ErrGetMonthlyTransferAmountReceiverFailed is returned when fetching monthly transfer amounts by receiver fails.
var ErrGetMonthlyTransferAmountReceiverFailed = errors.New("failed to get monthly transfer amount by receiver")

// ErrGetYearlyTransferAmountReceiverFailed is returned when fetching yearly transfer amounts by receiver fails.
var ErrGetYearlyTransferAmountReceiverFailed = errors.New("failed to get yearly transfer amount by receiver")
