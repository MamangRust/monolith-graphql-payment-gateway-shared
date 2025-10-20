package merchantserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrMerchantNotFoundRes is returned when the merchant is not found.
var ErrMerchantNotFoundRes = response.NewErrorResponse("Merchant not found", http.StatusNotFound)

// ErrFailedFindAllMerchants indicates failure in fetching all merchants.
var ErrFailedFindAllMerchants = response.NewErrorResponse("Failed to fetch Merchants", http.StatusInternalServerError)

// ErrFailedFindActiveMerchants indicates failure in fetching active merchants.
var ErrFailedFindActiveMerchants = response.NewErrorResponse("Failed to fetch active Merchants", http.StatusInternalServerError)

// ErrFailedFindTrashedMerchants indicates failure in fetching trashed merchants.
var ErrFailedFindTrashedMerchants = response.NewErrorResponse("Failed to fetch trashed Merchants", http.StatusInternalServerError)

// ErrFailedFindMerchantById is returned when a merchant cannot be found by ID.
var ErrFailedFindMerchantById = response.NewErrorResponse("Failed to find Merchant by ID", http.StatusInternalServerError)

// ErrFailedFindByApiKey is returned when a merchant cannot be found by API key.
var ErrFailedFindByApiKey = response.NewErrorResponse("Failed to find Merchant by API key", http.StatusInternalServerError)

// ErrFailedFindByMerchantUserId is returned when a merchant cannot be found by user ID.
var ErrFailedFindByMerchantUserId = response.NewErrorResponse("Failed to find Merchant by User ID", http.StatusInternalServerError)
