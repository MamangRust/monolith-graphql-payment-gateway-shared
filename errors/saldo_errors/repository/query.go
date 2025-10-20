package saldorepositoryerrors

import "errors"

// ErrFindAllSaldosFailed is returned when fetching all saldo records fails.
var ErrFindAllSaldosFailed = errors.New("failed to find all saldo records")

// ErrFindActiveSaldosFailed is returned when fetching active saldo records fails.
var ErrFindActiveSaldosFailed = errors.New("failed to find active saldo records")

// ErrFindTrashedSaldosFailed is returned when fetching trashed saldo records fails.
var ErrFindTrashedSaldosFailed = errors.New("failed to find trashed saldo records")

// ErrFindSaldoByIdFailed is returned when fetching a saldo by its ID fails.
var ErrFindSaldoByIdFailed = errors.New("failed to find saldo by ID")

// ErrFindSaldoByCardNumberFailed is returned when fetching saldo by card number fails.
var ErrFindSaldoByCardNumberFailed = errors.New("failed to find saldo by card number")
