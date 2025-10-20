package withdrawapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiWithdrawNotFound is an error response for withdraw not found.
var ErrApiWithdrawNotFound = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Withdraw not found", http.StatusNotFound)
}

// ErrApiFailedFindAllWithdraw is an error response for failed to fetch all withdraws.
var ErrApiFailedFindAllWithdraw = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all withdraws", http.StatusInternalServerError)
}

// ErrApiFailedFindAllWithdrawByCardNumber is an error response for failed to fetch withdraws by card number.
var ErrApiFailedFindAllWithdrawByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch withdraws by card number", http.StatusInternalServerError)
}

// ErrApiFailedFindByIdWithdraw is an error response for failed to fetch withdraw by ID.
var ErrApiFailedFindByIdWithdraw = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch withdraw by ID", http.StatusInternalServerError)
}

// ErrApiFailedFindByCardNumber is an error response for failed to fetch withdraws using card number.
var ErrApiFailedFindByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch withdraws using card number", http.StatusInternalServerError)
}

// ErrApiFailedFindByActiveWithdraw is an error response for failed to fetch active withdraws.
var ErrApiFailedFindByActiveWithdraw = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch active withdraws", http.StatusInternalServerError)
}

// ErrApiFailedFindByTrashedWithdraw is an error response for failed to fetch trashed withdraws.
var ErrApiFailedFindByTrashedWithdraw = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch trashed withdraws", http.StatusInternalServerError)
}
