package cardrepositoryerrors

import "errors"

// ErrGetTotalBalancesFailed is returned when fetching total balances fails.
var ErrGetTotalBalancesFailed = errors.New("failed to get total balances")

// ErrGetTotalTopAmountFailed is returned when fetching the total top-up amount fails.
var ErrGetTotalTopAmountFailed = errors.New("failed to get total topup amount")

// ErrGetTotalWithdrawAmountFailed is returned when fetching the total withdrawal amount fails.
var ErrGetTotalWithdrawAmountFailed = errors.New("failed to get total withdraw amount")

// ErrGetTotalTransactionAmountFailed is returned when fetching the total transaction amount fails.
var ErrGetTotalTransactionAmountFailed = errors.New("failed to get total transaction amount")

// ErrGetTotalTransferAmountFailed is returned when fetching the total transfer amount fails.
var ErrGetTotalTransferAmountFailed = errors.New("failed to get total transfer amount")

// ErrGetTotalBalanceByCardFailed is returned when fetching the total balance by card fails.
var ErrGetTotalBalanceByCardFailed = errors.New("failed to get total balance by card")

// ErrGetTotalTopupAmountByCardFailed is returned when fetching the total top-up amount by card fails.
var ErrGetTotalTopupAmountByCardFailed = errors.New("failed to get total topup amount by card")

// ErrGetTotalWithdrawAmountByCardFailed is returned when fetching the total withdrawal amount by card fails.
var ErrGetTotalWithdrawAmountByCardFailed = errors.New("failed to get total withdraw amount by card")

// ErrGetTotalTransactionAmountByCardFailed is returned when fetching the total transaction amount by card fails.
var ErrGetTotalTransactionAmountByCardFailed = errors.New("failed to get total transaction amount by card")

// ErrGetTotalTransferAmountBySenderFailed is returned when fetching the total transfer amount by sender fails.
var ErrGetTotalTransferAmountBySenderFailed = errors.New("failed to get total transfer amount by sender")

// ErrGetTotalTransferAmountByReceiverFailed is returned when fetching the total transfer amount by receiver fails.
var ErrGetTotalTransferAmountByReceiverFailed = errors.New("failed to get total transfer amount by receiver")
