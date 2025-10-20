package transferserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrFailedFindMonthlyTransferAmounts indicates a failure when retrieving the total monthly transfer amounts.
var ErrFailedFindMonthlyTransferAmounts = response.NewErrorResponse("Failed to fetch monthly transfer amounts", http.StatusInternalServerError)

// ErrFailedFindYearlyTransferAmounts indicates a failure when retrieving the total yearly transfer amounts.
var ErrFailedFindYearlyTransferAmounts = response.NewErrorResponse("Failed to fetch yearly transfer amounts", http.StatusInternalServerError)

// ErrFailedFindMonthlyTransferAmountsBySenderCard indicates a failure when retrieving monthly transfer amounts by sender card number.
var ErrFailedFindMonthlyTransferAmountsBySenderCard = response.NewErrorResponse("Failed to fetch monthly transfer amounts by sender card", http.StatusInternalServerError)

// ErrFailedFindMonthlyTransferAmountsByReceiverCard indicates a failure when retrieving monthly transfer amounts by receiver card number.
var ErrFailedFindMonthlyTransferAmountsByReceiverCard = response.NewErrorResponse("Failed to fetch monthly transfer amounts by receiver card", http.StatusInternalServerError)

// ErrFailedFindYearlyTransferAmountsBySenderCard indicates a failure when retrieving yearly transfer amounts by sender card number.
var ErrFailedFindYearlyTransferAmountsBySenderCard = response.NewErrorResponse("Failed to fetch yearly transfer amounts by sender card", http.StatusInternalServerError)

// ErrFailedFindYearlyTransferAmountsByReceiverCard indicates a failure when retrieving yearly transfer amounts by receiver card number.
var ErrFailedFindYearlyTransferAmountsByReceiverCard = response.NewErrorResponse("Failed to fetch yearly transfer amounts by receiver card", http.StatusInternalServerError)
