package roleserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrFailedCreateRole is returned when there is a failure in creating a role.
var ErrFailedCreateRole = response.NewErrorResponse("Failed to create Role", http.StatusInternalServerError)

// ErrFailedUpdateRole is returned when there is a failure in updating a role.
var ErrFailedUpdateRole = response.NewErrorResponse("Failed to update Role", http.StatusInternalServerError)

// ErrFailedTrashedRole is returned when there is a failure in moving a role to trash.
var ErrFailedTrashedRole = response.NewErrorResponse("Failed to move Role to trash", http.StatusInternalServerError)

// ErrFailedRestoreRole is returned when there is a failure in restoring a trashed role.
var ErrFailedRestoreRole = response.NewErrorResponse("Failed to restore Role", http.StatusInternalServerError)

// ErrFailedDeletePermanent is returned when there is a failure in permanently deleting a role.
var ErrFailedDeletePermanent = response.NewErrorResponse("Failed to delete Role permanently", http.StatusInternalServerError)

// ErrFailedRestoreAll is returned when there is a failure in restoring all trashed roles.
var ErrFailedRestoreAll = response.NewErrorResponse("Failed to restore all Roles", http.StatusInternalServerError)

// ErrFailedDeleteAll is returned when there is a failure in permanently deleting all roles.
var ErrFailedDeleteAll = response.NewErrorResponse("Failed to delete all Roles permanently", http.StatusInternalServerError)
