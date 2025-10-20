package transactonserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrFailedFindMonthlyAmounts indicates a failure when retrieving the total monthly transaction amounts.
var ErrFailedFindMonthlyAmounts = response.NewErrorResponse("Failed to fetch monthly amounts", http.StatusInternalServerError)

// ErrFailedFindYearlyAmounts indicates a failure when retrieving the total yearly transaction amounts.
var ErrFailedFindYearlyAmounts = response.NewErrorResponse("Failed to fetch yearly amounts", http.StatusInternalServerError)

// ErrFailedFindMonthlyAmountsByCard indicates a failure when retrieving monthly transaction amounts filtered by card.
var ErrFailedFindMonthlyAmountsByCard = response.NewErrorResponse("Failed to fetch monthly amounts by card", http.StatusInternalServerError)

// ErrFailedFindYearlyAmountsByCard indicates a failure when retrieving yearly transaction amounts filtered by card.
var ErrFailedFindYearlyAmountsByCard = response.NewErrorResponse("Failed to fetch yearly amounts by card", http.StatusInternalServerError)
