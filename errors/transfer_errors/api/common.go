package transactionapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiTransferInvalidID is an error response for invalid transfer ID.
var ErrApiTransferInvalidID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid Transfer ID", http.StatusBadRequest)
}

// ErrApiTransferInvalidMerchantID is an error response for invalid transfer merchant ID.
var ErrApiTransferInvalidMerchantID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid Transfer Merchant ID", http.StatusBadRequest)
}

// ErrApiInvalidCardNumber is an error response for invalid card number.
var ErrApiInvalidCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid card number", http.StatusBadRequest)
}

// ErrApiInvalidMonth is an error response for invalid month.
var ErrApiInvalidMonth = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid month", http.StatusBadRequest)
}

// ErrApiInvalidYear is an error response for invalid year.
var ErrApiInvalidYear = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid year", http.StatusBadRequest)
}
