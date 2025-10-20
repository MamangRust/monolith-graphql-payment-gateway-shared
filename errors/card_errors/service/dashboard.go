package cardserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrFailedDashboardCard is an error response when retrieving card dashboard
// fails. Its HTTP status code is 500.
var ErrFailedDashboardCard = response.NewErrorResponse("Failed to get Card dashboard", http.StatusInternalServerError)

// ErrFailedDashboardCardNumber is an error response when retrieving card
// dashboard by card number fails. Its HTTP status code is 500.
var ErrFailedDashboardCardNumber = response.NewErrorResponse("Failed to get Card dashboard by card number", http.StatusInternalServerError)
