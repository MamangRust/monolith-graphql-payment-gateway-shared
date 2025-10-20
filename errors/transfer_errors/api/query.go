package transactionapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiFailedFindAllTransfers is an error response for failed to fetch all transfers.
var ErrApiFailedFindAllTransfers = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all transfers", http.StatusInternalServerError)
}

// ErrApiFailedFindByIdTransfer is an error response for failed to fetch transfer by ID.
var ErrApiFailedFindByIdTransfer = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch transfer by ID", http.StatusInternalServerError)
}

// ErrApiFailedFindByTransferFrom is an error response for failed to fetch transfers by transfer_from.
var ErrApiFailedFindByTransferFrom = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch transfers by transfer_from", http.StatusInternalServerError)
}

// ErrApiFailedFindByTransferTo is an error response for failed to fetch transfers by transfer_to.
var ErrApiFailedFindByTransferTo = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch transfers by transfer_to", http.StatusInternalServerError)
}

// ErrApiFailedFindByActiveTransfer is an error response for failed to fetch active transfers.
var ErrApiFailedFindByActiveTransfer = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch active transfers", http.StatusInternalServerError)
}

// ErrApiFailedFindByTrashedTransfer returns an API error response when fetching trashed (soft-deleted) transfers fails.
var ErrApiFailedFindByTrashedTransfer = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch trashed transfers", http.StatusInternalServerError)
}
