package userserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrUserNotFoundRes is returned when a user is not found.
var ErrUserNotFoundRes = response.NewErrorResponse("User not found", http.StatusNotFound)

// ErrUserEmailAlready is returned when a user email already exists.
var ErrUserEmailAlready = response.NewErrorResponse("User email already exists", http.StatusBadRequest)

// ErrFailedFindAll is returned when fetching users fails.
var ErrFailedFindAll = response.NewErrorResponse("Failed to fetch users", http.StatusInternalServerError)

// ErrFailedFindActive is returned when fetching active users fails.
var ErrFailedFindActive = response.NewErrorResponse("Failed to fetch active users", http.StatusInternalServerError)

// ErrFailedFindTrashed is returned when fetching trashed users fails.
var ErrFailedFindTrashed = response.NewErrorResponse("Failed to fetch trashed users", http.StatusInternalServerError)
