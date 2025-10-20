package saldoserviceerror

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrFailedFindMonthlySaldoBalances returns a 500 error when fetching monthly saldo balances fails.
var ErrFailedFindMonthlySaldoBalances = response.NewErrorResponse("Failed to fetch monthly saldo balances", http.StatusInternalServerError)

// ErrFailedFindYearlySaldoBalances returns a 500 error when fetching yearly saldo balances fails.
var ErrFailedFindYearlySaldoBalances = response.NewErrorResponse("Failed to fetch yearly saldo balances", http.StatusInternalServerError)
