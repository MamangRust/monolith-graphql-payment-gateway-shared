package saldorepositoryerrors

import "errors"

// ErrCreateSaldoFailed is returned when creating a saldo record fails.
var ErrCreateSaldoFailed = errors.New("failed to create saldo record")

// ErrUpdateSaldoFailed is returned when updating a saldo record fails.
var ErrUpdateSaldoFailed = errors.New("failed to update saldo record")

// ErrUpdateSaldoBalanceFailed is returned when updating saldo balance fails.
var ErrUpdateSaldoBalanceFailed = errors.New("failed to update saldo balance")

// ErrUpdateSaldoWithdrawFailed is returned when updating saldo for a withdrawal fails.
var ErrUpdateSaldoWithdrawFailed = errors.New("failed to update saldo withdrawal")

// ErrTrashSaldoFailed is returned when soft-deleting (trashing) a saldo record fails.
var ErrTrashSaldoFailed = errors.New("failed to trash saldo record")

// ErrRestoreSaldoFailed is returned when restoring a trashed saldo record fails.
var ErrRestoreSaldoFailed = errors.New("failed to restore saldo record")

// ErrDeleteSaldoPermanentFailed is returned when permanently deleting a saldo record fails.
var ErrDeleteSaldoPermanentFailed = errors.New("failed to delete saldo permanently")

// ErrRestoreAllSaldosFailed is returned when restoring all trashed saldo records fails.
var ErrRestoreAllSaldosFailed = errors.New("failed to restore all saldo records")

// ErrDeleteAllSaldosPermanentFailed is returned when permanently deleting all saldo records fails.
var ErrDeleteAllSaldosPermanentFailed = errors.New("failed to delete all saldo records permanently")
