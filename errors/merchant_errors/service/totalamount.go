package merchantserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrFailedFindMonthlyTotalAmountMerchant indicates failure fetching monthly total amounts.
var ErrFailedFindMonthlyTotalAmountMerchant = response.NewErrorResponse("Failed to get monthly total amount", http.StatusInternalServerError)

// ErrFailedFindYearlyTotalAmountMerchant indicates failure fetching yearly total amounts.
var ErrFailedFindYearlyTotalAmountMerchant = response.NewErrorResponse("Failed to get yearly total amount", http.StatusInternalServerError)

// ErrFailedFindMonthlyTotalAmountByMerchants indicates failure fetching monthly total amounts by merchant.
var ErrFailedFindMonthlyTotalAmountByMerchants = response.NewErrorResponse("Failed to get monthly total amount by Merchant", http.StatusInternalServerError)

// ErrFailedFindYearlyTotalAmountByMerchants indicates failure fetching yearly total amounts by merchant.
var ErrFailedFindYearlyTotalAmountByMerchants = response.NewErrorResponse("Failed to get yearly total amount by Merchant", http.StatusInternalServerError)

// ErrFailedFindMonthlyTotalAmountByApikeys indicates failure fetching monthly total amounts by API key.
var ErrFailedFindMonthlyTotalAmountByApikeys = response.NewErrorResponse("Failed to get monthly total amount by API key", http.StatusInternalServerError)

// ErrFailedFindYearlyTotalAmountByApikeys indicates failure fetching yearly total amounts by API key.
var ErrFailedFindYearlyTotalAmountByApikeys = response.NewErrorResponse("Failed to get yearly total amount by API key", http.StatusInternalServerError)
