package saldoserviceerror

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrFailedInsuffientBalance returns a 400 error when a transaction is attempted with insufficient balance.
var ErrFailedInsuffientBalance = response.NewErrorResponse("Insufficient balance", http.StatusBadRequest)
