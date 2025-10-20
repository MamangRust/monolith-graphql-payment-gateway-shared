package topupserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrFailedFindMonthlyTopupMethods indicates failure in retrieving monthly top-up methods.
var ErrFailedFindMonthlyTopupMethods = response.NewErrorResponse("Failed to get monthly topup methods", http.StatusInternalServerError)

// ErrFailedFindYearlyTopupMethods indicates failure in retrieving yearly top-up methods.
var ErrFailedFindYearlyTopupMethods = response.NewErrorResponse("Failed to get yearly topup methods", http.StatusInternalServerError)

// ErrFailedFindMonthlyTopupMethodsByCard indicates failure in retrieving monthly top-up methods by card.
var ErrFailedFindMonthlyTopupMethodsByCard = response.NewErrorResponse("Failed to get monthly topup methods by card", http.StatusInternalServerError)

// ErrFailedFindYearlyTopupMethodsByCard indicates failure in retrieving yearly top-up methods by card.
var ErrFailedFindYearlyTopupMethodsByCard = response.NewErrorResponse("Failed to get yearly topup methods by card", http.StatusInternalServerError)
