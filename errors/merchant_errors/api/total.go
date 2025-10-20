package merchantapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiFailedFindMonthlyTotalAmountMerchant is returned when failing to retrieve total monthly transaction amount for a merchant.
var ErrApiFailedFindMonthlyTotalAmountMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly total amount by merchant", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTotalAmountMerchant is returned when failing to retrieve total yearly transaction amount for a merchant.
var ErrApiFailedFindYearlyTotalAmountMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly total amount by merchant", http.StatusInternalServerError)
}
