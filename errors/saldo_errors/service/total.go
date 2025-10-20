package saldoserviceerror

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrFailedFindMonthlyTotalSaldoBalance returns a 500 error when fetching monthly total saldo balance fails.
var ErrFailedFindMonthlyTotalSaldoBalance = response.NewErrorResponse("Failed to fetch monthly total saldo balance", http.StatusInternalServerError)

// ErrFailedFindYearTotalSaldoBalance returns a 500 error when fetching yearly total saldo balance fails.
var ErrFailedFindYearTotalSaldoBalance = response.NewErrorResponse("Failed to fetch yearly total saldo balance", http.StatusInternalServerError)
