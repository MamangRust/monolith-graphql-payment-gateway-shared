package merchantapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiFailedFindMonthlyPaymentMethodsMerchant is returned when failing to retrieve monthly payment methods for a merchant.
var ErrApiFailedFindMonthlyPaymentMethodsMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly payment methods by merchant", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyPaymentMethodMerchant is returned when failing to retrieve yearly payment methods for a merchant.
var ErrApiFailedFindYearlyPaymentMethodMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly payment methods by merchant", http.StatusInternalServerError)
}

// by merchant

// ErrApiFailedFindMonthlyPaymentMethodByMerchants is returned when failing to fetch monthly payment methods for all merchants.
var ErrApiFailedFindMonthlyPaymentMethodByMerchants = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly payment methods by merchant", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyPaymentMethodByMerchants is returned when failing to fetch yearly payment methods for all merchants.
var ErrApiFailedFindYearlyPaymentMethodByMerchants = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly payment methods by merchant", http.StatusInternalServerError)
}

// by apikey

// ErrApiFailedFindMonthlyPaymentMethodByApikeys is returned when failing to fetch monthly payment methods by API key.
var ErrApiFailedFindMonthlyPaymentMethodByApikeys = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly payment methods by API key", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyPaymentMethodByApikeys is returned when failing to fetch yearly payment methods by API key.
var ErrApiFailedFindYearlyPaymentMethodByApikeys = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly payment methods by API key", http.StatusInternalServerError)
}
