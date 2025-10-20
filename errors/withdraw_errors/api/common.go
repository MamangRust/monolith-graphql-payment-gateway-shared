package withdrawapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiWithdrawInvalidID is an error response for invalid withdraw ID.
var ErrApiWithdrawInvalidID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid Withdraw ID", http.StatusBadRequest)
}

// ErrApiWithdrawInvalidUserID is an error response for invalid withdraw merchant ID.
var ErrApiWithdrawInvalidUserID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid Withdraw Merchant ID", http.StatusBadRequest)
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
	return response.NewApiErrorResponse(c, "year", "Invalid year", http.StatusBadRequest)
}
