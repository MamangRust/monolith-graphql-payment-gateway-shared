package topupserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrTopupNotFoundRes indicates that a requested top-up was not found.
var ErrTopupNotFoundRes = response.NewErrorResponse("Topup not found", http.StatusNotFound)

// ErrFailedFindAllTopups indicates failure in retrieving all top-up records.
var ErrFailedFindAllTopups = response.NewErrorResponse("Failed to fetch Topups", http.StatusInternalServerError)

// ErrFailedFindAllTopupsByCardNumber indicates failure in retrieving top-ups by card number.
var ErrFailedFindAllTopupsByCardNumber = response.NewErrorResponse("Failed to fetch Topups by card number", http.StatusInternalServerError)

// ErrFailedFindTopupById indicates failure in finding a top-up by its ID.
var ErrFailedFindTopupById = response.NewErrorResponse("Failed to find Topup by ID", http.StatusInternalServerError)

// ErrFailedFindActiveTopups indicates failure in retrieving active (non-trashed) top-up records.
var ErrFailedFindActiveTopups = response.NewErrorResponse("Failed to fetch active Topups", http.StatusInternalServerError)

// ErrFailedFindTrashedTopups indicates failure in retrieving trashed (soft-deleted) top-up records.
var ErrFailedFindTrashedTopups = response.NewErrorResponse("Failed to fetch trashed Topups", http.StatusInternalServerError)
