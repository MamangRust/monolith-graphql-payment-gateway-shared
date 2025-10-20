package saldorepositoryerrors

import "errors"

// ErrGetMonthlySaldoBalancesFailed is returned when fetching monthly saldo balances fails.
var ErrGetMonthlySaldoBalancesFailed = errors.New("failed to get monthly saldo balances")

// ErrGetYearlySaldoBalancesFailed is returned when fetching yearly saldo balances fails.
var ErrGetYearlySaldoBalancesFailed = errors.New("failed to get yearly saldo balances")
