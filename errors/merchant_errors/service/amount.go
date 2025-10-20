package merchantserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrFailedFindMonthlyAmountMerchant indicates failure fetching monthly amounts.
var ErrFailedFindMonthlyAmountMerchant = response.NewErrorResponse("Failed to get monthly amount", http.StatusInternalServerError)

// ErrFailedFindYearlyAmountMerchant indicates failure fetching yearly amounts.
var ErrFailedFindYearlyAmountMerchant = response.NewErrorResponse("Failed to get yearly amount", http.StatusInternalServerError)

// ErrFailedFindMonthlyAmountByMerchants indicates failure fetching monthly amounts by merchant.
var ErrFailedFindMonthlyAmountByMerchants = response.NewErrorResponse("Failed to get monthly amount by Merchant", http.StatusInternalServerError)

// ErrFailedFindYearlyAmountByMerchants indicates failure fetching yearly amounts by merchant.
var ErrFailedFindYearlyAmountByMerchants = response.NewErrorResponse("Failed to get yearly amount by Merchant", http.StatusInternalServerError)

// ErrFailedFindMonthlyAmountByApikeys indicates failure fetching monthly amounts by API key.
var ErrFailedFindMonthlyAmountByApikeys = response.NewErrorResponse("Failed to get monthly amount by API key", http.StatusInternalServerError)

// ErrFailedFindYearlyAmountByApikeys indicates failure fetching yearly amounts by API key.
var ErrFailedFindYearlyAmountByApikeys = response.NewErrorResponse("Failed to get yearly amount by API key", http.StatusInternalServerError)
