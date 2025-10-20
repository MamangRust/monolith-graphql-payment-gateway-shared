package userserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrFailedCreateUser is returned when creating a user fails.
var ErrFailedCreateUser = response.NewErrorResponse("Failed to create user", http.StatusInternalServerError)

// ErrFailedUpdateUser is returned when updating a user fails.
var ErrFailedUpdateUser = response.NewErrorResponse("Failed to update user", http.StatusInternalServerError)

// ErrFailedTrashedUser is returned when moving a user to trash fails.
var ErrFailedTrashedUser = response.NewErrorResponse("Failed to move user to trash", http.StatusInternalServerError)

// ErrFailedRestoreUser is returned when restoring a user fails.
var ErrFailedRestoreUser = response.NewErrorResponse("Failed to restore user", http.StatusInternalServerError)

// ErrFailedDeletePermanent is returned when permanently deleting a user fails.
var ErrFailedDeletePermanent = response.NewErrorResponse("Failed to delete user permanently", http.StatusInternalServerError)

// ErrFailedRestoreAll is returned when restoring all users fails.
var ErrFailedRestoreAll = response.NewErrorResponse("Failed to restore all users", http.StatusInternalServerError)

// ErrFailedDeleteAll is returned when permanently deleting all users fails.
var ErrFailedDeleteAll = response.NewErrorResponse("Failed to delete all users permanently", http.StatusInternalServerError)
