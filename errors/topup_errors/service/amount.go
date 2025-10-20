package topupserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrFailedFindMonthlyTopupAmounts indicates failure in retrieving monthly top-up amounts.
var ErrFailedFindMonthlyTopupAmounts = response.NewErrorResponse("Failed to get monthly topup amounts", http.StatusInternalServerError)

// ErrFailedFindYearlyTopupAmounts indicates failure in retrieving yearly top-up amounts.
var ErrFailedFindYearlyTopupAmounts = response.NewErrorResponse("Failed to get yearly topup amounts", http.StatusInternalServerError)

// ErrFailedFindMonthlyTopupAmountsByCard indicates failure in retrieving monthly top-up amounts by card.
var ErrFailedFindMonthlyTopupAmountsByCard = response.NewErrorResponse("Failed to get monthly topup amounts by card", http.StatusInternalServerError)

// ErrFailedFindYearlyTopupAmountsByCard indicates failure in retrieving yearly top-up amounts by card.
var ErrFailedFindYearlyTopupAmountsByCard = response.NewErrorResponse("Failed to get yearly topup amounts by card", http.StatusInternalServerError)
