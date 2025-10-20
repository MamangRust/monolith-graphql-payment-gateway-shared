package withdrawserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrFailedFindAllWithdraws is used when failed to fetch all withdraws
var ErrFailedFindAllWithdraws = response.NewErrorResponse("Failed to fetch all withdraws", http.StatusInternalServerError)

// ErrWithdrawNotFound is used when withdraw is not found
var ErrWithdrawNotFound = response.NewErrorResponse("Withdraw not found", http.StatusNotFound)

// ErrFailedFindAllWithdrawsByCard is used when failed to fetch all withdraws by card number
var ErrFailedFindAllWithdrawsByCard = response.NewErrorResponse("Failed to fetch all withdraws by card number", http.StatusInternalServerError)

// ErrFailedFindActiveWithdraws is used when failed to fetch active withdraws
var ErrFailedFindActiveWithdraws = response.NewErrorResponse("Failed to fetch active withdraws", http.StatusInternalServerError)

// ErrFailedFindTrashedWithdraws is used when failed to fetch trashed withdraws
var ErrFailedFindTrashedWithdraws = response.NewErrorResponse("Failed to fetch trashed withdraws", http.StatusInternalServerError)
