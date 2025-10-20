package cardrepositoryerrors

import "errors"

// ErrGetMonthlyBalanceFailed is returned when fetching monthly balances fails.
var ErrGetMonthlyBalanceFailed = errors.New("failed to get monthly balance")

// ErrGetYearlyBalanceFailed is returned when fetching yearly balances fails.
var ErrGetYearlyBalanceFailed = errors.New("failed to get yearly balance")

// ErrGetMonthlyBalanceByCardFailed is returned when fetching monthly balances by card fails.
var ErrGetMonthlyBalanceByCardFailed = errors.New("failed to get monthly balance by card")

// ErrGetYearlyBalanceByCardFailed is returned when fetching yearly balances by card fails.
var ErrGetYearlyBalanceByCardFailed = errors.New("failed to get yearly balance by card")
