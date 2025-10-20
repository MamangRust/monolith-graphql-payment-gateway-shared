package saldoserviceerror

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrFailedFindAllSaldos returns a 500 error when fetching all saldo records fails.
var ErrFailedFindAllSaldos = response.NewErrorResponse("Failed to fetch saldos", http.StatusInternalServerError)

// ErrFailedSaldoNotFound returns a 404 error when a requested saldo record is not found.
var ErrFailedSaldoNotFound = response.NewErrorResponse("Saldo not found", http.StatusNotFound)

// ErrFailedFindSaldoByCardNumber returns a 500 error when fetching saldo by card number fails.
var ErrFailedFindSaldoByCardNumber = response.NewErrorResponse("Failed to find saldo by card number", http.StatusInternalServerError)

// ErrFailedFindActiveSaldos returns a 500 error when fetching active saldo records fails.
var ErrFailedFindActiveSaldos = response.NewErrorResponse("Failed to fetch active saldos", http.StatusInternalServerError)

// ErrFailedFindTrashedSaldos returns a 500 error when fetching trashed saldo records fails.
var ErrFailedFindTrashedSaldos = response.NewErrorResponse("Failed to fetch trashed saldos", http.StatusInternalServerError)
