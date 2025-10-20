package withdrawserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrFailedFindMonthlyWithdraws is used when failed to fetch monthly withdraw amounts
var ErrFailedFindMonthlyWithdraws = response.NewErrorResponse("Failed to fetch monthly withdraw amounts", http.StatusInternalServerError)

// ErrFailedFindYearlyWithdraws is used when failed to fetch yearly withdraw amounts
var ErrFailedFindYearlyWithdraws = response.NewErrorResponse("Failed to fetch yearly withdraw amounts", http.StatusInternalServerError)

// ErrFailedFindMonthlyWithdrawsByCardNumber is used when failed to fetch monthly withdraw amounts by card
var ErrFailedFindMonthlyWithdrawsByCardNumber = response.NewErrorResponse("Failed to fetch monthly withdraw amounts by card", http.StatusInternalServerError)

// ErrFailedFindYearlyWithdrawsByCardNumber is used when failed to fetch yearly withdraw amounts by card
var ErrFailedFindYearlyWithdrawsByCardNumber = response.NewErrorResponse("Failed to fetch yearly withdraw amounts by card", http.StatusInternalServerError)
