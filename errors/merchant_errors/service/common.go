package merchantserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrFailedSendEmail indicates a failure when sending an email related to merchant operations.
var ErrFailedSendEmail = response.NewErrorResponse("Failed to send email", http.StatusInternalServerError)
