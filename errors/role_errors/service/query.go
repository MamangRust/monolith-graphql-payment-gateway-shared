package roleserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrRoleNotFoundRes is returned when the requested role is not found.
var ErrRoleNotFoundRes = response.NewErrorResponse("Role not found", http.StatusNotFound)

// ErrFailedFindAll is returned when there is a failure in fetching all roles.
var ErrFailedFindAll = response.NewErrorResponse("Failed to fetch Roles", http.StatusInternalServerError)

// ErrFailedFindActive is returned when there is a failure in fetching active roles.
var ErrFailedFindActive = response.NewErrorResponse("Failed to fetch active Roles", http.StatusInternalServerError)

// ErrFailedFindTrashed is returned when there is a failure in fetching trashed roles.
var ErrFailedFindTrashed = response.NewErrorResponse("Failed to fetch trashed Roles", http.StatusInternalServerError)
