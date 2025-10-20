package merchantapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiFailedFindMonthlyAmountMerchant is returned when failing to retrieve monthly transaction amount for a merchant.
var ErrApiFailedFindMonthlyAmountMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly amount by merchant", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyAmountMerchant is returned when failing to retrieve yearly transaction amount for a merchant.
var ErrApiFailedFindYearlyAmountMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly amount by merchant", http.StatusInternalServerError)
}

// by merchant

// ErrApiFailedFindMonthlyAmountByMerchants is returned when failing to fetch monthly amounts for all merchants.
var ErrApiFailedFindMonthlyAmountByMerchants = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly amounts by merchant", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyAmountByMerchants is returned when failing to fetch yearly amounts for all merchants.
var ErrApiFailedFindYearlyAmountByMerchants = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly amounts by merchant", http.StatusInternalServerError)
}

// by apikey

// ErrApiFailedFindMonthlyAmountByApikeys is returned when failing to fetch monthly amounts by API key.
var ErrApiFailedFindMonthlyAmountByApikeys = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly amounts by API key", http.StatusInternalServerError)
}

var ErrApiFailedFindYearlyAmountByApikeys = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly amounts by API key", http.StatusInternalServerError)
}
