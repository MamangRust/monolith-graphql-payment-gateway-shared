package saldoapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

var ErrApiFailedCreateSaldo = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to create saldo", http.StatusInternalServerError)
}

var ErrApiFailedUpdateSaldo = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to update saldo", http.StatusInternalServerError)
}

// ErrApiFailedTrashSaldo is returned when soft-deleting saldo fails.
var ErrApiFailedTrashSaldo = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to trash saldo", http.StatusInternalServerError)
}

// ErrApiFailedRestoreSaldo is returned when restoring saldo fails.
var ErrApiFailedRestoreSaldo = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore saldo", http.StatusInternalServerError)
}

// ErrApiFailedDeleteSaldoPermanent is returned when permanently deleting saldo fails.
var ErrApiFailedDeleteSaldoPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete saldo", http.StatusInternalServerError)
}

// ErrApiFailedRestoreAllSaldo is returned when restoring all saldos fails.
var ErrApiFailedRestoreAllSaldo = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore all saldos", http.StatusInternalServerError)
}

// ErrApiFailedDeleteAllSaldoPermanent is returned when permanently deleting all saldos fails.
var ErrApiFailedDeleteAllSaldoPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete all saldos", http.StatusInternalServerError)
}
