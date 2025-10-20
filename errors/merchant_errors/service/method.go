package merchantserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrFailedFindMonthlyPaymentMethodsMerchant indicates failure fetching monthly payment methods.
var ErrFailedFindMonthlyPaymentMethodsMerchant = response.NewErrorResponse("Failed to get monthly payment methods", http.StatusInternalServerError)

// ErrFailedFindYearlyPaymentMethodMerchant indicates failure fetching yearly payment methods.
var ErrFailedFindYearlyPaymentMethodMerchant = response.NewErrorResponse("Failed to get yearly payment method", http.StatusInternalServerError)

// ErrFailedFindMonthlyPaymentMethodByMerchants indicates failure fetching monthly payment methods by merchant.
var ErrFailedFindMonthlyPaymentMethodByMerchants = response.NewErrorResponse("Failed to get monthly payment methods by Merchant", http.StatusInternalServerError)

// ErrFailedFindYearlyPaymentMethodByMerchants indicates failure fetching yearly payment methods by merchant.
var ErrFailedFindYearlyPaymentMethodByMerchants = response.NewErrorResponse("Failed to get yearly payment method by Merchant", http.StatusInternalServerError)

// ErrFailedFindMonthlyPaymentMethodByApikeys indicates failure fetching monthly payment methods by API key.
var ErrFailedFindMonthlyPaymentMethodByApikeys = response.NewErrorResponse("Failed to get monthly payment methods by API key", http.StatusInternalServerError)

// ErrFailedFindYearlyPaymentMethodByApikeys indicates failure fetching yearly payment methods by API key.
var ErrFailedFindYearlyPaymentMethodByApikeys = response.NewErrorResponse("Failed to get yearly payment method by API key", http.StatusInternalServerError)
