package cardserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrFailedFindMonthlyBalance indicates a failure in retrieving the monthly balance.
var ErrFailedFindMonthlyBalance = response.NewErrorResponse("Failed to get monthly balance", http.StatusInternalServerError)

// ErrFailedFindYearlyBalance indicates a failure in retrieving the yearly balance.
var ErrFailedFindYearlyBalance = response.NewErrorResponse("Failed to get yearly balance", http.StatusInternalServerError)

// ErrFailedFindMonthlyBalanceByCard returns an error response when retrieving
// monthly balance by card number fails. Its HTTP status code is 500.
var ErrFailedFindMonthlyBalanceByCard = response.NewErrorResponse("Failed to get monthly balance by card", http.StatusInternalServerError)

// ErrFailedFindYearlyBalanceByCard returns an error response when retrieving
// yearly balance by card number fails. Its HTTP status code is 500.
var ErrFailedFindYearlyBalanceByCard = response.NewErrorResponse("Failed to get yearly balance by card", http.StatusInternalServerError)
