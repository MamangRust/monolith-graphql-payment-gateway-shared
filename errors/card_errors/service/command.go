package cardserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrFailedCreateCard is returned when creating a new Card record fails.
// Its HTTP status code is 500.
var ErrFailedCreateCard = response.NewErrorResponse("Failed to create Card", http.StatusInternalServerError)

// ErrFailedUpdateCard is returned when updating an existing Card record fails.
// Its HTTP status code is 500.
var ErrFailedUpdateCard = response.NewErrorResponse("Failed to update Card", http.StatusInternalServerError)

// ErrFailedTrashCard is returned when moving a Card to trash fails.
// Its HTTP status code is 500.
var ErrFailedTrashCard = response.NewErrorResponse("Failed to trash Card", http.StatusInternalServerError)

// ErrFailedRestoreCard is returned when restoring a trashed Card fails.
// Its HTTP status code is 500.
var ErrFailedRestoreCard = response.NewErrorResponse("Failed to restore Card", http.StatusInternalServerError)

// ErrFailedDeleteCard is returned when permanently deleting a Card fails.
// Its HTTP status code is 500.
var ErrFailedDeleteCard = response.NewErrorResponse("Failed to delete Card permanently", http.StatusInternalServerError)

// ErrFailedRestoreAllCards is returned when restoring all trashed Cards fails.
// Its HTTP status code is 500.
var ErrFailedRestoreAllCards = response.NewErrorResponse("Failed to restore all Cards", http.StatusInternalServerError)

// ErrFailedDeleteAllCards is returned when permanently deleting all Cards fails.
// Its HTTP status code is 500.
var ErrFailedDeleteAllCards = response.NewErrorResponse("Failed to delete all Cards permanently", http.StatusInternalServerError)
