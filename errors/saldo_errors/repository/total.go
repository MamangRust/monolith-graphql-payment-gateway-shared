package saldorepositoryerrors

import "errors"

// ErrGetMonthlyTotalSaldoBalanceFailed is returned when fetching monthly total saldo balance fails.
var ErrGetMonthlyTotalSaldoBalanceFailed = errors.New("failed to get monthly total saldo balance")

// ErrGetYearTotalSaldoBalanceFailed is returned when fetching yearly total saldo balance fails.
var ErrGetYearTotalSaldoBalanceFailed = errors.New("failed to get yearly total saldo balance")
