package withdrawserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrFailedSendEmail is used when failed to send email
var ErrFailedSendEmail = response.NewErrorResponse("Failed to send email", http.StatusInternalServerError)
