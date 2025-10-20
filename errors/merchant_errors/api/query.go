package merchantapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiFailedFindAllMerchants is returned when failing to retrieve all merchants.
var ErrApiFailedFindAllMerchants = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all merchants", http.StatusInternalServerError)
}

// ErrApiFailedFindAllMerchantsActive is returned when failing to retrieve all active merchants.
var ErrApiFailedFindAllMerchantsActive = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all merchants active", http.StatusInternalServerError)
}

// ErrApiFailedFindAllMerchantsTrashed is returned when failing to retrieve all trashed merchants.
var ErrApiFailedFindAllMerchantsTrashed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all merchants trashed", http.StatusInternalServerError)
}

// ErrApiFailedFindByUserId is returned when failing to retrieve a merchant by user ID.
var ErrApiFailedFindByUserId = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch merchant by user ID", http.StatusInternalServerError)
}

// ErrApiFailedFindByIdMerchant is returned when failing to retrieve a merchant by ID.
var ErrApiFailedFindByIdMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch merchant by ID", http.StatusInternalServerError)
}

// ErrApiFailedFindByApiKeyMerchant is returned when failing to retrieve a merchant by API key.
var ErrApiFailedFindByApiKeyMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch merchant by api key", http.StatusInternalServerError)
}
